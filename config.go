package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

// Represents the struct extracted from the yaml config file
type Config struct {
	Logto     LogtoConfig
	Resources []LogtoResource
	Roles     []LogtoRole
}

// The basic logto config for communicating w/ the Logto instance.
type LogtoConfig struct {
	Url       string
	AppID     string
	AppSecret string
}

// Represents a logto resource
//
// This is for defining RBAC on each resource
type LogtoResource struct {
	BaseUrl   string
	Endpoints map[string][]string
}

// Logto roles
// See testdata folder for config examples
type LogtoRole map[string]map[string][]string

func NewConfig(filename string) (Config, error) {
	var k = koanf.New("::")
	k.Load(file.Provider(filename), yaml.Parser())

	// log.Println(k.Raw())

	var config Config
	k.Unmarshal("", &config)
	// log.Println("config unmarshal: ", config)

	return config, nil
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (c *Config) FetchAccessTokenResponse() (AccessTokenResponse, error) {
	bodyParams := map[string]string{
		"grant_type": "client_credentials",
		"resource":   "https://default.logto.app/api",
		"scope":      "all",
	}

	formData := url.Values{}
	for key, value := range bodyParams {
		formData.Set(key, value)
	}

	req, err := http.NewRequest(http.MethodPost, c.Logto.Url+"/oidc/token", bytes.NewBufferString(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.Logto.AppID, c.Logto.AppSecret)

	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("HTTP Request error: %v", err)
	}

	log.Println("HTTP request: ", req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("HTTP Response error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AccessTokenResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	var result AccessTokenResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return AccessTokenResponse{}, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return result, nil
}

func (c *Config) ProvisionLogto(accessTokenResp AccessTokenResponse) error {
	// Provision resources
	// for _, resource := range c.Resources {
	// 	for endpoint, scopes := range resource.Endpoints {
	// 		// TODO: create resource for each endpoint
	// 		// TODO: create scopes for each resource
	// 	}
	// }

	// Create an API resource
	createdResource, err := createResource(c, accessTokenResp, "default", "testing api", "http://api.test.io")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	// Provision roles

	// Provision users

	return nil
}

type Resource struct {
	TenantId       string `json:"tenant_id"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	Indicator      string `json:"indicator"`
	IsDefault      bool   `json:"isDefault"`
	AccessTokenTtl int    `json:"accessTokenTtl"`
}

type ResourceWithScopes struct {
	Resource
	Scopes []Scope `json:"scopes"`
}

type Scope struct {
	TenantId    string `json:"tenant_id"`
	Id          string `json:"id"`
	ResourceId  string `json:"resourceId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int    `json:"createdAt"`
}

// createResource creates a new resource in the Logto system.
//
// TODO: add unit tests
func createResource(c *Config, accessTokenResp AccessTokenResponse, tenantId, name, indicator string) (ResourceWithScopes, error) {
	bodyData := map[string]string{
		"tenantId":  tenantId,
		"name":      name,
		"indicator": indicator,
	}

	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return ResourceWithScopes{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Logto.Url+"/api/resources", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return ResourceWithScopes{}, err
	}
	req.Header.Set("Authorization", accessTokenResp.TokenType+" "+accessTokenResp.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ResourceWithScopes{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnprocessableEntity {
		log.Printf("Unprocessable Entity : already created resource.\nThe request were ignored. Please check your Logto dashboard.\nFull response details: %v", resp)
		// TODO: better handle this case.
		return ResourceWithScopes{}, nil
	}

	if resp.StatusCode != http.StatusCreated {
		return ResourceWithScopes{}, fmt.Errorf("unexpected status code: %v\nResponse: %v", resp.StatusCode, resp)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResourceWithScopes{}, err
	}

	var result ResourceWithScopes
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ResourceWithScopes{}, err
	}
	fmt.Println("Created API Resource: ", result)

	return result, nil
}

// func (c *Config) provisionResource(accessTokenResp AccessTokenResponse, resource LogtoResource) error {

// 	return nil
// }

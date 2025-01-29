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

func (c Config) FetchAccessTokenResponse() (AccessTokenResponse, error) {
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

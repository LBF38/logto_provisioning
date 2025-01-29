package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openapi "github.com/lbf38/logto_provisioning/logtoClient/client"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("Logto Provisioning")

	// Load provisioning config from file
	config, err := NewConfig("provisioning/logto_provisioning_dev.yaml")

	if err != nil {
		log.Fatalf("Failed loading config file: %v\n", err)
	}

	// Fetch access token for interacting w/ Management API
	accessTokenResponse, err := config.FetchAccessTokenResponse()
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}
	log.Printf("Access Token response: %v\n", accessTokenResponse)

	// Create OpenAPI instance
	cfg := openapi.NewConfiguration()
	client := openapi.NewAPIClient(cfg)
	fmt.Printf("OpenAPI configuration: %+v\n", *cfg)
	fmt.Printf("OpenAPI client: %+v\n", *client)

	// OpenAPI OAuth2
	token := oauth2.Token{
		AccessToken: accessTokenResponse.AccessToken,
	}
	tokenSource := oauth2.StaticTokenSource(&token)
	auth := context.WithValue(context.Background(), openapi.ContextOAuth2, tokenSource)

	//! Provision Logto instance
	// Resources provisioning
	req := *openapi.NewCreateResourceRequest("test openapi", "http://test.io")
	resp, r, err := client.ResourcesAPI.CreateResource(auth).CreateResourceRequest(req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.CreateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResource`: ListResources200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.CreateResource`: %v\n", resp)

	resp2, r, err := client.ResourcesAPI.ListResources(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ListResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ListResources`: %v\n", resp2)

	// Roles provisioning

	// Users provisioning

	// Organizations provisioning
}

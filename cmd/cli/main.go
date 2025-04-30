package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/lbf38/logto_provisioning/internal"
	openapi "github.com/lbf38/logto_provisioning/logtoClient/client"
	"golang.org/x/oauth2"
)

func main() {
	configFlag := flag.String("config", "logto.provisioning.yml", "Specify the config file for provisioning Logto")

	flag.Parse()

	fmt.Println("Logto Provisioning")

	// Load provisioning config from file
	config, err := internal.NewConfig(*configFlag)

	if err != nil {
		log.Fatalf("Failed loading config file: %v\n", err)
	}

	// DONE: handle the logto URL defined in the provisioning file.
	// => it should be replaced in the OpenAPI servers config.
	if config.Logto.Url != "" {
		internal.OpenapiCfg.Servers = openapi.ServerConfigurations{
			{
				URL:         config.Logto.Url,
				Description: "Logto instance URL",
			},
		}
	}

	// Fetch access token for interacting w/ Management API
	accessTokenResponse, err := config.FetchAccessTokenResponse()
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}
	log.Printf("Access Token response: %v\n", accessTokenResponse)

	// OpenAPI OAuth2
	token := oauth2.Token{
		AccessToken: accessTokenResponse.AccessToken,
	}
	tokenSource := oauth2.StaticTokenSource(&token)
	auth := context.WithValue(context.Background(), openapi.ContextOAuth2, tokenSource)

	//! Provision Logto instance
	// Resources provisioning

	var resourceScopesMap = make(map[string]map[string]string)
	for _, resource := range config.Resources {
		for endpoint, scopes := range resource.Endpoints {
			indicator := resource.BaseUrl + endpoint
			createdResource, err := internal.CreateResource(auth, endpoint, indicator)
			if err != nil {
				log.Fatalf("Resources provisioning: %v", err)
			}
			resourceScopesMap[indicator] = make(map[string]string)
			for _, scope := range scopes {
				createdScope, err := internal.CreateResourceScope(auth, scope, "description", createdResource)
				if err != nil {
					log.Fatalf("Resources provisioning: %v", err)
				}
				fmt.Println(createdScope)
				resourceScopesMap[indicator][scope] = createdScope.Id
			}
		}
	}

	// Roles provisioning
	for _, role := range config.Roles {
		for role_name, resources := range role {
			fmt.Println(resources)
			// DONE: link each role w/ resources scopes
			scopeIds := []string{}
			for indicator, scopes := range resources {
				for _, scope := range scopes {
					scopeIds = append(scopeIds, resourceScopesMap[indicator][scope])
				}
			}
			fmt.Println("scopeIds: ", scopeIds)

			createdRole, err := internal.CreateRole(auth, role_name, scopeIds)
			if err != nil {
				log.Fatalf("Roles provisioning: %v", err)
			}
			fmt.Println(createdRole)
		}
	}

	// Users provisioning

	// Organizations provisioning

}

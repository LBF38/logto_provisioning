package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Logto Provisioning")

	config, err := NewConfig("provisioning/logto_provisioning_dev.yaml")

	if err != nil {
		log.Fatalf("Failed loading config file: %v\n", err)
	}

	// Fetch access token
	accessTokenResponse, err := config.FetchAccessTokenResponse()
	if err != nil {
		log.Fatalf("Error getting access token: %v", err)
	}

	// Interact w/ Management API (provisioning, fetching data, ...)
	log.Printf("Access Token response: %v\n", accessTokenResponse)

	// Provision Logto
	err = config.ProvisionLogto(accessTokenResponse)
	if err != nil {
		log.Fatalf("Error from provisioning: %v\n", err)
	}
}

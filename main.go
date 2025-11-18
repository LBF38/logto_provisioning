package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

func main() {
	file, err := os.ReadFile("testdata/config_minimal.yaml")
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	data := make(map[string]interface{})
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	fmt.Println("data = ", data)

	var config LogtoConfig
	err = mapstructure.Decode(data, &config)
	if err != nil {
		log.Fatalf("error decoding config: %v", err)
	}

	fmt.Println("config = ", config)

	fmt.Println("Logto url : ", config.Url)
	fmt.Println("parent's ID is = ", config.Resources)
}

package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := NewConfig("testdata/config_minimal.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Config: %+v\n", config)
}

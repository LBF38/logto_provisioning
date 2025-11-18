package main

import (
	"io"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
)

type LogtoConfig struct {
	Url        string          `yaml:"url"`
	App_id     string          `yaml:"app_id"`
	App_secret string          `yaml:"app_secret"`
	Resources  []LogtoResource `yaml:"resources"`
	Roles      []LogtoRole
}

type LogtoResource struct {
	Name        string
	Permissions []string
}

type LogtoRole struct {
	Name        string
	Permissions []string
}

func NewConfig(filename string) (LogtoConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return LogtoConfig{}, err
	}
	defer file.Close()

	log.Println("Reading config file")

	var data []byte
	data, err = io.ReadAll(file)
	if err != nil {
		return LogtoConfig{}, err
	}
	log.Printf("Data read %v", data)

	var config LogtoConfig
	// yaml.Unmarshal(data, &config)
	err = mapstructure.Decode(data, &config)
	if err != nil {
		return LogtoConfig{}, err
	}
	log.Printf("Config %v", config)

	return config, nil
}

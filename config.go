package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logto     LogtoConfig     `yaml:"logto"`
	Resources []Resource      `yaml:"resources"`
	Roles     map[string]Role `yaml:"roles"`
}

type LogtoConfig struct {
	URL       string `yaml:"url"`
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
}

type Resource struct {
	BaseURL   string              `yaml:"baseUrl"`
	Endpoints map[string]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Permissions []string `yaml:"permissions"`
}

type Role struct {
	Permissions map[string][]string `yaml:",inline"`
}

func NewConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

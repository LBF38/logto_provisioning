package main

import (
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

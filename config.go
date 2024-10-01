package main

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type Config struct {
	Logto     LogtoConfig
	Resources []LogtoResource
	Roles     []LogtoRole
}

type LogtoConfig struct {
	Url       string
	AppID     string
	AppSecret string
}

type LogtoResource struct {
	BaseUrl   string
	Endpoints map[string][]string
}

type LogtoRole struct {
	name        string
	permissions []string
}

func NewConfig(filename string) (Config, error) {
	var k = koanf.New("::")
	k.Load(file.Provider(filename), yaml.Parser())

	// log.Println(k.Raw())

	var config Config
	k.Unmarshal("", &config)
	// log.Println("config unmarshal: ", config)

	return config, nil
}

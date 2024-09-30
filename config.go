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
	name        string
	permissions []string
}

type LogtoRole struct {
	name        string
	permissions []string
}

func NewConfig(filename string) (Config, error) {
	var k = koanf.New("::")
	k.Load(file.Provider(filename), yaml.Parser())

	return Config{
		Logto: LogtoConfig{
			Url:       k.String("logto::url"),
			AppID:     k.String("logto::appID"),
			AppSecret: k.String("logto::appSecret"),
		},
	}, nil
}

package main

type LogtoConfig struct {
	url        string
	app_id     string
	app_secret string
	resources  []LogtoResource
	roles      []LogtoRole
}

type LogtoResource struct {
	name        string
	permissions []string
}

type LogtoRole struct {
	name        string
	permissions []string
}

func NewConfig(filename string) (LogtoConfig, error) {
	return LogtoConfig{}, nil
}

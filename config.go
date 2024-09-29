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
	return LogtoConfig{
		url:        "https://[tenant-id].logto.app/api",
		app_id:     "APP_ID",
		app_secret: "APP_SECRET",
	}, nil
}

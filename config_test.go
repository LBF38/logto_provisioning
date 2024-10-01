package main

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		file string
		want Config
	}{
		{
			name: "Logto config",
			file: "testdata/config_logto.yaml",
			want: Config{
				Logto: LogtoConfig{
					Url:       "https://[tenant-id].logto.app/api",
					AppID:     "APP_ID",
					AppSecret: "APP_SECRET",
				},
			},
		},
		{
			name: "Logto config 2",
			file: "testdata/config_logto_2.yaml",
			want: Config{
				Logto: LogtoConfig{
					Url:       "https://localhost:3001/api",
					AppID:     "LOCAL_APP_ID",
					AppSecret: "LOCAL_APP_SECRET",
				},
			},
		},
		{
			name: "Resources config",
			file: "testdata/config_resources.yaml",
			want: Config{
				Resources: []LogtoResource{
					{
						BaseUrl: "https://api.store.io",
						Endpoints: map[string][]string{
							"/orders": {
								"read:order",
								"write:order",
								"delete:order",
							},
							"/products": {
								"read:product",
								"write:product",
								"delete:product",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.file)

			if err != nil {
				t.Errorf("Error should be nil, but got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		file string
		want LogtoConfig
	}{
		{
			name: "Logto config",
			file: "testdata/config_logto.yaml",
			want: LogtoConfig{
				url:        "https://[tenant-id].logto.app/api",
				app_id:     "APP_ID",
				app_secret: "APP_SECRET",
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

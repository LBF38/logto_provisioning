package main

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("Logto config", func(t *testing.T) {
		got, err := NewConfig("testdata/config_logto.yaml")
		want := LogtoConfig{
			url:        "https://[tenant-id].logto.app/api",
			app_id:     "APP_ID",
			app_secret: "APP_SECRET",
		}

		if err != nil {
			t.Errorf("Error should be nil, but got %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	// got, err := NewConfig("testdata/config_minimal.yaml")
	// want := LogtoConfig{
	// 	url:        "https://[tenant-id].logto.app/api",
	// 	app_id:     "APP_ID",
	// 	app_secret: "APP_SECRET",
	// 	resources: []LogtoResource{
	// 		{
	// 			name:        "https://api.store.io/orders",
	// 			permissions: []string{"read:order", "write:order", "delete:order"},
	// 		},
	// 	},
	// 	roles: []LogtoRole{
	// 		{
	// 			name:        "order_admin",
	// 			permissions: []string{"read:order", "write:order", "delete:order"},
	// 		},
	// 	},
	// }

	// if err != nil {
	// 	t.Errorf("Error should be nil, but got %v", err)
	// }

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("got %v, want %v", got, want)
	// }

}

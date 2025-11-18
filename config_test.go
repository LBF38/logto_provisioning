package main

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	got, err := NewConfig("testdata/config_minimal.yaml")
	want := LogtoConfig{
		Url:        "https://[tenant-id].logto.app/api",
		App_id:     "APP_ID",
		App_secret: "APP_SECRET",
		Resources: []LogtoResource{
			{
				Name:        "https://api.store.io/orders",
				Permissions: []string{"read:order", "write:order", "delete:order"},
			},
		},
		Roles: []LogtoRole{
			{
				Name:        "order_admin",
				Permissions: []string{"read:order", "write:order", "delete:order"},
			},
		},
	}

	if err != nil {
		t.Errorf("Error should be nil, but got %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}

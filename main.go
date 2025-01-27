package main

import (
	"fmt"

	openapi "github.com/lbf38/logto_provisioning/logtoClient/client"
)

func main() {
	fmt.Println("Logto Provisioning")
	// config, err := NewConfig("testdata/config_logto.yaml")
	// if err != nil {
	// 	panic(err)
	// }
	cfg := openapi.NewConfiguration()
	client := openapi.NewAPIClient(cfg)
	ConnectLogto(client)
}

package main

import (
	"context"
	"fmt"

	openapi "github.com/lbf38/logto_provisioning/logtoClient/client"
)

func ConnectLogto(client *openapi.APIClient) {
	createResourceReq := *openapi.NewCreateResourceRequest("hello", "something")

	resourcesList, response, err := client.ResourcesAPI.CreateResource(context.Background()).CreateResourceRequest(createResourceReq).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("resources: ", resourcesList, "\nresponse: ", response)
}

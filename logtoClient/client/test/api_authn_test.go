/*
Logto API references

Testing AuthnAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/lbf38/logto_provisioning/logtoClient/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_AuthnAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthnAPIService AssertSaml", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		httpRes, err := apiClient.AuthnAPI.AssertSaml(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthnAPIService AssertSingleSignOnSaml", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		httpRes, err := apiClient.AuthnAPI.AssertSingleSignOnSaml(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthnAPIService GetHasuraAuth", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AuthnAPI.GetHasuraAuth(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

/*
ForestVPN API

Testing AppleApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package forestvpn_api_test

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "github.com/forestvpn/api-client-go"
)

func Test_forestvpn_api_AppleApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AppleApiService VerifyAppStoreReceipt", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.AppleApi.VerifyAppStoreReceipt(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

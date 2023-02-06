/*
ForestVPN API

Testing BillingApiService

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

func Test_forestvpn_api_BillingApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test BillingApiService CancelSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionID string

        resp, httpRes, err := apiClient.BillingApi.CancelSubscription(context.Background(), subscriptionID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService CreatePaymentMethodStripeSetupIntent", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.CreatePaymentMethodStripeSetupIntent(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService CreateSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.CreateSubscription(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService DeletePaymentMethod", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var paymentMethodID string

        resp, httpRes, err := apiClient.BillingApi.DeletePaymentMethod(context.Background(), paymentMethodID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService GetBillingAccount", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.GetBillingAccount(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService GetBillingBundle", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var bundleID string

        resp, httpRes, err := apiClient.BillingApi.GetBillingBundle(context.Background(), bundleID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService GetBillingPaymentOption", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var paymentOptionID string

        resp, httpRes, err := apiClient.BillingApi.GetBillingPaymentOption(context.Background(), paymentOptionID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService GetBillingProduct", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var productID string

        resp, httpRes, err := apiClient.BillingApi.GetBillingProduct(context.Background(), productID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService GetPaymentMethod", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var paymentMethodID string

        resp, httpRes, err := apiClient.BillingApi.GetPaymentMethod(context.Background(), paymentMethodID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService GetSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionID string

        resp, httpRes, err := apiClient.BillingApi.GetSubscription(context.Background(), subscriptionID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService GetSubscriptionItem", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionItemID string

        resp, httpRes, err := apiClient.BillingApi.GetSubscriptionItem(context.Background(), subscriptionItemID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ListBillingBundles", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.ListBillingBundles(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ListBillingFeatures", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.ListBillingFeatures(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ListBillingPaymentOptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.ListBillingPaymentOptions(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ListBillingProducts", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.ListBillingProducts(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ListPaymentMethods", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.ListPaymentMethods(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ListSubscriptionItems", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.ListSubscriptionItems(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ListSubscriptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.ListSubscriptions(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService ReactivateSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionID string

        resp, httpRes, err := apiClient.BillingApi.ReactivateSubscription(context.Background(), subscriptionID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService UpdateBillingAccount", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BillingApi.UpdateBillingAccount(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BillingApiService UpdateSubscriptionItem", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionItemID string

        resp, httpRes, err := apiClient.BillingApi.UpdateSubscriptionItem(context.Background(), subscriptionItemID).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}

# \CheckoutApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCouponCheckoutSession**](CheckoutApi.md#ApplyCouponCheckoutSession) | **Post** /checkout/sessions/{sessionID}/apply_coupon/ | Apply coupon to session
[**CreateCheckoutSession**](CheckoutApi.md#CreateCheckoutSession) | **Post** /checkout/sessions/ | Create checkout session
[**ExpireCheckoutSession**](CheckoutApi.md#ExpireCheckoutSession) | **Post** /checkout/sessions/{sessionID}/expire/ | Expire checkout session
[**GetCheckoutSession**](CheckoutApi.md#GetCheckoutSession) | **Get** /checkout/sessions/{sessionID}/ | Checkout session details
[**GetStripeCheckoutSession**](CheckoutApi.md#GetStripeCheckoutSession) | **Get** /checkout/sessions/{sessionID}/stripe/checkout/session/ | Stripe checkout session details
[**GetStripePaymentIntent**](CheckoutApi.md#GetStripePaymentIntent) | **Get** /checkout/sessions/{sessionID}/stripe/payment/intent/ | Stripe payment intent details
[**ProcessCloudPaymentsAuth**](CheckoutApi.md#ProcessCloudPaymentsAuth) | **Post** /checkout/sessions/{sessionID}/cloud-payments/auth/ | Cloud payments auth
[**ProcessCloudPaymentsPost3ds**](CheckoutApi.md#ProcessCloudPaymentsPost3ds) | **Post** /checkout/sessions/{sessionID}/cloud-payments/post3ds/ | Cloud payments post3ds



## ApplyCouponCheckoutSession

> CouponCheckoutSession ApplyCouponCheckoutSession(ctx, sessionID).CreateCouponCheckoutSession(createCouponCheckoutSession).Execute()

Apply coupon to session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createCouponCheckoutSession := *openapiclient.NewCreateCouponCheckoutSession("Key_example") // CreateCouponCheckoutSession | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.ApplyCouponCheckoutSession(context.Background(), sessionID).CreateCouponCheckoutSession(createCouponCheckoutSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.ApplyCouponCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyCouponCheckoutSession`: CouponCheckoutSession
    fmt.Fprintf(os.Stdout, "Response from `CheckoutApi.ApplyCouponCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyCouponCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCouponCheckoutSession** | [**CreateCouponCheckoutSession**](CreateCouponCheckoutSession.md) |  | 

### Return type

[**CouponCheckoutSession**](CouponCheckoutSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCheckoutSession

> CheckoutSession CreateCheckoutSession(ctx).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()

Create checkout session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createCheckoutSessionRequest := *openapiclient.NewCreateCheckoutSessionRequest("CancelUrl_example", "SuccessUrl_example", []openapiclient.CreateCheckoutSessionProduct{*openapiclient.NewCreateCheckoutSessionProduct("Product_example", int32(123))}) // CreateCheckoutSessionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.CreateCheckoutSession(context.Background()).CreateCheckoutSessionRequest(createCheckoutSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.CreateCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCheckoutSession`: CheckoutSession
    fmt.Fprintf(os.Stdout, "Response from `CheckoutApi.CreateCheckoutSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheckoutSessionRequest** | [**CreateCheckoutSessionRequest**](CreateCheckoutSessionRequest.md) |  | 

### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpireCheckoutSession

> ExpireCheckoutSession(ctx, sessionID).Execute()

Expire checkout session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.ExpireCheckoutSession(context.Background(), sessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.ExpireCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckoutSession

> CheckoutSession GetCheckoutSession(ctx, sessionID).Execute()

Checkout session details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.GetCheckoutSession(context.Background(), sessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.GetCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckoutSession`: CheckoutSession
    fmt.Fprintf(os.Stdout, "Response from `CheckoutApi.GetCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStripeCheckoutSession

> StripeCheckoutSession GetStripeCheckoutSession(ctx, sessionID).Execute()

Stripe checkout session details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.GetStripeCheckoutSession(context.Background(), sessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.GetStripeCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStripeCheckoutSession`: StripeCheckoutSession
    fmt.Fprintf(os.Stdout, "Response from `CheckoutApi.GetStripeCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStripeCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StripeCheckoutSession**](StripeCheckoutSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStripePaymentIntent

> StripePaymentIntent GetStripePaymentIntent(ctx, sessionID).UseStripeSdk(useStripeSdk).ReturnUrl(returnUrl).Execute()

Stripe payment intent details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    useStripeSdk := true // bool |  (optional)
    returnUrl := "returnUrl_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.GetStripePaymentIntent(context.Background(), sessionID).UseStripeSdk(useStripeSdk).ReturnUrl(returnUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.GetStripePaymentIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStripePaymentIntent`: StripePaymentIntent
    fmt.Fprintf(os.Stdout, "Response from `CheckoutApi.GetStripePaymentIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStripePaymentIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useStripeSdk** | **bool** |  | 
 **returnUrl** | **string** |  | 

### Return type

[**StripePaymentIntent**](StripePaymentIntent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessCloudPaymentsAuth

> CloudPaymentsAuth ProcessCloudPaymentsAuth(ctx, sessionID).CreateCloudPaymentsAuth(createCloudPaymentsAuth).Execute()

Cloud payments auth

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createCloudPaymentsAuth := *openapiclient.NewCreateCloudPaymentsAuth("Cryptogram_example", "Name_example") // CreateCloudPaymentsAuth | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.ProcessCloudPaymentsAuth(context.Background(), sessionID).CreateCloudPaymentsAuth(createCloudPaymentsAuth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.ProcessCloudPaymentsAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessCloudPaymentsAuth`: CloudPaymentsAuth
    fmt.Fprintf(os.Stdout, "Response from `CheckoutApi.ProcessCloudPaymentsAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessCloudPaymentsAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCloudPaymentsAuth** | [**CreateCloudPaymentsAuth**](CreateCloudPaymentsAuth.md) |  | 

### Return type

[**CloudPaymentsAuth**](CloudPaymentsAuth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessCloudPaymentsPost3ds

> CloudPaymentsPost3ds ProcessCloudPaymentsPost3ds(ctx, sessionID).CreateCloudPaymentsPost3ds(createCloudPaymentsPost3ds).Execute()

Cloud payments post3ds

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sessionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createCloudPaymentsPost3ds := *openapiclient.NewCreateCloudPaymentsPost3ds("PaRes_example") // CreateCloudPaymentsPost3ds | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CheckoutApi.ProcessCloudPaymentsPost3ds(context.Background(), sessionID).CreateCloudPaymentsPost3ds(createCloudPaymentsPost3ds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.ProcessCloudPaymentsPost3ds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessCloudPaymentsPost3ds`: CloudPaymentsPost3ds
    fmt.Fprintf(os.Stdout, "Response from `CheckoutApi.ProcessCloudPaymentsPost3ds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessCloudPaymentsPost3dsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCloudPaymentsPost3ds** | [**CreateCloudPaymentsPost3ds**](CreateCloudPaymentsPost3ds.md) |  | 

### Return type

[**CloudPaymentsPost3ds**](CloudPaymentsPost3ds.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


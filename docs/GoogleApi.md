# \GoogleApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyPlayStorePurchase**](GoogleApi.md#VerifyPlayStorePurchase) | **Post** /purchase/google/verify/ | Play store purchase verification



## VerifyPlayStorePurchase

> VerifyPlayStorePurchase(ctx).PlayStorePurchaseVerificationRequest(playStorePurchaseVerificationRequest).Execute()

Play store purchase verification

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
    playStorePurchaseVerificationRequest := *openapiclient.NewPlayStorePurchaseVerificationRequest("ProductSku_example", "PurchaseToken_example") // PlayStorePurchaseVerificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleApi.VerifyPlayStorePurchase(context.Background()).PlayStorePurchaseVerificationRequest(playStorePurchaseVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleApi.VerifyPlayStorePurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPlayStorePurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playStorePurchaseVerificationRequest** | [**PlayStorePurchaseVerificationRequest**](PlayStorePurchaseVerificationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


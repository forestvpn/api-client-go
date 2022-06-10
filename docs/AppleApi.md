# \AppleApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyAppStoreReceipt**](AppleApi.md#VerifyAppStoreReceipt) | **Post** /purchase/apple/verify/ | App store receipt verification



## VerifyAppStoreReceipt

> VerifyAppStoreReceipt(ctx).AppStoreReceiptVerificationRequest(appStoreReceiptVerificationRequest).Execute()

App store receipt verification

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
    appStoreReceiptVerificationRequest := *openapiclient.NewAppStoreReceiptVerificationRequest(string(123)) // AppStoreReceiptVerificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppleApi.VerifyAppStoreReceipt(context.Background()).AppStoreReceiptVerificationRequest(appStoreReceiptVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppleApi.VerifyAppStoreReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAppStoreReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appStoreReceiptVerificationRequest** | [**AppStoreReceiptVerificationRequest**](AppStoreReceiptVerificationRequest.md) |  | 

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


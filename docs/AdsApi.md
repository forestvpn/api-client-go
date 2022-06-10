# \AdsApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdUnitRequestLog**](AdsApi.md#CreateAdUnitRequestLog) | **Post** /ads/request-logs/ | Create ad unit request log
[**ListAdPlacements**](AdsApi.md#ListAdPlacements) | **Get** /ads/placements/ | Get ad placement list



## CreateAdUnitRequestLog

> CreateAdUnitRequestLogRequest CreateAdUnitRequestLog(ctx).CreateAdUnitRequestLogRequest(createAdUnitRequestLogRequest).Execute()

Create ad unit request log

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    createAdUnitRequestLogRequest := *openapiclient.NewCreateAdUnitRequestLogRequest("Unit_example", time.Now(), "PT0.125S - 125ms") // CreateAdUnitRequestLogRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.CreateAdUnitRequestLog(context.Background()).CreateAdUnitRequestLogRequest(createAdUnitRequestLogRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateAdUnitRequestLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdUnitRequestLog`: CreateAdUnitRequestLogRequest
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateAdUnitRequestLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdUnitRequestLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAdUnitRequestLogRequest** | [**CreateAdUnitRequestLogRequest**](CreateAdUnitRequestLogRequest.md) |  | 

### Return type

[**CreateAdUnitRequestLogRequest**](CreateAdUnitRequestLogRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdPlacements

> []AdPlacement ListAdPlacements(ctx).Execute()

Get ad placement list

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.ListAdPlacements(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListAdPlacements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAdPlacements`: []AdPlacement
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListAdPlacements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAdPlacementsRequest struct via the builder pattern


### Return type

[**[]AdPlacement**](AdPlacement.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


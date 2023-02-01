# \AppApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentUserDevice**](AppApi.md#GetCurrentUserDevice) | **Get** /app/devices/current/ | Get user device info
[**UpdateCurrentUserDevice**](AppApi.md#UpdateCurrentUserDevice) | **Patch** /app/devices/current/ | Update user device



## GetCurrentUserDevice

> UserDevice GetCurrentUserDevice(ctx).Execute()

Get user device info

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
    resp, r, err := apiClient.AppApi.GetCurrentUserDevice(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.GetCurrentUserDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUserDevice`: UserDevice
    fmt.Fprintf(os.Stdout, "Response from `AppApi.GetCurrentUserDevice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserDeviceRequest struct via the builder pattern


### Return type

[**UserDevice**](UserDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentUserDevice

> UpdateCurrentUserDevice(ctx).UpdateUserDeviceRequest(updateUserDeviceRequest).Execute()

Update user device

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
    updateUserDeviceRequest := *openapiclient.NewUpdateUserDeviceRequest() // UpdateUserDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppApi.UpdateCurrentUserDevice(context.Background()).UpdateUserDeviceRequest(updateUserDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.UpdateCurrentUserDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentUserDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserDeviceRequest** | [**UpdateUserDeviceRequest**](UpdateUserDeviceRequest.md) |  | 

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


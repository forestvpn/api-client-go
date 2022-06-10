# \FcmApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFCMDevice**](FcmApi.md#CreateFCMDevice) | **Post** /notification/fcm/ | Device registration for push notification through out Firebase Cloud Messaging
[**DeleteFCMDevice**](FcmApi.md#DeleteFCMDevice) | **Delete** /notification/fcm/{registrationID}/ | Delete fcm device
[**GetFCMDevice**](FcmApi.md#GetFCMDevice) | **Get** /notification/fcm/{registrationID}/ | Device info
[**UpdateFCMDevice**](FcmApi.md#UpdateFCMDevice) | **Patch** /notification/fcm/{registrationID}/ | Update device fcm properties



## CreateFCMDevice

> FCMDevice CreateFCMDevice(ctx).CreateFCMDeviceRequest(createFCMDeviceRequest).Execute()

Device registration for push notification through out Firebase Cloud Messaging

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
    createFCMDeviceRequest := *openapiclient.NewCreateFCMDeviceRequest("RegistrationId_example", false) // CreateFCMDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FcmApi.CreateFCMDevice(context.Background()).CreateFCMDeviceRequest(createFCMDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FcmApi.CreateFCMDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFCMDevice`: FCMDevice
    fmt.Fprintf(os.Stdout, "Response from `FcmApi.CreateFCMDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFCMDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFCMDeviceRequest** | [**CreateFCMDeviceRequest**](CreateFCMDeviceRequest.md) |  | 

### Return type

[**FCMDevice**](FCMDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFCMDevice

> DeleteFCMDevice(ctx, registrationID).Execute()

Delete fcm device

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
    registrationID := "registrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FcmApi.DeleteFCMDevice(context.Background(), registrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FcmApi.DeleteFCMDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFCMDeviceRequest struct via the builder pattern


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


## GetFCMDevice

> FCMDevice GetFCMDevice(ctx, registrationID).Execute()

Device info

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
    registrationID := "registrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FcmApi.GetFCMDevice(context.Background(), registrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FcmApi.GetFCMDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFCMDevice`: FCMDevice
    fmt.Fprintf(os.Stdout, "Response from `FcmApi.GetFCMDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFCMDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FCMDevice**](FCMDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFCMDevice

> FCMDevice UpdateFCMDevice(ctx, registrationID).UpdateFCMDeviceRequest(updateFCMDeviceRequest).Execute()

Update device fcm properties

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
    registrationID := "registrationID_example" // string | 
    updateFCMDeviceRequest := *openapiclient.NewUpdateFCMDeviceRequest() // UpdateFCMDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FcmApi.UpdateFCMDevice(context.Background(), registrationID).UpdateFCMDeviceRequest(updateFCMDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FcmApi.UpdateFCMDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFCMDevice`: FCMDevice
    fmt.Fprintf(os.Stdout, "Response from `FcmApi.UpdateFCMDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFCMDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFCMDeviceRequest** | [**UpdateFCMDeviceRequest**](UpdateFCMDeviceRequest.md) |  | 

### Return type

[**FCMDevice**](FCMDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


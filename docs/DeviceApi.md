# \DeviceApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DeviceApi.md#CreateDevice) | **Post** /devices/ | Create new device
[**CreateDevicePortForwarding**](DeviceApi.md#CreateDevicePortForwarding) | **Post** /devices/{deviceID}/port-forwarding/ | Create new device port forwarding
[**DeleteDevice**](DeviceApi.md#DeleteDevice) | **Delete** /devices/{deviceID}/ | Delete Device
[**DeleteDevicePortForwarding**](DeviceApi.md#DeleteDevicePortForwarding) | **Delete** /devices/{deviceID}/port-forwarding/{portForwardingID}/ | Delete Device&#39;s Port Forwarding
[**GetDevice**](DeviceApi.md#GetDevice) | **Get** /devices/{deviceID}/ | Device Info
[**GetDeviceStats**](DeviceApi.md#GetDeviceStats) | **Get** /devices/{deviceID}/stats/{statsID}/ | Device&#39;s stats detail
[**GetDeviceWireGuard**](DeviceApi.md#GetDeviceWireGuard) | **Get** /devices/{deviceID}/wireguards/{wireGuardID}/ | Device&#39;s wireguard profile detail
[**ListDeviceBindings**](DeviceApi.md#ListDeviceBindings) | **Get** /devices/{deviceID}/bindings/ | Device bindings
[**ListDeviceConnectionModes**](DeviceApi.md#ListDeviceConnectionModes) | **Get** /devices/{deviceID}/connection-modes/ | Device connection modes
[**ListDeviceDetailStats**](DeviceApi.md#ListDeviceDetailStats) | **Get** /devices/{deviceID}/detail-stats/ | Device&#39;s detail stats list
[**ListDevicePortForwardings**](DeviceApi.md#ListDevicePortForwardings) | **Get** /devices/{deviceID}/port-forwarding/ | Device Port Forwarding List
[**ListDeviceStats**](DeviceApi.md#ListDeviceStats) | **Get** /devices/{deviceID}/stats/ | Device&#39;s stats list
[**ListDeviceWireGuardPeers**](DeviceApi.md#ListDeviceWireGuardPeers) | **Get** /devices/{deviceID}/wireguards/{wireGuardID}/peers/ | Device&#39;s wireguard peers
[**ListDeviceWireGuards**](DeviceApi.md#ListDeviceWireGuards) | **Get** /devices/{deviceID}/wireguards/ | Device&#39;s wireguard profiles list
[**ListDevices**](DeviceApi.md#ListDevices) | **Get** /devices/ | Device List
[**UpdateDevice**](DeviceApi.md#UpdateDevice) | **Patch** /devices/{deviceID}/ | Update device properties
[**UpdateDevicePortForwarding**](DeviceApi.md#UpdateDevicePortForwarding) | **Patch** /devices/{deviceID}/port-forwarding/{portForwardingID}/ | Update device&#39;s port forwarding



## CreateDevice

> Device CreateDevice(ctx).CreateOrUpdateDeviceRequest(createOrUpdateDeviceRequest).Execute()

Create new device



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
    createOrUpdateDeviceRequest := *openapiclient.NewCreateOrUpdateDeviceRequest() // CreateOrUpdateDeviceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.CreateDevice(context.Background()).CreateOrUpdateDeviceRequest(createOrUpdateDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrUpdateDeviceRequest** | [**CreateOrUpdateDeviceRequest**](CreateOrUpdateDeviceRequest.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevicePortForwarding

> PortForwarding CreateDevicePortForwarding(ctx, deviceID).CreateOrUpdatePortForwardingRequest(createOrUpdatePortForwardingRequest).Execute()

Create new device port forwarding



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createOrUpdatePortForwardingRequest := *openapiclient.NewCreateOrUpdatePortForwardingRequest() // CreateOrUpdatePortForwardingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.CreateDevicePortForwarding(context.Background(), deviceID).CreateOrUpdatePortForwardingRequest(createOrUpdatePortForwardingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDevicePortForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevicePortForwarding`: PortForwarding
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDevicePortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDevicePortForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrUpdatePortForwardingRequest** | [**CreateOrUpdatePortForwardingRequest**](CreateOrUpdatePortForwardingRequest.md) |  | 

### Return type

[**PortForwarding**](PortForwarding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, deviceID).Execute()

Delete Device



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.DeleteDevice(context.Background(), deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## DeleteDevicePortForwarding

> DeleteDevicePortForwarding(ctx, deviceID, portForwardingID).Execute()

Delete Device's Port Forwarding



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    portForwardingID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.DeleteDevicePortForwarding(context.Background(), deviceID, portForwardingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevicePortForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 
**portForwardingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevicePortForwardingRequest struct via the builder pattern


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


## GetDevice

> Device GetDevice(ctx, deviceID).Execute()

Device Info



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDevice(context.Background(), deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceStats

> DeviceStats GetDeviceStats(ctx, deviceID, statsID).Execute()

Device's stats detail

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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    statsID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceStats(context.Background(), deviceID, statsID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceStats`: DeviceStats
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 
**statsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceStats**](DeviceStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWireGuard

> WireGuard GetDeviceWireGuard(ctx, deviceID, wireGuardID).Execute()

Device's wireguard profile detail

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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    wireGuardID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDeviceWireGuard(context.Background(), deviceID, wireGuardID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDeviceWireGuard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWireGuard`: WireGuard
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDeviceWireGuard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 
**wireGuardID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWireGuardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WireGuard**](WireGuard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceBindings

> []string ListDeviceBindings(ctx, deviceID).Execute()

Device bindings



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDeviceBindings(context.Background(), deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDeviceBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceBindings`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDeviceBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceConnectionModes

> []ConnectionMode ListDeviceConnectionModes(ctx, deviceID).XAndroidPackage(xAndroidPackage).XAndroidSHA1(xAndroidSHA1).Execute()

Device connection modes



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    xAndroidPackage := "xAndroidPackage_example" // string |  (optional)
    xAndroidSHA1 := "xAndroidSHA1_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDeviceConnectionModes(context.Background(), deviceID).XAndroidPackage(xAndroidPackage).XAndroidSHA1(xAndroidSHA1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDeviceConnectionModes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceConnectionModes`: []ConnectionMode
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDeviceConnectionModes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceConnectionModesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAndroidPackage** | **string** |  | 
 **xAndroidSHA1** | **string** |  | 

### Return type

[**[]ConnectionMode**](ConnectionMode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceDetailStats

> []DeviceStats ListDeviceDetailStats(ctx, deviceID).DateTimeAfter(dateTimeAfter).DateTimeBefore(dateTimeBefore).PerPage(perPage).Page(page).Execute()

Device's detail stats list

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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dateTimeAfter := "dateTimeAfter_example" // string |  (optional)
    dateTimeBefore := "dateTimeBefore_example" // string |  (optional)
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDeviceDetailStats(context.Background(), deviceID).DateTimeAfter(dateTimeAfter).DateTimeBefore(dateTimeBefore).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDeviceDetailStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceDetailStats`: []DeviceStats
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDeviceDetailStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceDetailStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateTimeAfter** | **string** |  | 
 **dateTimeBefore** | **string** |  | 
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**[]DeviceStats**](DeviceStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicePortForwardings

> []PortForwarding ListDevicePortForwardings(ctx, deviceID).PerPage(perPage).Page(page).Execute()

Device Port Forwarding List



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDevicePortForwardings(context.Background(), deviceID).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDevicePortForwardings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevicePortForwardings`: []PortForwarding
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDevicePortForwardings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDevicePortForwardingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**[]PortForwarding**](PortForwarding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceStats

> []DeviceStats ListDeviceStats(ctx, deviceID).DateAfter(dateAfter).DateBefore(dateBefore).PerPage(perPage).Page(page).Execute()

Device's stats list

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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dateAfter := time.Now() // string |  (optional)
    dateBefore := time.Now() // string |  (optional)
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDeviceStats(context.Background(), deviceID).DateAfter(dateAfter).DateBefore(dateBefore).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDeviceStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceStats`: []DeviceStats
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDeviceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateAfter** | **string** |  | 
 **dateBefore** | **string** |  | 
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**[]DeviceStats**](DeviceStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceWireGuardPeers

> []WireGuardPeer ListDeviceWireGuardPeers(ctx, deviceID, wireGuardID).Execute()

Device's wireguard peers

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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    wireGuardID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDeviceWireGuardPeers(context.Background(), deviceID, wireGuardID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDeviceWireGuardPeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceWireGuardPeers`: []WireGuardPeer
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDeviceWireGuardPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 
**wireGuardID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceWireGuardPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WireGuardPeer**](WireGuardPeer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceWireGuards

> []WireGuard ListDeviceWireGuards(ctx, deviceID).PerPage(perPage).Page(page).Execute()

Device's wireguard profiles list

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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDeviceWireGuards(context.Background(), deviceID).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDeviceWireGuards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceWireGuards`: []WireGuard
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDeviceWireGuards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceWireGuardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**[]WireGuard**](WireGuard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevices

> []Device ListDevices(ctx).Q(q).ExternalKey(externalKey).RecentlyActive(recentlyActive).LastActiveAtAfter(lastActiveAtAfter).LastActiveAtBefore(lastActiveAtBefore).Sort(sort).PerPage(perPage).Page(page).Execute()

Device List



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
    q := "q_example" // string | Filter by search query (optional)
    externalKey := "externalKey_example" // string | Filter by external_key (optional)
    recentlyActive := true // bool | Filter by recently active (optional)
    lastActiveAtAfter := time.Now() // time.Time | Filter by last active at date-time after provided value (optional)
    lastActiveAtBefore := time.Now() // time.Time | Filter by last active at date-time before provided value (optional)
    sort := "sort_example" // string | Sort by provided field (optional)
    perPage := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDevices(context.Background()).Q(q).ExternalKey(externalKey).RecentlyActive(recentlyActive).LastActiveAtAfter(lastActiveAtAfter).LastActiveAtBefore(lastActiveAtBefore).Sort(sort).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevices`: []Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Filter by search query | 
 **externalKey** | **string** | Filter by external_key | 
 **recentlyActive** | **bool** | Filter by recently active | 
 **lastActiveAtAfter** | **time.Time** | Filter by last active at date-time after provided value | 
 **lastActiveAtBefore** | **time.Time** | Filter by last active at date-time before provided value | 
 **sort** | **string** | Sort by provided field | 
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

[**[]Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> Device UpdateDevice(ctx, deviceID).CreateOrUpdateDeviceRequest(createOrUpdateDeviceRequest).Execute()

Update device properties



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createOrUpdateDeviceRequest := *openapiclient.NewCreateOrUpdateDeviceRequest() // CreateOrUpdateDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UpdateDevice(context.Background(), deviceID).CreateOrUpdateDeviceRequest(createOrUpdateDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrUpdateDeviceRequest** | [**CreateOrUpdateDeviceRequest**](CreateOrUpdateDeviceRequest.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevicePortForwarding

> PortForwarding UpdateDevicePortForwarding(ctx, deviceID, portForwardingID).CreateOrUpdatePortForwardingRequest(createOrUpdatePortForwardingRequest).Execute()

Update device's port forwarding



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
    deviceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    portForwardingID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createOrUpdatePortForwardingRequest := *openapiclient.NewCreateOrUpdatePortForwardingRequest() // CreateOrUpdatePortForwardingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UpdateDevicePortForwarding(context.Background(), deviceID, portForwardingID).CreateOrUpdatePortForwardingRequest(createOrUpdatePortForwardingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateDevicePortForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevicePortForwarding`: PortForwarding
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateDevicePortForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceID** | **string** |  | 
**portForwardingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDevicePortForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOrUpdatePortForwardingRequest** | [**CreateOrUpdatePortForwardingRequest**](CreateOrUpdatePortForwardingRequest.md) |  | 

### Return type

[**PortForwarding**](PortForwarding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


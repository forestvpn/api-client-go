# \WireguardApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWireGuardPeerInfo**](WireguardApi.md#GetWireGuardPeerInfo) | **Get** /wireguard/peers/{pubKey}/ | Wireguard peer info
[**ListWireGuardPeers**](WireguardApi.md#ListWireGuardPeers) | **Get** /wireguard/peers/ | Wireguard peers list



## GetWireGuardPeerInfo

> WireGuardPeerInfo GetWireGuardPeerInfo(ctx, pubKey).Execute()

Wireguard peer info

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
    pubKey := "pubKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WireguardApi.GetWireGuardPeerInfo(context.Background(), pubKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardApi.GetWireGuardPeerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWireGuardPeerInfo`: WireGuardPeerInfo
    fmt.Fprintf(os.Stdout, "Response from `WireguardApi.GetWireGuardPeerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireGuardPeerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WireGuardPeerInfo**](WireGuardPeerInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWireGuardPeers

> []WireGuardPeerInfo ListWireGuardPeers(ctx).Execute()

Wireguard peers list

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
    resp, r, err := apiClient.WireguardApi.ListWireGuardPeers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardApi.ListWireGuardPeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWireGuardPeers`: []WireGuardPeerInfo
    fmt.Fprintf(os.Stdout, "Response from `WireguardApi.ListWireGuardPeers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWireGuardPeersRequest struct via the builder pattern


### Return type

[**[]WireGuardPeerInfo**](WireGuardPeerInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


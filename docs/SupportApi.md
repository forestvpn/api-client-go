# \SupportApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportTicket**](SupportApi.md#CreateSupportTicket) | **Post** /support/tickets/ | Create support ticket
[**GetSupportTicketCategory**](SupportApi.md#GetSupportTicketCategory) | **Get** /support/ticket-categories/ | Get ticket categories



## CreateSupportTicket

> CreateSupportTicket(ctx).Text(text).Category(category).Files(files).Execute()

Create support ticket

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
    text := "text_example" // string | 
    category := "category_example" // string | Ticket category's slug
    files := []*os.File{"TODO"} // []*os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SupportApi.CreateSupportTicket(context.Background()).Text(text).Category(category).Files(files).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.CreateSupportTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** |  | 
 **category** | **string** | Ticket category&#39;s slug | 
 **files** | **[]*os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicketCategory

> []TicketCategory GetSupportTicketCategory(ctx).Execute()

Get ticket categories

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
    resp, r, err := apiClient.SupportApi.GetSupportTicketCategory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportApi.GetSupportTicketCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportTicketCategory`: []TicketCategory
    fmt.Fprintf(os.Stdout, "Response from `SupportApi.GetSupportTicketCategory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketCategoryRequest struct via the builder pattern


### Return type

[**[]TicketCategory**](TicketCategory.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


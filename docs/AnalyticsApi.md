# \AnalyticsApi

All URIs are relative to *https://api.forestvpn.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataUsageStats**](AnalyticsApi.md#GetDataUsageStats) | **Get** /analytics/data-usage/ | Data Usage Stats



## GetDataUsageStats

> []AggregatedDataUsageStats GetDataUsageStats(ctx).AggrInterval(aggrInterval).DateAfter(dateAfter).DateBefore(dateBefore).DeviceTypeIn(deviceTypeIn).DeviceIn(deviceIn).SortBy(sortBy).Execute()

Data Usage Stats



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
    aggrInterval := "hour / day / week / month" // string | No more than 3 days for hour aggregation type
    dateAfter := time.Now() // string | No more than 3 months between date_after and date_before (optional)
    dateBefore := time.Now() // string | No more than 3 months between date_after and date_before (optional)
    deviceTypeIn := []string{"Inner_example"} // []string | Filter by device type (optional)
    deviceIn := []string{"Inner_example"} // []string | Filter by device (optional)
    sortBy := "sortBy_example" // string | Sort output by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetDataUsageStats(context.Background()).AggrInterval(aggrInterval).DateAfter(dateAfter).DateBefore(dateBefore).DeviceTypeIn(deviceTypeIn).DeviceIn(deviceIn).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetDataUsageStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataUsageStats`: []AggregatedDataUsageStats
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetDataUsageStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataUsageStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aggrInterval** | **string** | No more than 3 days for hour aggregation type | 
 **dateAfter** | **string** | No more than 3 months between date_after and date_before | 
 **dateBefore** | **string** | No more than 3 months between date_after and date_before | 
 **deviceTypeIn** | **[]string** | Filter by device type | 
 **deviceIn** | **[]string** | Filter by device | 
 **sortBy** | **string** | Sort output by | 

### Return type

[**[]AggregatedDataUsageStats**](AggregatedDataUsageStats.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


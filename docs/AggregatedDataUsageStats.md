# AggregatedDataUsageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggrInterval** | Pointer to **string** | Aggregation value. It might de a hour, day, week, or month | [optional] 
**DeviceId** | Pointer to **string** | User device name, useful for retrieve extra data through device API | [optional] 
**DeviceName** | Pointer to **string** | User device name, useful for showing in the chart | [optional] 
**ReceivedBytes** | Pointer to **int32** | Sum of the received bytes aggregated around the device and aggr_interval | [optional] 
**TransmittedBytes** | Pointer to **int32** | Sum of the transmitted bytes aggregated around the device and aggr_interval | [optional] 
**TotalBytes** | Pointer to **int32** | Sum of the received + transmitted bytes aggregated around the device and aggr_interval | [optional] 
**UsageTime** | Pointer to **int32** | Minutes of usage time aggregated around the device and aggr_interval | [optional] 

## Methods

### NewAggregatedDataUsageStats

`func NewAggregatedDataUsageStats() *AggregatedDataUsageStats`

NewAggregatedDataUsageStats instantiates a new AggregatedDataUsageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedDataUsageStatsWithDefaults

`func NewAggregatedDataUsageStatsWithDefaults() *AggregatedDataUsageStats`

NewAggregatedDataUsageStatsWithDefaults instantiates a new AggregatedDataUsageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggrInterval

`func (o *AggregatedDataUsageStats) GetAggrInterval() string`

GetAggrInterval returns the AggrInterval field if non-nil, zero value otherwise.

### GetAggrIntervalOk

`func (o *AggregatedDataUsageStats) GetAggrIntervalOk() (*string, bool)`

GetAggrIntervalOk returns a tuple with the AggrInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrInterval

`func (o *AggregatedDataUsageStats) SetAggrInterval(v string)`

SetAggrInterval sets AggrInterval field to given value.

### HasAggrInterval

`func (o *AggregatedDataUsageStats) HasAggrInterval() bool`

HasAggrInterval returns a boolean if a field has been set.

### GetDeviceId

`func (o *AggregatedDataUsageStats) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AggregatedDataUsageStats) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AggregatedDataUsageStats) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AggregatedDataUsageStats) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *AggregatedDataUsageStats) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AggregatedDataUsageStats) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AggregatedDataUsageStats) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AggregatedDataUsageStats) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetReceivedBytes

`func (o *AggregatedDataUsageStats) GetReceivedBytes() int32`

GetReceivedBytes returns the ReceivedBytes field if non-nil, zero value otherwise.

### GetReceivedBytesOk

`func (o *AggregatedDataUsageStats) GetReceivedBytesOk() (*int32, bool)`

GetReceivedBytesOk returns a tuple with the ReceivedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBytes

`func (o *AggregatedDataUsageStats) SetReceivedBytes(v int32)`

SetReceivedBytes sets ReceivedBytes field to given value.

### HasReceivedBytes

`func (o *AggregatedDataUsageStats) HasReceivedBytes() bool`

HasReceivedBytes returns a boolean if a field has been set.

### GetTransmittedBytes

`func (o *AggregatedDataUsageStats) GetTransmittedBytes() int32`

GetTransmittedBytes returns the TransmittedBytes field if non-nil, zero value otherwise.

### GetTransmittedBytesOk

`func (o *AggregatedDataUsageStats) GetTransmittedBytesOk() (*int32, bool)`

GetTransmittedBytesOk returns a tuple with the TransmittedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmittedBytes

`func (o *AggregatedDataUsageStats) SetTransmittedBytes(v int32)`

SetTransmittedBytes sets TransmittedBytes field to given value.

### HasTransmittedBytes

`func (o *AggregatedDataUsageStats) HasTransmittedBytes() bool`

HasTransmittedBytes returns a boolean if a field has been set.

### GetTotalBytes

`func (o *AggregatedDataUsageStats) GetTotalBytes() int32`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *AggregatedDataUsageStats) GetTotalBytesOk() (*int32, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *AggregatedDataUsageStats) SetTotalBytes(v int32)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *AggregatedDataUsageStats) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.

### GetUsageTime

`func (o *AggregatedDataUsageStats) GetUsageTime() int32`

GetUsageTime returns the UsageTime field if non-nil, zero value otherwise.

### GetUsageTimeOk

`func (o *AggregatedDataUsageStats) GetUsageTimeOk() (*int32, bool)`

GetUsageTimeOk returns a tuple with the UsageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTime

`func (o *AggregatedDataUsageStats) SetUsageTime(v int32)`

SetUsageTime sets UsageTime field to given value.

### HasUsageTime

`func (o *AggregatedDataUsageStats) HasUsageTime() bool`

HasUsageTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DeviceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**ReceivedBytes** | Pointer to **float32** |  | [optional] 
**TransmittedBytes** | Pointer to **float32** |  | [optional] 
**UsageTime** | Pointer to **float32** |  | [optional] 

## Methods

### NewDeviceRecord

`func NewDeviceRecord() *DeviceRecord`

NewDeviceRecord instantiates a new DeviceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRecordWithDefaults

`func NewDeviceRecordWithDefaults() *DeviceRecord`

NewDeviceRecordWithDefaults instantiates a new DeviceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *DeviceRecord) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DeviceRecord) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DeviceRecord) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *DeviceRecord) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetReceivedBytes

`func (o *DeviceRecord) GetReceivedBytes() float32`

GetReceivedBytes returns the ReceivedBytes field if non-nil, zero value otherwise.

### GetReceivedBytesOk

`func (o *DeviceRecord) GetReceivedBytesOk() (*float32, bool)`

GetReceivedBytesOk returns a tuple with the ReceivedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBytes

`func (o *DeviceRecord) SetReceivedBytes(v float32)`

SetReceivedBytes sets ReceivedBytes field to given value.

### HasReceivedBytes

`func (o *DeviceRecord) HasReceivedBytes() bool`

HasReceivedBytes returns a boolean if a field has been set.

### GetTransmittedBytes

`func (o *DeviceRecord) GetTransmittedBytes() float32`

GetTransmittedBytes returns the TransmittedBytes field if non-nil, zero value otherwise.

### GetTransmittedBytesOk

`func (o *DeviceRecord) GetTransmittedBytesOk() (*float32, bool)`

GetTransmittedBytesOk returns a tuple with the TransmittedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmittedBytes

`func (o *DeviceRecord) SetTransmittedBytes(v float32)`

SetTransmittedBytes sets TransmittedBytes field to given value.

### HasTransmittedBytes

`func (o *DeviceRecord) HasTransmittedBytes() bool`

HasTransmittedBytes returns a boolean if a field has been set.

### GetUsageTime

`func (o *DeviceRecord) GetUsageTime() float32`

GetUsageTime returns the UsageTime field if non-nil, zero value otherwise.

### GetUsageTimeOk

`func (o *DeviceRecord) GetUsageTimeOk() (*float32, bool)`

GetUsageTimeOk returns a tuple with the UsageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageTime

`func (o *DeviceRecord) SetUsageTime(v float32)`

SetUsageTime sets UsageTime field to given value.

### HasUsageTime

`func (o *DeviceRecord) HasUsageTime() bool`

HasUsageTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DeviceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Connections** | Pointer to **int32** |  | [optional] 
**ReceivedBytes** | Pointer to **int32** |  | [optional] 
**TransmittedBytes** | Pointer to **int32** |  | [optional] 
**BlockedAds** | Pointer to **int32** |  | [optional] 
**BlockedMalwares** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeviceStats

`func NewDeviceStats() *DeviceStats`

NewDeviceStats instantiates a new DeviceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStatsWithDefaults

`func NewDeviceStatsWithDefaults() *DeviceStats`

NewDeviceStatsWithDefaults instantiates a new DeviceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceStats) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConnections

`func (o *DeviceStats) GetConnections() int32`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *DeviceStats) GetConnectionsOk() (*int32, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *DeviceStats) SetConnections(v int32)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *DeviceStats) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetReceivedBytes

`func (o *DeviceStats) GetReceivedBytes() int32`

GetReceivedBytes returns the ReceivedBytes field if non-nil, zero value otherwise.

### GetReceivedBytesOk

`func (o *DeviceStats) GetReceivedBytesOk() (*int32, bool)`

GetReceivedBytesOk returns a tuple with the ReceivedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBytes

`func (o *DeviceStats) SetReceivedBytes(v int32)`

SetReceivedBytes sets ReceivedBytes field to given value.

### HasReceivedBytes

`func (o *DeviceStats) HasReceivedBytes() bool`

HasReceivedBytes returns a boolean if a field has been set.

### GetTransmittedBytes

`func (o *DeviceStats) GetTransmittedBytes() int32`

GetTransmittedBytes returns the TransmittedBytes field if non-nil, zero value otherwise.

### GetTransmittedBytesOk

`func (o *DeviceStats) GetTransmittedBytesOk() (*int32, bool)`

GetTransmittedBytesOk returns a tuple with the TransmittedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmittedBytes

`func (o *DeviceStats) SetTransmittedBytes(v int32)`

SetTransmittedBytes sets TransmittedBytes field to given value.

### HasTransmittedBytes

`func (o *DeviceStats) HasTransmittedBytes() bool`

HasTransmittedBytes returns a boolean if a field has been set.

### GetBlockedAds

`func (o *DeviceStats) GetBlockedAds() int32`

GetBlockedAds returns the BlockedAds field if non-nil, zero value otherwise.

### GetBlockedAdsOk

`func (o *DeviceStats) GetBlockedAdsOk() (*int32, bool)`

GetBlockedAdsOk returns a tuple with the BlockedAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedAds

`func (o *DeviceStats) SetBlockedAds(v int32)`

SetBlockedAds sets BlockedAds field to given value.

### HasBlockedAds

`func (o *DeviceStats) HasBlockedAds() bool`

HasBlockedAds returns a boolean if a field has been set.

### GetBlockedMalwares

`func (o *DeviceStats) GetBlockedMalwares() int32`

GetBlockedMalwares returns the BlockedMalwares field if non-nil, zero value otherwise.

### GetBlockedMalwaresOk

`func (o *DeviceStats) GetBlockedMalwaresOk() (*int32, bool)`

GetBlockedMalwaresOk returns a tuple with the BlockedMalwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMalwares

`func (o *DeviceStats) SetBlockedMalwares(v int32)`

SetBlockedMalwares sets BlockedMalwares field to given value.

### HasBlockedMalwares

`func (o *DeviceStats) HasBlockedMalwares() bool`

HasBlockedMalwares returns a boolean if a field has been set.

### GetDate

`func (o *DeviceStats) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DeviceStats) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DeviceStats) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *DeviceStats) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



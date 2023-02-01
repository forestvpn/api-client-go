# CreateOrUpdateDeviceRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DeviceType**](DeviceType.md) |  | 
**Info** | **map[string]string** |  | 

## Methods

### NewCreateOrUpdateDeviceRequestInfo

`func NewCreateOrUpdateDeviceRequestInfo(type_ DeviceType, info map[string]string, ) *CreateOrUpdateDeviceRequestInfo`

NewCreateOrUpdateDeviceRequestInfo instantiates a new CreateOrUpdateDeviceRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateDeviceRequestInfoWithDefaults

`func NewCreateOrUpdateDeviceRequestInfoWithDefaults() *CreateOrUpdateDeviceRequestInfo`

NewCreateOrUpdateDeviceRequestInfoWithDefaults instantiates a new CreateOrUpdateDeviceRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateOrUpdateDeviceRequestInfo) GetType() DeviceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrUpdateDeviceRequestInfo) GetTypeOk() (*DeviceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrUpdateDeviceRequestInfo) SetType(v DeviceType)`

SetType sets Type field to given value.


### GetInfo

`func (o *CreateOrUpdateDeviceRequestInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CreateOrUpdateDeviceRequestInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CreateOrUpdateDeviceRequestInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



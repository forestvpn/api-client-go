# CreateOrUpdateDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalKey** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**TorOver** | Pointer to **bool** |  | [optional] [default to false]
**ConnectionMode** | Pointer to **NullableString** |  | [optional] 
**RandomServer** | Pointer to **bool** |  | [optional] [default to false]
**Info** | Pointer to [**NullableCreateOrUpdateDeviceRequestInfo**](CreateOrUpdateDeviceRequestInfo.md) |  | [optional] 

## Methods

### NewCreateOrUpdateDeviceRequest

`func NewCreateOrUpdateDeviceRequest() *CreateOrUpdateDeviceRequest`

NewCreateOrUpdateDeviceRequest instantiates a new CreateOrUpdateDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateDeviceRequestWithDefaults

`func NewCreateOrUpdateDeviceRequestWithDefaults() *CreateOrUpdateDeviceRequest`

NewCreateOrUpdateDeviceRequestWithDefaults instantiates a new CreateOrUpdateDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalKey

`func (o *CreateOrUpdateDeviceRequest) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *CreateOrUpdateDeviceRequest) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *CreateOrUpdateDeviceRequest) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *CreateOrUpdateDeviceRequest) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### SetExternalKeyNil

`func (o *CreateOrUpdateDeviceRequest) SetExternalKeyNil(b bool)`

 SetExternalKeyNil sets the value for ExternalKey to be an explicit nil

### UnsetExternalKey
`func (o *CreateOrUpdateDeviceRequest) UnsetExternalKey()`

UnsetExternalKey ensures that no value is present for ExternalKey, not even an explicit nil
### GetName

`func (o *CreateOrUpdateDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateDeviceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrUpdateDeviceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *CreateOrUpdateDeviceRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateOrUpdateDeviceRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateOrUpdateDeviceRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateOrUpdateDeviceRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *CreateOrUpdateDeviceRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *CreateOrUpdateDeviceRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetTorOver

`func (o *CreateOrUpdateDeviceRequest) GetTorOver() bool`

GetTorOver returns the TorOver field if non-nil, zero value otherwise.

### GetTorOverOk

`func (o *CreateOrUpdateDeviceRequest) GetTorOverOk() (*bool, bool)`

GetTorOverOk returns a tuple with the TorOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorOver

`func (o *CreateOrUpdateDeviceRequest) SetTorOver(v bool)`

SetTorOver sets TorOver field to given value.

### HasTorOver

`func (o *CreateOrUpdateDeviceRequest) HasTorOver() bool`

HasTorOver returns a boolean if a field has been set.

### GetConnectionMode

`func (o *CreateOrUpdateDeviceRequest) GetConnectionMode() string`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *CreateOrUpdateDeviceRequest) GetConnectionModeOk() (*string, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *CreateOrUpdateDeviceRequest) SetConnectionMode(v string)`

SetConnectionMode sets ConnectionMode field to given value.

### HasConnectionMode

`func (o *CreateOrUpdateDeviceRequest) HasConnectionMode() bool`

HasConnectionMode returns a boolean if a field has been set.

### SetConnectionModeNil

`func (o *CreateOrUpdateDeviceRequest) SetConnectionModeNil(b bool)`

 SetConnectionModeNil sets the value for ConnectionMode to be an explicit nil

### UnsetConnectionMode
`func (o *CreateOrUpdateDeviceRequest) UnsetConnectionMode()`

UnsetConnectionMode ensures that no value is present for ConnectionMode, not even an explicit nil
### GetRandomServer

`func (o *CreateOrUpdateDeviceRequest) GetRandomServer() bool`

GetRandomServer returns the RandomServer field if non-nil, zero value otherwise.

### GetRandomServerOk

`func (o *CreateOrUpdateDeviceRequest) GetRandomServerOk() (*bool, bool)`

GetRandomServerOk returns a tuple with the RandomServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomServer

`func (o *CreateOrUpdateDeviceRequest) SetRandomServer(v bool)`

SetRandomServer sets RandomServer field to given value.

### HasRandomServer

`func (o *CreateOrUpdateDeviceRequest) HasRandomServer() bool`

HasRandomServer returns a boolean if a field has been set.

### GetInfo

`func (o *CreateOrUpdateDeviceRequest) GetInfo() CreateOrUpdateDeviceRequestInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CreateOrUpdateDeviceRequest) GetInfoOk() (*CreateOrUpdateDeviceRequestInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CreateOrUpdateDeviceRequest) SetInfo(v CreateOrUpdateDeviceRequestInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CreateOrUpdateDeviceRequest) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *CreateOrUpdateDeviceRequest) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *CreateOrUpdateDeviceRequest) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



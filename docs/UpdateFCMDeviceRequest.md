# UpdateFCMDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateFCMDeviceRequest

`func NewUpdateFCMDeviceRequest() *UpdateFCMDeviceRequest`

NewUpdateFCMDeviceRequest instantiates a new UpdateFCMDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFCMDeviceRequestWithDefaults

`func NewUpdateFCMDeviceRequestWithDefaults() *UpdateFCMDeviceRequest`

NewUpdateFCMDeviceRequestWithDefaults instantiates a new UpdateFCMDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *UpdateFCMDeviceRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UpdateFCMDeviceRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UpdateFCMDeviceRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *UpdateFCMDeviceRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *UpdateFCMDeviceRequest) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *UpdateFCMDeviceRequest) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetActive

`func (o *UpdateFCMDeviceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateFCMDeviceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateFCMDeviceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateFCMDeviceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *UpdateFCMDeviceRequest) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *UpdateFCMDeviceRequest) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateFCMDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationId** | **string** |  | 
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**Active** | **bool** |  | 

## Methods

### NewCreateFCMDeviceRequest

`func NewCreateFCMDeviceRequest(registrationId string, active bool, ) *CreateFCMDeviceRequest`

NewCreateFCMDeviceRequest instantiates a new CreateFCMDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFCMDeviceRequestWithDefaults

`func NewCreateFCMDeviceRequestWithDefaults() *CreateFCMDeviceRequest`

NewCreateFCMDeviceRequestWithDefaults instantiates a new CreateFCMDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationId

`func (o *CreateFCMDeviceRequest) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *CreateFCMDeviceRequest) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *CreateFCMDeviceRequest) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.


### GetDeviceId

`func (o *CreateFCMDeviceRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CreateFCMDeviceRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CreateFCMDeviceRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *CreateFCMDeviceRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *CreateFCMDeviceRequest) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *CreateFCMDeviceRequest) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetActive

`func (o *CreateFCMDeviceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateFCMDeviceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateFCMDeviceRequest) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FCMDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationId** | **string** |  | 
**DeviceId** | Pointer to **string** |  | [optional] 
**Active** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewFCMDevice

`func NewFCMDevice(registrationId string, active bool, createdAt time.Time, ) *FCMDevice`

NewFCMDevice instantiates a new FCMDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCMDeviceWithDefaults

`func NewFCMDeviceWithDefaults() *FCMDevice`

NewFCMDeviceWithDefaults instantiates a new FCMDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationId

`func (o *FCMDevice) GetRegistrationId() string`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *FCMDevice) GetRegistrationIdOk() (*string, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *FCMDevice) SetRegistrationId(v string)`

SetRegistrationId sets RegistrationId field to given value.


### GetDeviceId

`func (o *FCMDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *FCMDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *FCMDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *FCMDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetActive

`func (o *FCMDevice) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FCMDevice) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FCMDevice) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *FCMDevice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FCMDevice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FCMDevice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



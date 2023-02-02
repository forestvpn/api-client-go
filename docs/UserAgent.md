# UserAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | Pointer to [**UserAgentOs**](UserAgentOs.md) |  | [optional] 
**Device** | Pointer to [**UserAgentDevice**](UserAgentDevice.md) |  | [optional] 
**Browser** | Pointer to [**UserAgentBrowser**](UserAgentBrowser.md) |  | [optional] 

## Methods

### NewUserAgent

`func NewUserAgent() *UserAgent`

NewUserAgent instantiates a new UserAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentWithDefaults

`func NewUserAgentWithDefaults() *UserAgent`

NewUserAgentWithDefaults instantiates a new UserAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *UserAgent) GetOs() UserAgentOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *UserAgent) GetOsOk() (*UserAgentOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *UserAgent) SetOs(v UserAgentOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *UserAgent) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetDevice

`func (o *UserAgent) GetDevice() UserAgentDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *UserAgent) GetDeviceOk() (*UserAgentDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *UserAgent) SetDevice(v UserAgentDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *UserAgent) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetBrowser

`func (o *UserAgent) GetBrowser() UserAgentBrowser`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *UserAgent) GetBrowserOk() (*UserAgentBrowser, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *UserAgent) SetBrowser(v UserAgentBrowser)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *UserAgent) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



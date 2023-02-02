# UserAgentOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | **string** |  | 
**Version** | Pointer to **string** | It might be empty string | [optional] 

## Methods

### NewUserAgentOs

`func NewUserAgentOs(family string, ) *UserAgentOs`

NewUserAgentOs instantiates a new UserAgentOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentOsWithDefaults

`func NewUserAgentOsWithDefaults() *UserAgentOs`

NewUserAgentOsWithDefaults instantiates a new UserAgentOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *UserAgentOs) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *UserAgentOs) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *UserAgentOs) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetVersion

`func (o *UserAgentOs) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserAgentOs) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserAgentOs) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserAgentOs) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



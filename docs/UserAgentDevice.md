# UserAgentDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | **string** | It might be \&quot;Other\&quot; in case if it can&#39;t be recognized | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewUserAgentDevice

`func NewUserAgentDevice(family string, type_ string, ) *UserAgentDevice`

NewUserAgentDevice instantiates a new UserAgentDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentDeviceWithDefaults

`func NewUserAgentDeviceWithDefaults() *UserAgentDevice`

NewUserAgentDeviceWithDefaults instantiates a new UserAgentDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *UserAgentDevice) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *UserAgentDevice) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *UserAgentDevice) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetBrand

`func (o *UserAgentDevice) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UserAgentDevice) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UserAgentDevice) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *UserAgentDevice) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *UserAgentDevice) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *UserAgentDevice) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetModel

`func (o *UserAgentDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UserAgentDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UserAgentDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *UserAgentDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *UserAgentDevice) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *UserAgentDevice) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetType

`func (o *UserAgentDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAgentDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAgentDevice) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



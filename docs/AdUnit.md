# AdUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ExternalKey** | Pointer to **string** |  | [optional] 
**Format** | **string** |  | 
**Reward** | Pointer to [**AdReward**](AdReward.md) |  | [optional] 
**Provider** | [**AdProvider**](AdProvider.md) |  | 

## Methods

### NewAdUnit

`func NewAdUnit(id string, name string, format string, provider AdProvider, ) *AdUnit`

NewAdUnit instantiates a new AdUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdUnitWithDefaults

`func NewAdUnitWithDefaults() *AdUnit`

NewAdUnitWithDefaults instantiates a new AdUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdUnit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdUnit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdUnit) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AdUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdUnit) SetName(v string)`

SetName sets Name field to given value.


### GetExternalKey

`func (o *AdUnit) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *AdUnit) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *AdUnit) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *AdUnit) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### GetFormat

`func (o *AdUnit) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AdUnit) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AdUnit) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetReward

`func (o *AdUnit) GetReward() AdReward`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *AdUnit) GetRewardOk() (*AdReward, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *AdUnit) SetReward(v AdReward)`

SetReward sets Reward field to given value.

### HasReward

`func (o *AdUnit) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetProvider

`func (o *AdUnit) GetProvider() AdProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AdUnit) GetProviderOk() (*AdProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AdUnit) SetProvider(v AdProvider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



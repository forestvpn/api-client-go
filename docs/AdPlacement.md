# AdPlacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Slug** | **string** |  | 
**Units** | [**[]AdUnit**](AdUnit.md) |  | 

## Methods

### NewAdPlacement

`func NewAdPlacement(id string, slug string, units []AdUnit, ) *AdPlacement`

NewAdPlacement instantiates a new AdPlacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdPlacementWithDefaults

`func NewAdPlacementWithDefaults() *AdPlacement`

NewAdPlacementWithDefaults instantiates a new AdPlacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdPlacement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdPlacement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdPlacement) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *AdPlacement) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AdPlacement) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AdPlacement) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUnits

`func (o *AdPlacement) GetUnits() []AdUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *AdPlacement) GetUnitsOk() (*[]AdUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *AdPlacement) SetUnits(v []AdUnit)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



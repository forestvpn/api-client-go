# ProductWithoutPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Bundle** | Pointer to [**Bundle**](Bundle.md) |  | [optional] 
**Recurring** | Pointer to [**Recurring**](Recurring.md) |  | [optional] 
**Discount** | Pointer to [**Discount**](Discount.md) |  | [optional] 
**IsMostPopular** | Pointer to **bool** |  | [optional] 

## Methods

### NewProductWithoutPrice

`func NewProductWithoutPrice(id string, name string, ) *ProductWithoutPrice`

NewProductWithoutPrice instantiates a new ProductWithoutPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithoutPriceWithDefaults

`func NewProductWithoutPriceWithDefaults() *ProductWithoutPrice`

NewProductWithoutPriceWithDefaults instantiates a new ProductWithoutPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductWithoutPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductWithoutPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductWithoutPrice) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProductWithoutPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductWithoutPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductWithoutPrice) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductWithoutPrice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductWithoutPrice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductWithoutPrice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductWithoutPrice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBundle

`func (o *ProductWithoutPrice) GetBundle() Bundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *ProductWithoutPrice) GetBundleOk() (*Bundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *ProductWithoutPrice) SetBundle(v Bundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *ProductWithoutPrice) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetRecurring

`func (o *ProductWithoutPrice) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *ProductWithoutPrice) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *ProductWithoutPrice) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *ProductWithoutPrice) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetDiscount

`func (o *ProductWithoutPrice) GetDiscount() Discount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ProductWithoutPrice) GetDiscountOk() (*Discount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ProductWithoutPrice) SetDiscount(v Discount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ProductWithoutPrice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetIsMostPopular

`func (o *ProductWithoutPrice) GetIsMostPopular() bool`

GetIsMostPopular returns the IsMostPopular field if non-nil, zero value otherwise.

### GetIsMostPopularOk

`func (o *ProductWithoutPrice) GetIsMostPopularOk() (*bool, bool)`

GetIsMostPopularOk returns a tuple with the IsMostPopular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMostPopular

`func (o *ProductWithoutPrice) SetIsMostPopular(v bool)`

SetIsMostPopular sets IsMostPopular field to given value.

### HasIsMostPopular

`func (o *ProductWithoutPrice) HasIsMostPopular() bool`

HasIsMostPopular returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



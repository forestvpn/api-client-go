# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Bundle** | Pointer to [**Bundle**](Bundle.md) |  | [optional] 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 
**Recurring** | Pointer to [**Recurring**](Recurring.md) |  | [optional] 
**Discount** | Pointer to [**Discount**](Discount.md) |  | [optional] 
**IsMostPopular** | Pointer to **bool** |  | [optional] 

## Methods

### NewProduct

`func NewProduct(id string, name string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBundle

`func (o *Product) GetBundle() Bundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *Product) GetBundleOk() (*Bundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *Product) SetBundle(v Bundle)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *Product) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Product) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRecurring

`func (o *Product) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *Product) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *Product) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *Product) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetDiscount

`func (o *Product) GetDiscount() Discount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Product) GetDiscountOk() (*Discount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Product) SetDiscount(v Discount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Product) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetIsMostPopular

`func (o *Product) GetIsMostPopular() bool`

GetIsMostPopular returns the IsMostPopular field if non-nil, zero value otherwise.

### GetIsMostPopularOk

`func (o *Product) GetIsMostPopularOk() (*bool, bool)`

GetIsMostPopularOk returns a tuple with the IsMostPopular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMostPopular

`func (o *Product) SetIsMostPopular(v bool)`

SetIsMostPopular sets IsMostPopular field to given value.

### HasIsMostPopular

`func (o *Product) HasIsMostPopular() bool`

HasIsMostPopular returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



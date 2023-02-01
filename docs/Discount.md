# Discount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float64** |  | [optional] 
**Recurring** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to **float64** |  | [optional] 

## Methods

### NewDiscount

`func NewDiscount() *Discount`

NewDiscount instantiates a new Discount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountWithDefaults

`func NewDiscountWithDefaults() *Discount`

NewDiscountWithDefaults instantiates a new Discount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *Discount) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Discount) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Discount) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Discount) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRecurring

`func (o *Discount) GetRecurring() string`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *Discount) GetRecurringOk() (*string, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *Discount) SetRecurring(v string)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *Discount) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetDiscount

`func (o *Discount) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Discount) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Discount) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Discount) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentMethodCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | **string** | Card brand | 
**Last4** | **string** | Last 4 digits of the card | 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**ExpiresAt** | **string** | Last day of the expiration year and month | 

## Methods

### NewPaymentMethodCard

`func NewPaymentMethodCard(brand string, last4 string, expiresAt string, ) *PaymentMethodCard`

NewPaymentMethodCard instantiates a new PaymentMethodCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardWithDefaults

`func NewPaymentMethodCardWithDefaults() *PaymentMethodCard`

NewPaymentMethodCardWithDefaults instantiates a new PaymentMethodCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *PaymentMethodCard) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaymentMethodCard) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaymentMethodCard) SetBrand(v string)`

SetBrand sets Brand field to given value.


### GetLast4

`func (o *PaymentMethodCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethodCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethodCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### GetCountry

`func (o *PaymentMethodCard) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodCard) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodCard) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodCard) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentMethodCard) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodCard) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodCard) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



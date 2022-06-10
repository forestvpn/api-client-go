# CreateCheckoutSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | **string** |  | 
**SuccessUrl** | **string** |  | 
**Currency** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Products** | [**[]CreateCheckoutSessionProduct**](CreateCheckoutSessionProduct.md) |  | 

## Methods

### NewCreateCheckoutSessionRequest

`func NewCreateCheckoutSessionRequest(cancelUrl string, successUrl string, products []CreateCheckoutSessionProduct, ) *CreateCheckoutSessionRequest`

NewCreateCheckoutSessionRequest instantiates a new CreateCheckoutSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutSessionRequestWithDefaults

`func NewCreateCheckoutSessionRequestWithDefaults() *CreateCheckoutSessionRequest`

NewCreateCheckoutSessionRequestWithDefaults instantiates a new CreateCheckoutSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *CreateCheckoutSessionRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CreateCheckoutSessionRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CreateCheckoutSessionRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetSuccessUrl

`func (o *CreateCheckoutSessionRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CreateCheckoutSessionRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CreateCheckoutSessionRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetCurrency

`func (o *CreateCheckoutSessionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCheckoutSessionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCheckoutSessionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCheckoutSessionRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLocale

`func (o *CreateCheckoutSessionRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateCheckoutSessionRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateCheckoutSessionRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateCheckoutSessionRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetEmail

`func (o *CreateCheckoutSessionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCheckoutSessionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCheckoutSessionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateCheckoutSessionRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetProducts

`func (o *CreateCheckoutSessionRequest) GetProducts() []CreateCheckoutSessionProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreateCheckoutSessionRequest) GetProductsOk() (*[]CreateCheckoutSessionProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreateCheckoutSessionRequest) SetProducts(v []CreateCheckoutSessionProduct)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



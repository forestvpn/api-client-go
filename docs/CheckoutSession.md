# CheckoutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CancelUrl** | **string** |  | 
**SuccessUrl** | **string** |  | 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**Currency** | **string** |  | 
**AmountSubtotal** | **float64** |  | 
**AmountTotal** | **float64** |  | 
**Locale** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Products** | [**[]CheckoutSessionProduct**](CheckoutSessionProduct.md) |  | 
**PaymentStatus** | **string** |  | 
**Status** | **string** |  | 
**TrialPeriod** | Pointer to **string** | Trial period duration in ISO 8601 format. | [optional] 
**User** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewCheckoutSession

`func NewCheckoutSession(id string, cancelUrl string, successUrl string, currency string, amountSubtotal float64, amountTotal float64, products []CheckoutSessionProduct, paymentStatus string, status string, createdAt time.Time, expiresAt time.Time, ) *CheckoutSession`

NewCheckoutSession instantiates a new CheckoutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionWithDefaults

`func NewCheckoutSessionWithDefaults() *CheckoutSession`

NewCheckoutSessionWithDefaults instantiates a new CheckoutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckoutSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSession) SetId(v string)`

SetId sets Id field to given value.


### GetCancelUrl

`func (o *CheckoutSession) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CheckoutSession) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CheckoutSession) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetSuccessUrl

`func (o *CheckoutSession) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutSession) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutSession) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetRedirectUrl

`func (o *CheckoutSession) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CheckoutSession) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CheckoutSession) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CheckoutSession) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetCurrency

`func (o *CheckoutSession) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutSession) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutSession) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmountSubtotal

`func (o *CheckoutSession) GetAmountSubtotal() float64`

GetAmountSubtotal returns the AmountSubtotal field if non-nil, zero value otherwise.

### GetAmountSubtotalOk

`func (o *CheckoutSession) GetAmountSubtotalOk() (*float64, bool)`

GetAmountSubtotalOk returns a tuple with the AmountSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSubtotal

`func (o *CheckoutSession) SetAmountSubtotal(v float64)`

SetAmountSubtotal sets AmountSubtotal field to given value.


### GetAmountTotal

`func (o *CheckoutSession) GetAmountTotal() float64`

GetAmountTotal returns the AmountTotal field if non-nil, zero value otherwise.

### GetAmountTotalOk

`func (o *CheckoutSession) GetAmountTotalOk() (*float64, bool)`

GetAmountTotalOk returns a tuple with the AmountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTotal

`func (o *CheckoutSession) SetAmountTotal(v float64)`

SetAmountTotal sets AmountTotal field to given value.


### GetLocale

`func (o *CheckoutSession) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CheckoutSession) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CheckoutSession) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CheckoutSession) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetEmail

`func (o *CheckoutSession) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CheckoutSession) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CheckoutSession) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CheckoutSession) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetProducts

`func (o *CheckoutSession) GetProducts() []CheckoutSessionProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CheckoutSession) GetProductsOk() (*[]CheckoutSessionProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CheckoutSession) SetProducts(v []CheckoutSessionProduct)`

SetProducts sets Products field to given value.


### GetPaymentStatus

`func (o *CheckoutSession) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CheckoutSession) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CheckoutSession) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetStatus

`func (o *CheckoutSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutSession) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTrialPeriod

`func (o *CheckoutSession) GetTrialPeriod() string`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *CheckoutSession) GetTrialPeriodOk() (*string, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *CheckoutSession) SetTrialPeriod(v string)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *CheckoutSession) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetUser

`func (o *CheckoutSession) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CheckoutSession) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CheckoutSession) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CheckoutSession) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CheckoutSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *CheckoutSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



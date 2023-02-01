/*
ForestVPN API

ForestVPN - Fast, secure, and modern VPN. It offers Distributed Computing, Crypto Mining, P2P, Ad Blocking, TOR over VPN, 30+ locations, and a free version with unlimited data. 

API version: 2.0
Contact: support@forestvpn.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forestvpn_api

import (
	"encoding/json"
	"time"
)

// CheckoutSession struct for CheckoutSession
type CheckoutSession struct {
	Id string `json:"id"`
	CancelUrl string `json:"cancel_url"`
	SuccessUrl string `json:"success_url"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
	Currency string `json:"currency"`
	AmountSubtotal float64 `json:"amount_subtotal"`
	AmountTotal float64 `json:"amount_total"`
	Locale *string `json:"locale,omitempty"`
	Email *string `json:"email,omitempty"`
	Products []CheckoutSessionProduct `json:"products"`
	PaymentStatus string `json:"payment_status"`
	Status string `json:"status"`
	// Trial period duration in ISO 8601 format.
	TrialPeriod *string `json:"trial_period,omitempty"`
	User *string `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// NewCheckoutSession instantiates a new CheckoutSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSession(id string, cancelUrl string, successUrl string, currency string, amountSubtotal float64, amountTotal float64, products []CheckoutSessionProduct, paymentStatus string, status string, createdAt time.Time, expiresAt time.Time) *CheckoutSession {
	this := CheckoutSession{}
	this.Id = id
	this.CancelUrl = cancelUrl
	this.SuccessUrl = successUrl
	this.Currency = currency
	this.AmountSubtotal = amountSubtotal
	this.AmountTotal = amountTotal
	this.Products = products
	this.PaymentStatus = paymentStatus
	this.Status = status
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	return &this
}

// NewCheckoutSessionWithDefaults instantiates a new CheckoutSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionWithDefaults() *CheckoutSession {
	this := CheckoutSession{}
	return &this
}

// GetId returns the Id field value
func (o *CheckoutSession) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutSession) SetId(v string) {
	o.Id = v
}

// GetCancelUrl returns the CancelUrl field value
func (o *CheckoutSession) GetCancelUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetCancelUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CancelUrl, true
}

// SetCancelUrl sets field value
func (o *CheckoutSession) SetCancelUrl(v string) {
	o.CancelUrl = v
}

// GetSuccessUrl returns the SuccessUrl field value
func (o *CheckoutSession) GetSuccessUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetSuccessUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SuccessUrl, true
}

// SetSuccessUrl sets field value
func (o *CheckoutSession) SetSuccessUrl(v string) {
	o.SuccessUrl = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *CheckoutSession) GetRedirectUrl() string {
	if o == nil || isNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetRedirectUrlOk() (*string, bool) {
	if o == nil || isNil(o.RedirectUrl) {
    return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *CheckoutSession) HasRedirectUrl() bool {
	if o != nil && !isNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *CheckoutSession) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetCurrency returns the Currency field value
func (o *CheckoutSession) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetCurrencyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CheckoutSession) SetCurrency(v string) {
	o.Currency = v
}

// GetAmountSubtotal returns the AmountSubtotal field value
func (o *CheckoutSession) GetAmountSubtotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AmountSubtotal
}

// GetAmountSubtotalOk returns a tuple with the AmountSubtotal field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetAmountSubtotalOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AmountSubtotal, true
}

// SetAmountSubtotal sets field value
func (o *CheckoutSession) SetAmountSubtotal(v float64) {
	o.AmountSubtotal = v
}

// GetAmountTotal returns the AmountTotal field value
func (o *CheckoutSession) GetAmountTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AmountTotal
}

// GetAmountTotalOk returns a tuple with the AmountTotal field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetAmountTotalOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AmountTotal, true
}

// SetAmountTotal sets field value
func (o *CheckoutSession) SetAmountTotal(v float64) {
	o.AmountTotal = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *CheckoutSession) GetLocale() string {
	if o == nil || isNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetLocaleOk() (*string, bool) {
	if o == nil || isNil(o.Locale) {
    return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *CheckoutSession) HasLocale() bool {
	if o != nil && !isNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *CheckoutSession) SetLocale(v string) {
	o.Locale = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CheckoutSession) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CheckoutSession) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CheckoutSession) SetEmail(v string) {
	o.Email = &v
}

// GetProducts returns the Products field value
func (o *CheckoutSession) GetProducts() []CheckoutSessionProduct {
	if o == nil {
		var ret []CheckoutSessionProduct
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetProductsOk() ([]CheckoutSessionProduct, bool) {
	if o == nil {
    return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *CheckoutSession) SetProducts(v []CheckoutSessionProduct) {
	o.Products = v
}

// GetPaymentStatus returns the PaymentStatus field value
func (o *CheckoutSession) GetPaymentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetPaymentStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PaymentStatus, true
}

// SetPaymentStatus sets field value
func (o *CheckoutSession) SetPaymentStatus(v string) {
	o.PaymentStatus = v
}

// GetStatus returns the Status field value
func (o *CheckoutSession) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CheckoutSession) SetStatus(v string) {
	o.Status = v
}

// GetTrialPeriod returns the TrialPeriod field value if set, zero value otherwise.
func (o *CheckoutSession) GetTrialPeriod() string {
	if o == nil || isNil(o.TrialPeriod) {
		var ret string
		return ret
	}
	return *o.TrialPeriod
}

// GetTrialPeriodOk returns a tuple with the TrialPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetTrialPeriodOk() (*string, bool) {
	if o == nil || isNil(o.TrialPeriod) {
    return nil, false
	}
	return o.TrialPeriod, true
}

// HasTrialPeriod returns a boolean if a field has been set.
func (o *CheckoutSession) HasTrialPeriod() bool {
	if o != nil && !isNil(o.TrialPeriod) {
		return true
	}

	return false
}

// SetTrialPeriod gets a reference to the given string and assigns it to the TrialPeriod field.
func (o *CheckoutSession) SetTrialPeriod(v string) {
	o.TrialPeriod = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CheckoutSession) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CheckoutSession) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CheckoutSession) SetUser(v string) {
	o.User = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CheckoutSession) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CheckoutSession) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *CheckoutSession) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *CheckoutSession) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

func (o CheckoutSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["cancel_url"] = o.CancelUrl
	}
	if true {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if !isNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["amount_subtotal"] = o.AmountSubtotal
	}
	if true {
		toSerialize["amount_total"] = o.AmountTotal
	}
	if !isNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["products"] = o.Products
	}
	if true {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TrialPeriod) {
		toSerialize["trial_period"] = o.TrialPeriod
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutSession struct {
	value *CheckoutSession
	isSet bool
}

func (v NullableCheckoutSession) Get() *CheckoutSession {
	return v.value
}

func (v *NullableCheckoutSession) Set(val *CheckoutSession) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSession) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSession(val *CheckoutSession) *NullableCheckoutSession {
	return &NullableCheckoutSession{value: val, isSet: true}
}

func (v NullableCheckoutSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



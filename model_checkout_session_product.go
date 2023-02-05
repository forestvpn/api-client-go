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
)

// checks if the CheckoutSessionProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutSessionProduct{}

// CheckoutSessionProduct struct for CheckoutSessionProduct
type CheckoutSessionProduct struct {
	Product Product `json:"product"`
	Quantity int32 `json:"quantity"`
}

// NewCheckoutSessionProduct instantiates a new CheckoutSessionProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionProduct(product Product, quantity int32) *CheckoutSessionProduct {
	this := CheckoutSessionProduct{}
	this.Product = product
	this.Quantity = quantity
	return &this
}

// NewCheckoutSessionProductWithDefaults instantiates a new CheckoutSessionProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionProductWithDefaults() *CheckoutSessionProduct {
	this := CheckoutSessionProduct{}
	return &this
}

// GetProduct returns the Product field value
func (o *CheckoutSessionProduct) GetProduct() Product {
	if o == nil {
		var ret Product
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionProduct) GetProductOk() (*Product, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *CheckoutSessionProduct) SetProduct(v Product) {
	o.Product = v
}

// GetQuantity returns the Quantity field value
func (o *CheckoutSessionProduct) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionProduct) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CheckoutSessionProduct) SetQuantity(v int32) {
	o.Quantity = v
}

func (o CheckoutSessionProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSessionProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

type NullableCheckoutSessionProduct struct {
	value *CheckoutSessionProduct
	isSet bool
}

func (v NullableCheckoutSessionProduct) Get() *CheckoutSessionProduct {
	return v.value
}

func (v *NullableCheckoutSessionProduct) Set(val *CheckoutSessionProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionProduct(val *CheckoutSessionProduct) *NullableCheckoutSessionProduct {
	return &NullableCheckoutSessionProduct{value: val, isSet: true}
}

func (v NullableCheckoutSessionProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



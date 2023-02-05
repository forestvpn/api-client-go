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

// checks if the CreateCheckoutSessionProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCheckoutSessionProduct{}

// CreateCheckoutSessionProduct struct for CreateCheckoutSessionProduct
type CreateCheckoutSessionProduct struct {
	Product string `json:"product"`
	Quantity int32 `json:"quantity"`
}

// NewCreateCheckoutSessionProduct instantiates a new CreateCheckoutSessionProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCheckoutSessionProduct(product string, quantity int32) *CreateCheckoutSessionProduct {
	this := CreateCheckoutSessionProduct{}
	this.Product = product
	this.Quantity = quantity
	return &this
}

// NewCreateCheckoutSessionProductWithDefaults instantiates a new CreateCheckoutSessionProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCheckoutSessionProductWithDefaults() *CreateCheckoutSessionProduct {
	this := CreateCheckoutSessionProduct{}
	return &this
}

// GetProduct returns the Product field value
func (o *CreateCheckoutSessionProduct) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionProduct) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *CreateCheckoutSessionProduct) SetProduct(v string) {
	o.Product = v
}

// GetQuantity returns the Quantity field value
func (o *CreateCheckoutSessionProduct) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionProduct) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CreateCheckoutSessionProduct) SetQuantity(v int32) {
	o.Quantity = v
}

func (o CreateCheckoutSessionProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCheckoutSessionProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

type NullableCreateCheckoutSessionProduct struct {
	value *CreateCheckoutSessionProduct
	isSet bool
}

func (v NullableCreateCheckoutSessionProduct) Get() *CreateCheckoutSessionProduct {
	return v.value
}

func (v *NullableCreateCheckoutSessionProduct) Set(val *CreateCheckoutSessionProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCheckoutSessionProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCheckoutSessionProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCheckoutSessionProduct(val *CreateCheckoutSessionProduct) *NullableCreateCheckoutSessionProduct {
	return &NullableCreateCheckoutSessionProduct{value: val, isSet: true}
}

func (v NullableCreateCheckoutSessionProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCheckoutSessionProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



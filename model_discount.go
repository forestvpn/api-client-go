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

// checks if the Discount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Discount{}

// Discount struct for Discount
type Discount struct {
	Price *float64 `json:"price,omitempty"`
	Recurring *string `json:"recurring,omitempty"`
	Discount *float64 `json:"discount,omitempty"`
}

// NewDiscount instantiates a new Discount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscount() *Discount {
	this := Discount{}
	return &this
}

// NewDiscountWithDefaults instantiates a new Discount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountWithDefaults() *Discount {
	this := Discount{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Discount) GetPrice() float64 {
	if o == nil || isNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetPriceOk() (*float64, bool) {
	if o == nil || isNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Discount) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *Discount) SetPrice(v float64) {
	o.Price = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *Discount) GetRecurring() string {
	if o == nil || isNil(o.Recurring) {
		var ret string
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetRecurringOk() (*string, bool) {
	if o == nil || isNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *Discount) HasRecurring() bool {
	if o != nil && !isNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given string and assigns it to the Recurring field.
func (o *Discount) SetRecurring(v string) {
	o.Recurring = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Discount) GetDiscount() float64 {
	if o == nil || isNil(o.Discount) {
		var ret float64
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discount) GetDiscountOk() (*float64, bool) {
	if o == nil || isNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Discount) HasDiscount() bool {
	if o != nil && !isNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float64 and assigns it to the Discount field.
func (o *Discount) SetDiscount(v float64) {
	o.Discount = &v
}

func (o Discount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Discount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	if !isNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	return toSerialize, nil
}

type NullableDiscount struct {
	value *Discount
	isSet bool
}

func (v NullableDiscount) Get() *Discount {
	return v.value
}

func (v *NullableDiscount) Set(val *Discount) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscount(val *Discount) *NullableDiscount {
	return &NullableDiscount{value: val, isSet: true}
}

func (v NullableDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// ProductWithoutPrice struct for ProductWithoutPrice
type ProductWithoutPrice struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Bundle *Bundle `json:"bundle,omitempty"`
	Recurring *Recurring `json:"recurring,omitempty"`
	Discount *Discount `json:"discount,omitempty"`
	IsMostPopular *bool `json:"is_most_popular,omitempty"`
}

// NewProductWithoutPrice instantiates a new ProductWithoutPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductWithoutPrice(id string, name string) *ProductWithoutPrice {
	this := ProductWithoutPrice{}
	this.Id = id
	this.Name = name
	return &this
}

// NewProductWithoutPriceWithDefaults instantiates a new ProductWithoutPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithoutPriceWithDefaults() *ProductWithoutPrice {
	this := ProductWithoutPrice{}
	return &this
}

// GetId returns the Id field value
func (o *ProductWithoutPrice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductWithoutPrice) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductWithoutPrice) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProductWithoutPrice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductWithoutPrice) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductWithoutPrice) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductWithoutPrice) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWithoutPrice) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductWithoutPrice) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductWithoutPrice) SetDescription(v string) {
	o.Description = &v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *ProductWithoutPrice) GetBundle() Bundle {
	if o == nil || isNil(o.Bundle) {
		var ret Bundle
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWithoutPrice) GetBundleOk() (*Bundle, bool) {
	if o == nil || isNil(o.Bundle) {
    return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *ProductWithoutPrice) HasBundle() bool {
	if o != nil && !isNil(o.Bundle) {
		return true
	}

	return false
}

// SetBundle gets a reference to the given Bundle and assigns it to the Bundle field.
func (o *ProductWithoutPrice) SetBundle(v Bundle) {
	o.Bundle = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *ProductWithoutPrice) GetRecurring() Recurring {
	if o == nil || isNil(o.Recurring) {
		var ret Recurring
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWithoutPrice) GetRecurringOk() (*Recurring, bool) {
	if o == nil || isNil(o.Recurring) {
    return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *ProductWithoutPrice) HasRecurring() bool {
	if o != nil && !isNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given Recurring and assigns it to the Recurring field.
func (o *ProductWithoutPrice) SetRecurring(v Recurring) {
	o.Recurring = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *ProductWithoutPrice) GetDiscount() Discount {
	if o == nil || isNil(o.Discount) {
		var ret Discount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWithoutPrice) GetDiscountOk() (*Discount, bool) {
	if o == nil || isNil(o.Discount) {
    return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *ProductWithoutPrice) HasDiscount() bool {
	if o != nil && !isNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given Discount and assigns it to the Discount field.
func (o *ProductWithoutPrice) SetDiscount(v Discount) {
	o.Discount = &v
}

// GetIsMostPopular returns the IsMostPopular field value if set, zero value otherwise.
func (o *ProductWithoutPrice) GetIsMostPopular() bool {
	if o == nil || isNil(o.IsMostPopular) {
		var ret bool
		return ret
	}
	return *o.IsMostPopular
}

// GetIsMostPopularOk returns a tuple with the IsMostPopular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductWithoutPrice) GetIsMostPopularOk() (*bool, bool) {
	if o == nil || isNil(o.IsMostPopular) {
    return nil, false
	}
	return o.IsMostPopular, true
}

// HasIsMostPopular returns a boolean if a field has been set.
func (o *ProductWithoutPrice) HasIsMostPopular() bool {
	if o != nil && !isNil(o.IsMostPopular) {
		return true
	}

	return false
}

// SetIsMostPopular gets a reference to the given bool and assigns it to the IsMostPopular field.
func (o *ProductWithoutPrice) SetIsMostPopular(v bool) {
	o.IsMostPopular = &v
}

func (o ProductWithoutPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Bundle) {
		toSerialize["bundle"] = o.Bundle
	}
	if !isNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	if !isNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !isNil(o.IsMostPopular) {
		toSerialize["is_most_popular"] = o.IsMostPopular
	}
	return json.Marshal(toSerialize)
}

type NullableProductWithoutPrice struct {
	value *ProductWithoutPrice
	isSet bool
}

func (v NullableProductWithoutPrice) Get() *ProductWithoutPrice {
	return v.value
}

func (v *NullableProductWithoutPrice) Set(val *ProductWithoutPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProductWithoutPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProductWithoutPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductWithoutPrice(val *ProductWithoutPrice) *NullableProductWithoutPrice {
	return &NullableProductWithoutPrice{value: val, isSet: true}
}

func (v NullableProductWithoutPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductWithoutPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


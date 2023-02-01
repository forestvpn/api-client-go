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

// SubscriptionItem struct for SubscriptionItem
type SubscriptionItem struct {
	Id string `json:"id"`
	Price Price `json:"price"`
	PriceId string `json:"price_id"`
	Quantity int32 `json:"quantity"`
}

// NewSubscriptionItem instantiates a new SubscriptionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionItem(id string, price Price, priceId string, quantity int32) *SubscriptionItem {
	this := SubscriptionItem{}
	this.Id = id
	this.Price = price
	this.PriceId = priceId
	this.Quantity = quantity
	return &this
}

// NewSubscriptionItemWithDefaults instantiates a new SubscriptionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionItemWithDefaults() *SubscriptionItem {
	this := SubscriptionItem{}
	return &this
}

// GetId returns the Id field value
func (o *SubscriptionItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionItem) SetId(v string) {
	o.Id = v
}

// GetPrice returns the Price field value
func (o *SubscriptionItem) GetPrice() Price {
	if o == nil {
		var ret Price
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetPriceOk() (*Price, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *SubscriptionItem) SetPrice(v Price) {
	o.Price = v
}

// GetPriceId returns the PriceId field value
func (o *SubscriptionItem) GetPriceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceId
}

// GetPriceIdOk returns a tuple with the PriceId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetPriceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PriceId, true
}

// SetPriceId sets field value
func (o *SubscriptionItem) SetPriceId(v string) {
	o.PriceId = v
}

// GetQuantity returns the Quantity field value
func (o *SubscriptionItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *SubscriptionItem) SetQuantity(v int32) {
	o.Quantity = v
}

func (o SubscriptionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["price_id"] = o.PriceId
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionItem struct {
	value *SubscriptionItem
	isSet bool
}

func (v NullableSubscriptionItem) Get() *SubscriptionItem {
	return v.value
}

func (v *NullableSubscriptionItem) Set(val *SubscriptionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionItem(val *SubscriptionItem) *NullableSubscriptionItem {
	return &NullableSubscriptionItem{value: val, isSet: true}
}

func (v NullableSubscriptionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



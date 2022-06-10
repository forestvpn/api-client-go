/*
ForestVPN API

ForestVPN defeats content restrictions and censorship to deliver unlimited access to video, music, social media, and more, from anywhere in the world. 

API version: 2.0
Contact: support@forestvpn.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forestvpn_api

import (
	"encoding/json"
)

// AdReward struct for AdReward
type AdReward struct {
	Item string `json:"item"`
	Amount int32 `json:"amount"`
}

// NewAdReward instantiates a new AdReward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdReward(item string, amount int32) *AdReward {
	this := AdReward{}
	this.Item = item
	this.Amount = amount
	return &this
}

// NewAdRewardWithDefaults instantiates a new AdReward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdRewardWithDefaults() *AdReward {
	this := AdReward{}
	return &this
}

// GetItem returns the Item field value
func (o *AdReward) GetItem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *AdReward) GetItemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *AdReward) SetItem(v string) {
	o.Item = v
}

// GetAmount returns the Amount field value
func (o *AdReward) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AdReward) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AdReward) SetAmount(v int32) {
	o.Amount = v
}

func (o AdReward) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableAdReward struct {
	value *AdReward
	isSet bool
}

func (v NullableAdReward) Get() *AdReward {
	return v.value
}

func (v *NullableAdReward) Set(val *AdReward) {
	v.value = val
	v.isSet = true
}

func (v NullableAdReward) IsSet() bool {
	return v.isSet
}

func (v *NullableAdReward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdReward(val *AdReward) *NullableAdReward {
	return &NullableAdReward{value: val, isSet: true}
}

func (v NullableAdReward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdReward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


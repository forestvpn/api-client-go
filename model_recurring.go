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

// Recurring struct for Recurring
type Recurring struct {
	// Recurring period in ISO 8601 format.
	Period string `json:"period"`
}

// NewRecurring instantiates a new Recurring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurring(period string) *Recurring {
	this := Recurring{}
	this.Period = period
	return &this
}

// NewRecurringWithDefaults instantiates a new Recurring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringWithDefaults() *Recurring {
	this := Recurring{}
	return &this
}

// GetPeriod returns the Period field value
func (o *Recurring) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *Recurring) GetPeriodOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *Recurring) SetPeriod(v string) {
	o.Period = v
}

func (o Recurring) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["period"] = o.Period
	}
	return json.Marshal(toSerialize)
}

type NullableRecurring struct {
	value *Recurring
	isSet bool
}

func (v NullableRecurring) Get() *Recurring {
	return v.value
}

func (v *NullableRecurring) Set(val *Recurring) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurring) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurring(val *Recurring) *NullableRecurring {
	return &NullableRecurring{value: val, isSet: true}
}

func (v NullableRecurring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



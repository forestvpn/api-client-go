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
	"fmt"
)

// SubscriptionSource the model 'SubscriptionSource'
type SubscriptionSource string

// List of SubscriptionSource
const (
	DEFAULT SubscriptionSource = "default"
	STRIPE SubscriptionSource = "stripe"
	GOOGLE SubscriptionSource = "google"
	APPLE SubscriptionSource = "apple"
	CLOUD_PAYMENTS SubscriptionSource = "cloud_payments"
)

// All allowed values of SubscriptionSource enum
var AllowedSubscriptionSourceEnumValues = []SubscriptionSource{
	"default",
	"stripe",
	"google",
	"apple",
	"cloud_payments",
}

func (v *SubscriptionSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionSource(value)
	for _, existing := range AllowedSubscriptionSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionSource", value)
}

// NewSubscriptionSourceFromValue returns a pointer to a valid SubscriptionSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionSourceFromValue(v string) (*SubscriptionSource, error) {
	ev := SubscriptionSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionSource: valid values are %v", v, AllowedSubscriptionSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionSource) IsValid() bool {
	for _, existing := range AllowedSubscriptionSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionSource value
func (v SubscriptionSource) Ptr() *SubscriptionSource {
	return &v
}

type NullableSubscriptionSource struct {
	value *SubscriptionSource
	isSet bool
}

func (v NullableSubscriptionSource) Get() *SubscriptionSource {
	return v.value
}

func (v *NullableSubscriptionSource) Set(val *SubscriptionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionSource(val *SubscriptionSource) *NullableSubscriptionSource {
	return &NullableSubscriptionSource{value: val, isSet: true}
}

func (v NullableSubscriptionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


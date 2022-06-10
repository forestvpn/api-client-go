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
	"time"
)

// BillingFeature struct for BillingFeature
type BillingFeature struct {
	// Billing feature's slug
	BundleId string `json:"bundle_id"`
	// Billing feature's expiry date
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	Constraints []Constraint `json:"constraints,omitempty"`
}

// NewBillingFeature instantiates a new BillingFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingFeature(bundleId string) *BillingFeature {
	this := BillingFeature{}
	this.BundleId = bundleId
	return &this
}

// NewBillingFeatureWithDefaults instantiates a new BillingFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingFeatureWithDefaults() *BillingFeature {
	this := BillingFeature{}
	return &this
}

// GetBundleId returns the BundleId field value
func (o *BillingFeature) GetBundleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value
// and a boolean to check if the value has been set.
func (o *BillingFeature) GetBundleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleId, true
}

// SetBundleId sets field value
func (o *BillingFeature) SetBundleId(v string) {
	o.BundleId = v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *BillingFeature) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFeature) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *BillingFeature) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *BillingFeature) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *BillingFeature) GetConstraints() []Constraint {
	if o == nil || o.Constraints == nil {
		var ret []Constraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingFeature) GetConstraintsOk() ([]Constraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *BillingFeature) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []Constraint and assigns it to the Constraints field.
func (o *BillingFeature) SetConstraints(v []Constraint) {
	o.Constraints = v
}

func (o BillingFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bundle_id"] = o.BundleId
	}
	if o.ExpiryDate != nil {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	return json.Marshal(toSerialize)
}

type NullableBillingFeature struct {
	value *BillingFeature
	isSet bool
}

func (v NullableBillingFeature) Get() *BillingFeature {
	return v.value
}

func (v *NullableBillingFeature) Set(val *BillingFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingFeature(val *BillingFeature) *NullableBillingFeature {
	return &NullableBillingFeature{value: val, isSet: true}
}

func (v NullableBillingFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// CreateCloudPaymentsAuth struct for CreateCloudPaymentsAuth
type CreateCloudPaymentsAuth struct {
	Cryptogram string `json:"cryptogram"`
	Name string `json:"name"`
}

// NewCreateCloudPaymentsAuth instantiates a new CreateCloudPaymentsAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCloudPaymentsAuth(cryptogram string, name string) *CreateCloudPaymentsAuth {
	this := CreateCloudPaymentsAuth{}
	this.Cryptogram = cryptogram
	this.Name = name
	return &this
}

// NewCreateCloudPaymentsAuthWithDefaults instantiates a new CreateCloudPaymentsAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCloudPaymentsAuthWithDefaults() *CreateCloudPaymentsAuth {
	this := CreateCloudPaymentsAuth{}
	return &this
}

// GetCryptogram returns the Cryptogram field value
func (o *CreateCloudPaymentsAuth) GetCryptogram() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cryptogram
}

// GetCryptogramOk returns a tuple with the Cryptogram field value
// and a boolean to check if the value has been set.
func (o *CreateCloudPaymentsAuth) GetCryptogramOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cryptogram, true
}

// SetCryptogram sets field value
func (o *CreateCloudPaymentsAuth) SetCryptogram(v string) {
	o.Cryptogram = v
}

// GetName returns the Name field value
func (o *CreateCloudPaymentsAuth) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCloudPaymentsAuth) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCloudPaymentsAuth) SetName(v string) {
	o.Name = v
}

func (o CreateCloudPaymentsAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cryptogram"] = o.Cryptogram
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCloudPaymentsAuth struct {
	value *CreateCloudPaymentsAuth
	isSet bool
}

func (v NullableCreateCloudPaymentsAuth) Get() *CreateCloudPaymentsAuth {
	return v.value
}

func (v *NullableCreateCloudPaymentsAuth) Set(val *CreateCloudPaymentsAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCloudPaymentsAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCloudPaymentsAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCloudPaymentsAuth(val *CreateCloudPaymentsAuth) *NullableCreateCloudPaymentsAuth {
	return &NullableCreateCloudPaymentsAuth{value: val, isSet: true}
}

func (v NullableCreateCloudPaymentsAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCloudPaymentsAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



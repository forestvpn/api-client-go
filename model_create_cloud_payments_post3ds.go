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

// CreateCloudPaymentsPost3ds struct for CreateCloudPaymentsPost3ds
type CreateCloudPaymentsPost3ds struct {
	PaRes string `json:"paRes"`
}

// NewCreateCloudPaymentsPost3ds instantiates a new CreateCloudPaymentsPost3ds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCloudPaymentsPost3ds(paRes string) *CreateCloudPaymentsPost3ds {
	this := CreateCloudPaymentsPost3ds{}
	this.PaRes = paRes
	return &this
}

// NewCreateCloudPaymentsPost3dsWithDefaults instantiates a new CreateCloudPaymentsPost3ds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCloudPaymentsPost3dsWithDefaults() *CreateCloudPaymentsPost3ds {
	this := CreateCloudPaymentsPost3ds{}
	return &this
}

// GetPaRes returns the PaRes field value
func (o *CreateCloudPaymentsPost3ds) GetPaRes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaRes
}

// GetPaResOk returns a tuple with the PaRes field value
// and a boolean to check if the value has been set.
func (o *CreateCloudPaymentsPost3ds) GetPaResOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaRes, true
}

// SetPaRes sets field value
func (o *CreateCloudPaymentsPost3ds) SetPaRes(v string) {
	o.PaRes = v
}

func (o CreateCloudPaymentsPost3ds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paRes"] = o.PaRes
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCloudPaymentsPost3ds struct {
	value *CreateCloudPaymentsPost3ds
	isSet bool
}

func (v NullableCreateCloudPaymentsPost3ds) Get() *CreateCloudPaymentsPost3ds {
	return v.value
}

func (v *NullableCreateCloudPaymentsPost3ds) Set(val *CreateCloudPaymentsPost3ds) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCloudPaymentsPost3ds) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCloudPaymentsPost3ds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCloudPaymentsPost3ds(val *CreateCloudPaymentsPost3ds) *NullableCreateCloudPaymentsPost3ds {
	return &NullableCreateCloudPaymentsPost3ds{value: val, isSet: true}
}

func (v NullableCreateCloudPaymentsPost3ds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCloudPaymentsPost3ds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



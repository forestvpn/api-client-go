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

// checks if the CloudPaymentsSecure3d type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudPaymentsSecure3d{}

// CloudPaymentsSecure3d struct for CloudPaymentsSecure3d
type CloudPaymentsSecure3d struct {
	PaReq string `json:"pa_req"`
	AcsUrl string `json:"acs_url"`
	TermUrl string `json:"term_url"`
}

// NewCloudPaymentsSecure3d instantiates a new CloudPaymentsSecure3d object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudPaymentsSecure3d(paReq string, acsUrl string, termUrl string) *CloudPaymentsSecure3d {
	this := CloudPaymentsSecure3d{}
	this.PaReq = paReq
	this.AcsUrl = acsUrl
	this.TermUrl = termUrl
	return &this
}

// NewCloudPaymentsSecure3dWithDefaults instantiates a new CloudPaymentsSecure3d object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudPaymentsSecure3dWithDefaults() *CloudPaymentsSecure3d {
	this := CloudPaymentsSecure3d{}
	return &this
}

// GetPaReq returns the PaReq field value
func (o *CloudPaymentsSecure3d) GetPaReq() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaReq
}

// GetPaReqOk returns a tuple with the PaReq field value
// and a boolean to check if the value has been set.
func (o *CloudPaymentsSecure3d) GetPaReqOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaReq, true
}

// SetPaReq sets field value
func (o *CloudPaymentsSecure3d) SetPaReq(v string) {
	o.PaReq = v
}

// GetAcsUrl returns the AcsUrl field value
func (o *CloudPaymentsSecure3d) GetAcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value
// and a boolean to check if the value has been set.
func (o *CloudPaymentsSecure3d) GetAcsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcsUrl, true
}

// SetAcsUrl sets field value
func (o *CloudPaymentsSecure3d) SetAcsUrl(v string) {
	o.AcsUrl = v
}

// GetTermUrl returns the TermUrl field value
func (o *CloudPaymentsSecure3d) GetTermUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TermUrl
}

// GetTermUrlOk returns a tuple with the TermUrl field value
// and a boolean to check if the value has been set.
func (o *CloudPaymentsSecure3d) GetTermUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermUrl, true
}

// SetTermUrl sets field value
func (o *CloudPaymentsSecure3d) SetTermUrl(v string) {
	o.TermUrl = v
}

func (o CloudPaymentsSecure3d) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudPaymentsSecure3d) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pa_req"] = o.PaReq
	toSerialize["acs_url"] = o.AcsUrl
	toSerialize["term_url"] = o.TermUrl
	return toSerialize, nil
}

type NullableCloudPaymentsSecure3d struct {
	value *CloudPaymentsSecure3d
	isSet bool
}

func (v NullableCloudPaymentsSecure3d) Get() *CloudPaymentsSecure3d {
	return v.value
}

func (v *NullableCloudPaymentsSecure3d) Set(val *CloudPaymentsSecure3d) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudPaymentsSecure3d) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudPaymentsSecure3d) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudPaymentsSecure3d(val *CloudPaymentsSecure3d) *NullableCloudPaymentsSecure3d {
	return &NullableCloudPaymentsSecure3d{value: val, isSet: true}
}

func (v NullableCloudPaymentsSecure3d) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudPaymentsSecure3d) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



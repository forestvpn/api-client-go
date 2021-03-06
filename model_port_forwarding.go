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

// PortForwarding struct for PortForwarding
type PortForwarding struct {
	Id string `json:"id"`
	SrcPort *int32 `json:"src_port,omitempty"`
	DstPort *int32 `json:"dst_port,omitempty"`
}

// NewPortForwarding instantiates a new PortForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortForwarding(id string) *PortForwarding {
	this := PortForwarding{}
	this.Id = id
	return &this
}

// NewPortForwardingWithDefaults instantiates a new PortForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortForwardingWithDefaults() *PortForwarding {
	this := PortForwarding{}
	return &this
}

// GetId returns the Id field value
func (o *PortForwarding) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PortForwarding) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PortForwarding) SetId(v string) {
	o.Id = v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *PortForwarding) GetSrcPort() int32 {
	if o == nil || o.SrcPort == nil {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortForwarding) GetSrcPortOk() (*int32, bool) {
	if o == nil || o.SrcPort == nil {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *PortForwarding) HasSrcPort() bool {
	if o != nil && o.SrcPort != nil {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *PortForwarding) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *PortForwarding) GetDstPort() int32 {
	if o == nil || o.DstPort == nil {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortForwarding) GetDstPortOk() (*int32, bool) {
	if o == nil || o.DstPort == nil {
		return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *PortForwarding) HasDstPort() bool {
	if o != nil && o.DstPort != nil {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *PortForwarding) SetDstPort(v int32) {
	o.DstPort = &v
}

func (o PortForwarding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.SrcPort != nil {
		toSerialize["src_port"] = o.SrcPort
	}
	if o.DstPort != nil {
		toSerialize["dst_port"] = o.DstPort
	}
	return json.Marshal(toSerialize)
}

type NullablePortForwarding struct {
	value *PortForwarding
	isSet bool
}

func (v NullablePortForwarding) Get() *PortForwarding {
	return v.value
}

func (v *NullablePortForwarding) Set(val *PortForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullablePortForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullablePortForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortForwarding(val *PortForwarding) *NullablePortForwarding {
	return &NullablePortForwarding{value: val, isSet: true}
}

func (v NullablePortForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



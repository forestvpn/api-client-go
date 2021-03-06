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

// CreateOrUpdatePortForwardingRequest struct for CreateOrUpdatePortForwardingRequest
type CreateOrUpdatePortForwardingRequest struct {
	DstPort *int32 `json:"dst_port,omitempty"`
}

// NewCreateOrUpdatePortForwardingRequest instantiates a new CreateOrUpdatePortForwardingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdatePortForwardingRequest() *CreateOrUpdatePortForwardingRequest {
	this := CreateOrUpdatePortForwardingRequest{}
	return &this
}

// NewCreateOrUpdatePortForwardingRequestWithDefaults instantiates a new CreateOrUpdatePortForwardingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdatePortForwardingRequestWithDefaults() *CreateOrUpdatePortForwardingRequest {
	this := CreateOrUpdatePortForwardingRequest{}
	return &this
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *CreateOrUpdatePortForwardingRequest) GetDstPort() int32 {
	if o == nil || o.DstPort == nil {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePortForwardingRequest) GetDstPortOk() (*int32, bool) {
	if o == nil || o.DstPort == nil {
		return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *CreateOrUpdatePortForwardingRequest) HasDstPort() bool {
	if o != nil && o.DstPort != nil {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *CreateOrUpdatePortForwardingRequest) SetDstPort(v int32) {
	o.DstPort = &v
}

func (o CreateOrUpdatePortForwardingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DstPort != nil {
		toSerialize["dst_port"] = o.DstPort
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdatePortForwardingRequest struct {
	value *CreateOrUpdatePortForwardingRequest
	isSet bool
}

func (v NullableCreateOrUpdatePortForwardingRequest) Get() *CreateOrUpdatePortForwardingRequest {
	return v.value
}

func (v *NullableCreateOrUpdatePortForwardingRequest) Set(val *CreateOrUpdatePortForwardingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdatePortForwardingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdatePortForwardingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdatePortForwardingRequest(val *CreateOrUpdatePortForwardingRequest) *NullableCreateOrUpdatePortForwardingRequest {
	return &NullableCreateOrUpdatePortForwardingRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdatePortForwardingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdatePortForwardingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



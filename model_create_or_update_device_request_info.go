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

// CreateOrUpdateDeviceRequestInfo struct for CreateOrUpdateDeviceRequestInfo
type CreateOrUpdateDeviceRequestInfo struct {
	Type DeviceType `json:"type"`
	Info map[string]string `json:"info"`
}

// NewCreateOrUpdateDeviceRequestInfo instantiates a new CreateOrUpdateDeviceRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateDeviceRequestInfo(type_ DeviceType, info map[string]string) *CreateOrUpdateDeviceRequestInfo {
	this := CreateOrUpdateDeviceRequestInfo{}
	this.Type = type_
	this.Info = info
	return &this
}

// NewCreateOrUpdateDeviceRequestInfoWithDefaults instantiates a new CreateOrUpdateDeviceRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateDeviceRequestInfoWithDefaults() *CreateOrUpdateDeviceRequestInfo {
	this := CreateOrUpdateDeviceRequestInfo{}
	return &this
}

// GetType returns the Type field value
func (o *CreateOrUpdateDeviceRequestInfo) GetType() DeviceType {
	if o == nil {
		var ret DeviceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDeviceRequestInfo) GetTypeOk() (*DeviceType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateOrUpdateDeviceRequestInfo) SetType(v DeviceType) {
	o.Type = v
}

// GetInfo returns the Info field value
func (o *CreateOrUpdateDeviceRequestInfo) GetInfo() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDeviceRequestInfo) GetInfoOk() (*map[string]string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *CreateOrUpdateDeviceRequestInfo) SetInfo(v map[string]string) {
	o.Info = v
}

func (o CreateOrUpdateDeviceRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["info"] = o.Info
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdateDeviceRequestInfo struct {
	value *CreateOrUpdateDeviceRequestInfo
	isSet bool
}

func (v NullableCreateOrUpdateDeviceRequestInfo) Get() *CreateOrUpdateDeviceRequestInfo {
	return v.value
}

func (v *NullableCreateOrUpdateDeviceRequestInfo) Set(val *CreateOrUpdateDeviceRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateDeviceRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateDeviceRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateDeviceRequestInfo(val *CreateOrUpdateDeviceRequestInfo) *NullableCreateOrUpdateDeviceRequestInfo {
	return &NullableCreateOrUpdateDeviceRequestInfo{value: val, isSet: true}
}

func (v NullableCreateOrUpdateDeviceRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateDeviceRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



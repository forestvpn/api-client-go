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

// WireGuardPeerInfo struct for WireGuardPeerInfo
type WireGuardPeerInfo struct {
	PubKey *string `json:"pub_key,omitempty"`
	User *WireGuardPeerUser `json:"user,omitempty"`
	Device *WireGuardPeerDevice `json:"device,omitempty"`
}

// NewWireGuardPeerInfo instantiates a new WireGuardPeerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireGuardPeerInfo() *WireGuardPeerInfo {
	this := WireGuardPeerInfo{}
	return &this
}

// NewWireGuardPeerInfoWithDefaults instantiates a new WireGuardPeerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireGuardPeerInfoWithDefaults() *WireGuardPeerInfo {
	this := WireGuardPeerInfo{}
	return &this
}

// GetPubKey returns the PubKey field value if set, zero value otherwise.
func (o *WireGuardPeerInfo) GetPubKey() string {
	if o == nil || o.PubKey == nil {
		var ret string
		return ret
	}
	return *o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerInfo) GetPubKeyOk() (*string, bool) {
	if o == nil || o.PubKey == nil {
		return nil, false
	}
	return o.PubKey, true
}

// HasPubKey returns a boolean if a field has been set.
func (o *WireGuardPeerInfo) HasPubKey() bool {
	if o != nil && o.PubKey != nil {
		return true
	}

	return false
}

// SetPubKey gets a reference to the given string and assigns it to the PubKey field.
func (o *WireGuardPeerInfo) SetPubKey(v string) {
	o.PubKey = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *WireGuardPeerInfo) GetUser() WireGuardPeerUser {
	if o == nil || o.User == nil {
		var ret WireGuardPeerUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerInfo) GetUserOk() (*WireGuardPeerUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *WireGuardPeerInfo) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given WireGuardPeerUser and assigns it to the User field.
func (o *WireGuardPeerInfo) SetUser(v WireGuardPeerUser) {
	o.User = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *WireGuardPeerInfo) GetDevice() WireGuardPeerDevice {
	if o == nil || o.Device == nil {
		var ret WireGuardPeerDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerInfo) GetDeviceOk() (*WireGuardPeerDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *WireGuardPeerInfo) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given WireGuardPeerDevice and assigns it to the Device field.
func (o *WireGuardPeerInfo) SetDevice(v WireGuardPeerDevice) {
	o.Device = &v
}

func (o WireGuardPeerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PubKey != nil {
		toSerialize["pub_key"] = o.PubKey
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	return json.Marshal(toSerialize)
}

type NullableWireGuardPeerInfo struct {
	value *WireGuardPeerInfo
	isSet bool
}

func (v NullableWireGuardPeerInfo) Get() *WireGuardPeerInfo {
	return v.value
}

func (v *NullableWireGuardPeerInfo) Set(val *WireGuardPeerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWireGuardPeerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWireGuardPeerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireGuardPeerInfo(val *WireGuardPeerInfo) *NullableWireGuardPeerInfo {
	return &NullableWireGuardPeerInfo{value: val, isSet: true}
}

func (v NullableWireGuardPeerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireGuardPeerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



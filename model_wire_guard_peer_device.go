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

// checks if the WireGuardPeerDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireGuardPeerDevice{}

// WireGuardPeerDevice struct for WireGuardPeerDevice
type WireGuardPeerDevice struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewWireGuardPeerDevice instantiates a new WireGuardPeerDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireGuardPeerDevice() *WireGuardPeerDevice {
	this := WireGuardPeerDevice{}
	return &this
}

// NewWireGuardPeerDeviceWithDefaults instantiates a new WireGuardPeerDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireGuardPeerDeviceWithDefaults() *WireGuardPeerDevice {
	this := WireGuardPeerDevice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WireGuardPeerDevice) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerDevice) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WireGuardPeerDevice) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WireGuardPeerDevice) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WireGuardPeerDevice) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerDevice) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WireGuardPeerDevice) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WireGuardPeerDevice) SetName(v string) {
	o.Name = &v
}

func (o WireGuardPeerDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireGuardPeerDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableWireGuardPeerDevice struct {
	value *WireGuardPeerDevice
	isSet bool
}

func (v NullableWireGuardPeerDevice) Get() *WireGuardPeerDevice {
	return v.value
}

func (v *NullableWireGuardPeerDevice) Set(val *WireGuardPeerDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableWireGuardPeerDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableWireGuardPeerDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireGuardPeerDevice(val *WireGuardPeerDevice) *NullableWireGuardPeerDevice {
	return &NullableWireGuardPeerDevice{value: val, isSet: true}
}

func (v NullableWireGuardPeerDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireGuardPeerDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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
	"time"
)

// FCMDevice struct for FCMDevice
type FCMDevice struct {
	RegistrationId string `json:"registration_id"`
	DeviceId *string `json:"device_id,omitempty"`
	Active bool `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

// NewFCMDevice instantiates a new FCMDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFCMDevice(registrationId string, active bool, createdAt time.Time) *FCMDevice {
	this := FCMDevice{}
	this.RegistrationId = registrationId
	this.Active = active
	this.CreatedAt = createdAt
	return &this
}

// NewFCMDeviceWithDefaults instantiates a new FCMDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFCMDeviceWithDefaults() *FCMDevice {
	this := FCMDevice{}
	return &this
}

// GetRegistrationId returns the RegistrationId field value
func (o *FCMDevice) GetRegistrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistrationId
}

// GetRegistrationIdOk returns a tuple with the RegistrationId field value
// and a boolean to check if the value has been set.
func (o *FCMDevice) GetRegistrationIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegistrationId, true
}

// SetRegistrationId sets field value
func (o *FCMDevice) SetRegistrationId(v string) {
	o.RegistrationId = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *FCMDevice) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCMDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *FCMDevice) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *FCMDevice) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetActive returns the Active field value
func (o *FCMDevice) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *FCMDevice) GetActiveOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *FCMDevice) SetActive(v bool) {
	o.Active = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FCMDevice) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FCMDevice) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FCMDevice) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o FCMDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["registration_id"] = o.RegistrationId
	}
	if !isNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableFCMDevice struct {
	value *FCMDevice
	isSet bool
}

func (v NullableFCMDevice) Get() *FCMDevice {
	return v.value
}

func (v *NullableFCMDevice) Set(val *FCMDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableFCMDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableFCMDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFCMDevice(val *FCMDevice) *NullableFCMDevice {
	return &NullableFCMDevice{value: val, isSet: true}
}

func (v NullableFCMDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFCMDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



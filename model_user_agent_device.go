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

// UserAgentDevice struct for UserAgentDevice
type UserAgentDevice struct {
	// It might be \"Other\" in case if it can't be recognized
	Family *string `json:"family,omitempty"`
	Brand NullableString `json:"brand,omitempty"`
	Model NullableString `json:"model,omitempty"`
}

// NewUserAgentDevice instantiates a new UserAgentDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAgentDevice() *UserAgentDevice {
	this := UserAgentDevice{}
	return &this
}

// NewUserAgentDeviceWithDefaults instantiates a new UserAgentDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAgentDeviceWithDefaults() *UserAgentDevice {
	this := UserAgentDevice{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *UserAgentDevice) GetFamily() string {
	if o == nil || isNil(o.Family) {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentDevice) GetFamilyOk() (*string, bool) {
	if o == nil || isNil(o.Family) {
    return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *UserAgentDevice) HasFamily() bool {
	if o != nil && !isNil(o.Family) {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *UserAgentDevice) SetFamily(v string) {
	o.Family = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAgentDevice) GetBrand() string {
	if o == nil || isNil(o.Brand.Get()) {
		var ret string
		return ret
	}
	return *o.Brand.Get()
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAgentDevice) GetBrandOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Brand.Get(), o.Brand.IsSet()
}

// HasBrand returns a boolean if a field has been set.
func (o *UserAgentDevice) HasBrand() bool {
	if o != nil && o.Brand.IsSet() {
		return true
	}

	return false
}

// SetBrand gets a reference to the given NullableString and assigns it to the Brand field.
func (o *UserAgentDevice) SetBrand(v string) {
	o.Brand.Set(&v)
}
// SetBrandNil sets the value for Brand to be an explicit nil
func (o *UserAgentDevice) SetBrandNil() {
	o.Brand.Set(nil)
}

// UnsetBrand ensures that no value is present for Brand, not even an explicit nil
func (o *UserAgentDevice) UnsetBrand() {
	o.Brand.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAgentDevice) GetModel() string {
	if o == nil || isNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAgentDevice) GetModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *UserAgentDevice) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *UserAgentDevice) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *UserAgentDevice) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *UserAgentDevice) UnsetModel() {
	o.Model.Unset()
}

func (o UserAgentDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Family) {
		toSerialize["family"] = o.Family
	}
	if o.Brand.IsSet() {
		toSerialize["brand"] = o.Brand.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserAgentDevice struct {
	value *UserAgentDevice
	isSet bool
}

func (v NullableUserAgentDevice) Get() *UserAgentDevice {
	return v.value
}

func (v *NullableUserAgentDevice) Set(val *UserAgentDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAgentDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAgentDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAgentDevice(val *UserAgentDevice) *NullableUserAgentDevice {
	return &NullableUserAgentDevice{value: val, isSet: true}
}

func (v NullableUserAgentDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAgentDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



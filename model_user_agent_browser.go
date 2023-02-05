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

// checks if the UserAgentBrowser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAgentBrowser{}

// UserAgentBrowser struct for UserAgentBrowser
type UserAgentBrowser struct {
	// It might be \"Other\" in case if it can't be recognized
	Family *string `json:"family,omitempty"`
	// It might be empty string
	Version *string `json:"version,omitempty"`
}

// NewUserAgentBrowser instantiates a new UserAgentBrowser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAgentBrowser() *UserAgentBrowser {
	this := UserAgentBrowser{}
	return &this
}

// NewUserAgentBrowserWithDefaults instantiates a new UserAgentBrowser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAgentBrowserWithDefaults() *UserAgentBrowser {
	this := UserAgentBrowser{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *UserAgentBrowser) GetFamily() string {
	if o == nil || isNil(o.Family) {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentBrowser) GetFamilyOk() (*string, bool) {
	if o == nil || isNil(o.Family) {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *UserAgentBrowser) HasFamily() bool {
	if o != nil && !isNil(o.Family) {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *UserAgentBrowser) SetFamily(v string) {
	o.Family = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UserAgentBrowser) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAgentBrowser) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UserAgentBrowser) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UserAgentBrowser) SetVersion(v string) {
	o.Version = &v
}

func (o UserAgentBrowser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAgentBrowser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Family) {
		toSerialize["family"] = o.Family
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableUserAgentBrowser struct {
	value *UserAgentBrowser
	isSet bool
}

func (v NullableUserAgentBrowser) Get() *UserAgentBrowser {
	return v.value
}

func (v *NullableUserAgentBrowser) Set(val *UserAgentBrowser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAgentBrowser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAgentBrowser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAgentBrowser(val *UserAgentBrowser) *NullableUserAgentBrowser {
	return &NullableUserAgentBrowser{value: val, isSet: true}
}

func (v NullableUserAgentBrowser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAgentBrowser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



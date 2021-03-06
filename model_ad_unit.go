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

// AdUnit struct for AdUnit
type AdUnit struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ExternalKey *string `json:"external_key,omitempty"`
	Format string `json:"format"`
	Reward *AdReward `json:"reward,omitempty"`
	Provider AdProvider `json:"provider"`
}

// NewAdUnit instantiates a new AdUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdUnit(id string, name string, format string, provider AdProvider) *AdUnit {
	this := AdUnit{}
	this.Id = id
	this.Name = name
	this.Format = format
	this.Provider = provider
	return &this
}

// NewAdUnitWithDefaults instantiates a new AdUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdUnitWithDefaults() *AdUnit {
	this := AdUnit{}
	return &this
}

// GetId returns the Id field value
func (o *AdUnit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdUnit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdUnit) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AdUnit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdUnit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdUnit) SetName(v string) {
	o.Name = v
}

// GetExternalKey returns the ExternalKey field value if set, zero value otherwise.
func (o *AdUnit) GetExternalKey() string {
	if o == nil || o.ExternalKey == nil {
		var ret string
		return ret
	}
	return *o.ExternalKey
}

// GetExternalKeyOk returns a tuple with the ExternalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdUnit) GetExternalKeyOk() (*string, bool) {
	if o == nil || o.ExternalKey == nil {
		return nil, false
	}
	return o.ExternalKey, true
}

// HasExternalKey returns a boolean if a field has been set.
func (o *AdUnit) HasExternalKey() bool {
	if o != nil && o.ExternalKey != nil {
		return true
	}

	return false
}

// SetExternalKey gets a reference to the given string and assigns it to the ExternalKey field.
func (o *AdUnit) SetExternalKey(v string) {
	o.ExternalKey = &v
}

// GetFormat returns the Format field value
func (o *AdUnit) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *AdUnit) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *AdUnit) SetFormat(v string) {
	o.Format = v
}

// GetReward returns the Reward field value if set, zero value otherwise.
func (o *AdUnit) GetReward() AdReward {
	if o == nil || o.Reward == nil {
		var ret AdReward
		return ret
	}
	return *o.Reward
}

// GetRewardOk returns a tuple with the Reward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdUnit) GetRewardOk() (*AdReward, bool) {
	if o == nil || o.Reward == nil {
		return nil, false
	}
	return o.Reward, true
}

// HasReward returns a boolean if a field has been set.
func (o *AdUnit) HasReward() bool {
	if o != nil && o.Reward != nil {
		return true
	}

	return false
}

// SetReward gets a reference to the given AdReward and assigns it to the Reward field.
func (o *AdUnit) SetReward(v AdReward) {
	o.Reward = &v
}

// GetProvider returns the Provider field value
func (o *AdUnit) GetProvider() AdProvider {
	if o == nil {
		var ret AdProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AdUnit) GetProviderOk() (*AdProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AdUnit) SetProvider(v AdProvider) {
	o.Provider = v
}

func (o AdUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExternalKey != nil {
		toSerialize["external_key"] = o.ExternalKey
	}
	if true {
		toSerialize["format"] = o.Format
	}
	if o.Reward != nil {
		toSerialize["reward"] = o.Reward
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableAdUnit struct {
	value *AdUnit
	isSet bool
}

func (v NullableAdUnit) Get() *AdUnit {
	return v.value
}

func (v *NullableAdUnit) Set(val *AdUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableAdUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableAdUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdUnit(val *AdUnit) *NullableAdUnit {
	return &NullableAdUnit{value: val, isSet: true}
}

func (v NullableAdUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



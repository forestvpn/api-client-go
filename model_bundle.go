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

// checks if the Bundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bundle{}

// Bundle struct for Bundle
type Bundle struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// Trial period duration in ISO 8601 format.
	TrialPeriod *string `json:"trial_period,omitempty"`
	Products []Product `json:"products,omitempty"`
}

// NewBundle instantiates a new Bundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundle(id string, name string) *Bundle {
	this := Bundle{}
	this.Id = id
	this.Name = name
	return &this
}

// NewBundleWithDefaults instantiates a new Bundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleWithDefaults() *Bundle {
	this := Bundle{}
	return &this
}

// GetId returns the Id field value
func (o *Bundle) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Bundle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Bundle) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Bundle) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Bundle) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Bundle) SetName(v string) {
	o.Name = v
}

// GetTrialPeriod returns the TrialPeriod field value if set, zero value otherwise.
func (o *Bundle) GetTrialPeriod() string {
	if o == nil || isNil(o.TrialPeriod) {
		var ret string
		return ret
	}
	return *o.TrialPeriod
}

// GetTrialPeriodOk returns a tuple with the TrialPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetTrialPeriodOk() (*string, bool) {
	if o == nil || isNil(o.TrialPeriod) {
		return nil, false
	}
	return o.TrialPeriod, true
}

// HasTrialPeriod returns a boolean if a field has been set.
func (o *Bundle) HasTrialPeriod() bool {
	if o != nil && !isNil(o.TrialPeriod) {
		return true
	}

	return false
}

// SetTrialPeriod gets a reference to the given string and assigns it to the TrialPeriod field.
func (o *Bundle) SetTrialPeriod(v string) {
	o.TrialPeriod = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Bundle) GetProducts() []Product {
	if o == nil || isNil(o.Products) {
		var ret []Product
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bundle) GetProductsOk() ([]Product, bool) {
	if o == nil || isNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Bundle) HasProducts() bool {
	if o != nil && !isNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []Product and assigns it to the Products field.
func (o *Bundle) SetProducts(v []Product) {
	o.Products = v
}

func (o Bundle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !isNil(o.TrialPeriod) {
		toSerialize["trial_period"] = o.TrialPeriod
	}
	if !isNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableBundle struct {
	value *Bundle
	isSet bool
}

func (v NullableBundle) Get() *Bundle {
	return v.value
}

func (v *NullableBundle) Set(val *Bundle) {
	v.value = val
	v.isSet = true
}

func (v NullableBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundle(val *Bundle) *NullableBundle {
	return &NullableBundle{value: val, isSet: true}
}

func (v NullableBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



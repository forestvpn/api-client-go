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

// Location struct for Location
type Location struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Country Country `json:"country"`
	AlternativeNames []string `json:"alternative_names,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation(id string, name string, latitude float32, longitude float32, country Country) *Location {
	this := Location{}
	this.Id = id
	this.Name = name
	this.Latitude = latitude
	this.Longitude = longitude
	this.Country = country
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetId returns the Id field value
func (o *Location) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Location) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Location) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Location) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Location) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Location) SetName(v string) {
	o.Name = v
}

// GetLatitude returns the Latitude field value
func (o *Location) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *Location) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *Location) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *Location) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *Location) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *Location) SetLongitude(v float32) {
	o.Longitude = v
}

// GetCountry returns the Country field value
func (o *Location) GetCountry() Country {
	if o == nil {
		var ret Country
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *Location) GetCountryOk() (*Country, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Location) SetCountry(v Country) {
	o.Country = v
}

// GetAlternativeNames returns the AlternativeNames field value if set, zero value otherwise.
func (o *Location) GetAlternativeNames() []string {
	if o == nil || o.AlternativeNames == nil {
		var ret []string
		return ret
	}
	return o.AlternativeNames
}

// GetAlternativeNamesOk returns a tuple with the AlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAlternativeNamesOk() ([]string, bool) {
	if o == nil || o.AlternativeNames == nil {
		return nil, false
	}
	return o.AlternativeNames, true
}

// HasAlternativeNames returns a boolean if a field has been set.
func (o *Location) HasAlternativeNames() bool {
	if o != nil && o.AlternativeNames != nil {
		return true
	}

	return false
}

// SetAlternativeNames gets a reference to the given []string and assigns it to the AlternativeNames field.
func (o *Location) SetAlternativeNames(v []string) {
	o.AlternativeNames = v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["latitude"] = o.Latitude
	}
	if true {
		toSerialize["longitude"] = o.Longitude
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if o.AlternativeNames != nil {
		toSerialize["alternative_names"] = o.AlternativeNames
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


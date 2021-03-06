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

// WireGuardPeerUser struct for WireGuardPeerUser
type WireGuardPeerUser struct {
	Id *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Email *string `json:"email,omitempty"`
	PhotoUrl *string `json:"photo_url,omitempty"`
}

// NewWireGuardPeerUser instantiates a new WireGuardPeerUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireGuardPeerUser() *WireGuardPeerUser {
	this := WireGuardPeerUser{}
	return &this
}

// NewWireGuardPeerUserWithDefaults instantiates a new WireGuardPeerUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireGuardPeerUserWithDefaults() *WireGuardPeerUser {
	this := WireGuardPeerUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WireGuardPeerUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WireGuardPeerUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WireGuardPeerUser) SetId(v string) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *WireGuardPeerUser) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerUser) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *WireGuardPeerUser) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *WireGuardPeerUser) SetUsername(v string) {
	o.Username = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *WireGuardPeerUser) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerUser) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *WireGuardPeerUser) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *WireGuardPeerUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *WireGuardPeerUser) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerUser) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *WireGuardPeerUser) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *WireGuardPeerUser) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *WireGuardPeerUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *WireGuardPeerUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *WireGuardPeerUser) SetEmail(v string) {
	o.Email = &v
}

// GetPhotoUrl returns the PhotoUrl field value if set, zero value otherwise.
func (o *WireGuardPeerUser) GetPhotoUrl() string {
	if o == nil || o.PhotoUrl == nil {
		var ret string
		return ret
	}
	return *o.PhotoUrl
}

// GetPhotoUrlOk returns a tuple with the PhotoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireGuardPeerUser) GetPhotoUrlOk() (*string, bool) {
	if o == nil || o.PhotoUrl == nil {
		return nil, false
	}
	return o.PhotoUrl, true
}

// HasPhotoUrl returns a boolean if a field has been set.
func (o *WireGuardPeerUser) HasPhotoUrl() bool {
	if o != nil && o.PhotoUrl != nil {
		return true
	}

	return false
}

// SetPhotoUrl gets a reference to the given string and assigns it to the PhotoUrl field.
func (o *WireGuardPeerUser) SetPhotoUrl(v string) {
	o.PhotoUrl = &v
}

func (o WireGuardPeerUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhotoUrl != nil {
		toSerialize["photo_url"] = o.PhotoUrl
	}
	return json.Marshal(toSerialize)
}

type NullableWireGuardPeerUser struct {
	value *WireGuardPeerUser
	isSet bool
}

func (v NullableWireGuardPeerUser) Get() *WireGuardPeerUser {
	return v.value
}

func (v *NullableWireGuardPeerUser) Set(val *WireGuardPeerUser) {
	v.value = val
	v.isSet = true
}

func (v NullableWireGuardPeerUser) IsSet() bool {
	return v.isSet
}

func (v *NullableWireGuardPeerUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireGuardPeerUser(val *WireGuardPeerUser) *NullableWireGuardPeerUser {
	return &NullableWireGuardPeerUser{value: val, isSet: true}
}

func (v NullableWireGuardPeerUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireGuardPeerUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// TokenLogin struct for TokenLogin
type TokenLogin struct {
	FirebaseToken string `json:"firebase_token"`
}

// NewTokenLogin instantiates a new TokenLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenLogin(firebaseToken string) *TokenLogin {
	this := TokenLogin{}
	this.FirebaseToken = firebaseToken
	return &this
}

// NewTokenLoginWithDefaults instantiates a new TokenLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenLoginWithDefaults() *TokenLogin {
	this := TokenLogin{}
	return &this
}

// GetFirebaseToken returns the FirebaseToken field value
func (o *TokenLogin) GetFirebaseToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirebaseToken
}

// GetFirebaseTokenOk returns a tuple with the FirebaseToken field value
// and a boolean to check if the value has been set.
func (o *TokenLogin) GetFirebaseTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirebaseToken, true
}

// SetFirebaseToken sets field value
func (o *TokenLogin) SetFirebaseToken(v string) {
	o.FirebaseToken = v
}

func (o TokenLogin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["firebase_token"] = o.FirebaseToken
	}
	return json.Marshal(toSerialize)
}

type NullableTokenLogin struct {
	value *TokenLogin
	isSet bool
}

func (v NullableTokenLogin) Get() *TokenLogin {
	return v.value
}

func (v *NullableTokenLogin) Set(val *TokenLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenLogin(val *TokenLogin) *NullableTokenLogin {
	return &NullableTokenLogin{value: val, isSet: true}
}

func (v NullableTokenLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



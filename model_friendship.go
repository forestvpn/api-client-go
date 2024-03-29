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

// checks if the Friendship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Friendship{}

// Friendship struct for Friendship
type Friendship struct {
	Id string `json:"id"`
	User *User `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// NewFriendship instantiates a new Friendship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFriendship(id string, createdAt time.Time) *Friendship {
	this := Friendship{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewFriendshipWithDefaults instantiates a new Friendship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFriendshipWithDefaults() *Friendship {
	this := Friendship{}
	return &this
}

// GetId returns the Id field value
func (o *Friendship) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Friendship) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Friendship) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Friendship) GetUser() User {
	if o == nil || isNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Friendship) GetUserOk() (*User, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Friendship) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Friendship) SetUser(v User) {
	o.User = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Friendship) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Friendship) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Friendship) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o Friendship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Friendship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableFriendship struct {
	value *Friendship
	isSet bool
}

func (v NullableFriendship) Get() *Friendship {
	return v.value
}

func (v *NullableFriendship) Set(val *Friendship) {
	v.value = val
	v.isSet = true
}

func (v NullableFriendship) IsSet() bool {
	return v.isSet
}

func (v *NullableFriendship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFriendship(val *Friendship) *NullableFriendship {
	return &NullableFriendship{value: val, isSet: true}
}

func (v NullableFriendship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFriendship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



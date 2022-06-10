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

// NotificationUnreadCount struct for NotificationUnreadCount
type NotificationUnreadCount struct {
	Count int32 `json:"count"`
}

// NewNotificationUnreadCount instantiates a new NotificationUnreadCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationUnreadCount(count int32) *NotificationUnreadCount {
	this := NotificationUnreadCount{}
	this.Count = count
	return &this
}

// NewNotificationUnreadCountWithDefaults instantiates a new NotificationUnreadCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationUnreadCountWithDefaults() *NotificationUnreadCount {
	this := NotificationUnreadCount{}
	return &this
}

// GetCount returns the Count field value
func (o *NotificationUnreadCount) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *NotificationUnreadCount) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *NotificationUnreadCount) SetCount(v int32) {
	o.Count = v
}

func (o NotificationUnreadCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationUnreadCount struct {
	value *NotificationUnreadCount
	isSet bool
}

func (v NullableNotificationUnreadCount) Get() *NotificationUnreadCount {
	return v.value
}

func (v *NullableNotificationUnreadCount) Set(val *NotificationUnreadCount) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationUnreadCount) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationUnreadCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationUnreadCount(val *NotificationUnreadCount) *NullableNotificationUnreadCount {
	return &NullableNotificationUnreadCount{value: val, isSet: true}
}

func (v NullableNotificationUnreadCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationUnreadCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


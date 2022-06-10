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

// Constraint struct for Constraint
type Constraint struct {
	Namespace *string `json:"namespace,omitempty"`
	Relation *string `json:"relation,omitempty"`
	Subject []string `json:"subject,omitempty"`
}

// NewConstraint instantiates a new Constraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraint() *Constraint {
	this := Constraint{}
	return &this
}

// NewConstraintWithDefaults instantiates a new Constraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintWithDefaults() *Constraint {
	this := Constraint{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Constraint) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Constraint) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Constraint) SetNamespace(v string) {
	o.Namespace = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *Constraint) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *Constraint) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *Constraint) SetRelation(v string) {
	o.Relation = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Constraint) GetSubject() []string {
	if o == nil || o.Subject == nil {
		var ret []string
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetSubjectOk() ([]string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Constraint) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []string and assigns it to the Subject field.
func (o *Constraint) SetSubject(v []string) {
	o.Subject = v
}

func (o Constraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableConstraint struct {
	value *Constraint
	isSet bool
}

func (v NullableConstraint) Get() *Constraint {
	return v.value
}

func (v *NullableConstraint) Set(val *Constraint) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraint(val *Constraint) *NullableConstraint {
	return &NullableConstraint{value: val, isSet: true}
}

func (v NullableConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



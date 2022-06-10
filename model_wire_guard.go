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

// WireGuard struct for WireGuard
type WireGuard struct {
	Id string `json:"id"`
	PrivKey string `json:"priv_key"`
	PubKey string `json:"pub_key"`
	Peers []WireGuardPeer `json:"peers"`
}

// NewWireGuard instantiates a new WireGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireGuard(id string, privKey string, pubKey string, peers []WireGuardPeer) *WireGuard {
	this := WireGuard{}
	this.Id = id
	this.PrivKey = privKey
	this.PubKey = pubKey
	this.Peers = peers
	return &this
}

// NewWireGuardWithDefaults instantiates a new WireGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireGuardWithDefaults() *WireGuard {
	this := WireGuard{}
	return &this
}

// GetId returns the Id field value
func (o *WireGuard) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WireGuard) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WireGuard) SetId(v string) {
	o.Id = v
}

// GetPrivKey returns the PrivKey field value
func (o *WireGuard) GetPrivKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivKey
}

// GetPrivKeyOk returns a tuple with the PrivKey field value
// and a boolean to check if the value has been set.
func (o *WireGuard) GetPrivKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivKey, true
}

// SetPrivKey sets field value
func (o *WireGuard) SetPrivKey(v string) {
	o.PrivKey = v
}

// GetPubKey returns the PubKey field value
func (o *WireGuard) GetPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value
// and a boolean to check if the value has been set.
func (o *WireGuard) GetPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubKey, true
}

// SetPubKey sets field value
func (o *WireGuard) SetPubKey(v string) {
	o.PubKey = v
}

// GetPeers returns the Peers field value
func (o *WireGuard) GetPeers() []WireGuardPeer {
	if o == nil {
		var ret []WireGuardPeer
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *WireGuard) GetPeersOk() ([]WireGuardPeer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *WireGuard) SetPeers(v []WireGuardPeer) {
	o.Peers = v
}

func (o WireGuard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["priv_key"] = o.PrivKey
	}
	if true {
		toSerialize["pub_key"] = o.PubKey
	}
	if true {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableWireGuard struct {
	value *WireGuard
	isSet bool
}

func (v NullableWireGuard) Get() *WireGuard {
	return v.value
}

func (v *NullableWireGuard) Set(val *WireGuard) {
	v.value = val
	v.isSet = true
}

func (v NullableWireGuard) IsSet() bool {
	return v.isSet
}

func (v *NullableWireGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireGuard(val *WireGuard) *NullableWireGuard {
	return &NullableWireGuard{value: val, isSet: true}
}

func (v NullableWireGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



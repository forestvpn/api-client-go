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

// DeviceStats struct for DeviceStats
type DeviceStats struct {
	Id *string `json:"id,omitempty"`
	Connections *int32 `json:"connections,omitempty"`
	ReceivedBytes *int32 `json:"received_bytes,omitempty"`
	TransmittedBytes *int32 `json:"transmitted_bytes,omitempty"`
	BlockedAds *int32 `json:"blocked_ads,omitempty"`
	BlockedMalwares *int32 `json:"blocked_malwares,omitempty"`
	Date *time.Time `json:"date,omitempty"`
}

// NewDeviceStats instantiates a new DeviceStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceStats() *DeviceStats {
	this := DeviceStats{}
	return &this
}

// NewDeviceStatsWithDefaults instantiates a new DeviceStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceStatsWithDefaults() *DeviceStats {
	this := DeviceStats{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceStats) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStats) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceStats) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceStats) SetId(v string) {
	o.Id = &v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *DeviceStats) GetConnections() int32 {
	if o == nil || isNil(o.Connections) {
		var ret int32
		return ret
	}
	return *o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStats) GetConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.Connections) {
    return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *DeviceStats) HasConnections() bool {
	if o != nil && !isNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given int32 and assigns it to the Connections field.
func (o *DeviceStats) SetConnections(v int32) {
	o.Connections = &v
}

// GetReceivedBytes returns the ReceivedBytes field value if set, zero value otherwise.
func (o *DeviceStats) GetReceivedBytes() int32 {
	if o == nil || isNil(o.ReceivedBytes) {
		var ret int32
		return ret
	}
	return *o.ReceivedBytes
}

// GetReceivedBytesOk returns a tuple with the ReceivedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStats) GetReceivedBytesOk() (*int32, bool) {
	if o == nil || isNil(o.ReceivedBytes) {
    return nil, false
	}
	return o.ReceivedBytes, true
}

// HasReceivedBytes returns a boolean if a field has been set.
func (o *DeviceStats) HasReceivedBytes() bool {
	if o != nil && !isNil(o.ReceivedBytes) {
		return true
	}

	return false
}

// SetReceivedBytes gets a reference to the given int32 and assigns it to the ReceivedBytes field.
func (o *DeviceStats) SetReceivedBytes(v int32) {
	o.ReceivedBytes = &v
}

// GetTransmittedBytes returns the TransmittedBytes field value if set, zero value otherwise.
func (o *DeviceStats) GetTransmittedBytes() int32 {
	if o == nil || isNil(o.TransmittedBytes) {
		var ret int32
		return ret
	}
	return *o.TransmittedBytes
}

// GetTransmittedBytesOk returns a tuple with the TransmittedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStats) GetTransmittedBytesOk() (*int32, bool) {
	if o == nil || isNil(o.TransmittedBytes) {
    return nil, false
	}
	return o.TransmittedBytes, true
}

// HasTransmittedBytes returns a boolean if a field has been set.
func (o *DeviceStats) HasTransmittedBytes() bool {
	if o != nil && !isNil(o.TransmittedBytes) {
		return true
	}

	return false
}

// SetTransmittedBytes gets a reference to the given int32 and assigns it to the TransmittedBytes field.
func (o *DeviceStats) SetTransmittedBytes(v int32) {
	o.TransmittedBytes = &v
}

// GetBlockedAds returns the BlockedAds field value if set, zero value otherwise.
func (o *DeviceStats) GetBlockedAds() int32 {
	if o == nil || isNil(o.BlockedAds) {
		var ret int32
		return ret
	}
	return *o.BlockedAds
}

// GetBlockedAdsOk returns a tuple with the BlockedAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStats) GetBlockedAdsOk() (*int32, bool) {
	if o == nil || isNil(o.BlockedAds) {
    return nil, false
	}
	return o.BlockedAds, true
}

// HasBlockedAds returns a boolean if a field has been set.
func (o *DeviceStats) HasBlockedAds() bool {
	if o != nil && !isNil(o.BlockedAds) {
		return true
	}

	return false
}

// SetBlockedAds gets a reference to the given int32 and assigns it to the BlockedAds field.
func (o *DeviceStats) SetBlockedAds(v int32) {
	o.BlockedAds = &v
}

// GetBlockedMalwares returns the BlockedMalwares field value if set, zero value otherwise.
func (o *DeviceStats) GetBlockedMalwares() int32 {
	if o == nil || isNil(o.BlockedMalwares) {
		var ret int32
		return ret
	}
	return *o.BlockedMalwares
}

// GetBlockedMalwaresOk returns a tuple with the BlockedMalwares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStats) GetBlockedMalwaresOk() (*int32, bool) {
	if o == nil || isNil(o.BlockedMalwares) {
    return nil, false
	}
	return o.BlockedMalwares, true
}

// HasBlockedMalwares returns a boolean if a field has been set.
func (o *DeviceStats) HasBlockedMalwares() bool {
	if o != nil && !isNil(o.BlockedMalwares) {
		return true
	}

	return false
}

// SetBlockedMalwares gets a reference to the given int32 and assigns it to the BlockedMalwares field.
func (o *DeviceStats) SetBlockedMalwares(v int32) {
	o.BlockedMalwares = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DeviceStats) GetDate() time.Time {
	if o == nil || isNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceStats) GetDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DeviceStats) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *DeviceStats) SetDate(v time.Time) {
	o.Date = &v
}

func (o DeviceStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !isNil(o.ReceivedBytes) {
		toSerialize["received_bytes"] = o.ReceivedBytes
	}
	if !isNil(o.TransmittedBytes) {
		toSerialize["transmitted_bytes"] = o.TransmittedBytes
	}
	if !isNil(o.BlockedAds) {
		toSerialize["blocked_ads"] = o.BlockedAds
	}
	if !isNil(o.BlockedMalwares) {
		toSerialize["blocked_malwares"] = o.BlockedMalwares
	}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceStats struct {
	value *DeviceStats
	isSet bool
}

func (v NullableDeviceStats) Get() *DeviceStats {
	return v.value
}

func (v *NullableDeviceStats) Set(val *DeviceStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceStats(val *DeviceStats) *NullableDeviceStats {
	return &NullableDeviceStats{value: val, isSet: true}
}

func (v NullableDeviceStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the AggregatedDataUsageStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregatedDataUsageStats{}

// AggregatedDataUsageStats struct for AggregatedDataUsageStats
type AggregatedDataUsageStats struct {
	// Aggregation value. It might de a hour, day, week, or month
	AggrInterval *string `json:"aggr_interval,omitempty"`
	// User device name, useful for retrieve extra data through device API
	DeviceId *string `json:"device_id,omitempty"`
	// User device name, useful for showing in the chart
	DeviceName *string `json:"device_name,omitempty"`
	// Sum of the received bytes aggregated around the device and aggr_interval
	ReceivedBytes *int32 `json:"received_bytes,omitempty"`
	// Sum of the transmitted bytes aggregated around the device and aggr_interval
	TransmittedBytes *int32 `json:"transmitted_bytes,omitempty"`
	// Sum of the received + transmitted bytes aggregated around the device and aggr_interval
	TotalBytes *int32 `json:"total_bytes,omitempty"`
	// Minutes of usage time aggregated around the device and aggr_interval
	UsageTime *int32 `json:"usage_time,omitempty"`
}

// NewAggregatedDataUsageStats instantiates a new AggregatedDataUsageStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregatedDataUsageStats() *AggregatedDataUsageStats {
	this := AggregatedDataUsageStats{}
	return &this
}

// NewAggregatedDataUsageStatsWithDefaults instantiates a new AggregatedDataUsageStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatedDataUsageStatsWithDefaults() *AggregatedDataUsageStats {
	this := AggregatedDataUsageStats{}
	return &this
}

// GetAggrInterval returns the AggrInterval field value if set, zero value otherwise.
func (o *AggregatedDataUsageStats) GetAggrInterval() string {
	if o == nil || isNil(o.AggrInterval) {
		var ret string
		return ret
	}
	return *o.AggrInterval
}

// GetAggrIntervalOk returns a tuple with the AggrInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedDataUsageStats) GetAggrIntervalOk() (*string, bool) {
	if o == nil || isNil(o.AggrInterval) {
		return nil, false
	}
	return o.AggrInterval, true
}

// HasAggrInterval returns a boolean if a field has been set.
func (o *AggregatedDataUsageStats) HasAggrInterval() bool {
	if o != nil && !isNil(o.AggrInterval) {
		return true
	}

	return false
}

// SetAggrInterval gets a reference to the given string and assigns it to the AggrInterval field.
func (o *AggregatedDataUsageStats) SetAggrInterval(v string) {
	o.AggrInterval = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *AggregatedDataUsageStats) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedDataUsageStats) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AggregatedDataUsageStats) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *AggregatedDataUsageStats) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *AggregatedDataUsageStats) GetDeviceName() string {
	if o == nil || isNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedDataUsageStats) GetDeviceNameOk() (*string, bool) {
	if o == nil || isNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *AggregatedDataUsageStats) HasDeviceName() bool {
	if o != nil && !isNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *AggregatedDataUsageStats) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetReceivedBytes returns the ReceivedBytes field value if set, zero value otherwise.
func (o *AggregatedDataUsageStats) GetReceivedBytes() int32 {
	if o == nil || isNil(o.ReceivedBytes) {
		var ret int32
		return ret
	}
	return *o.ReceivedBytes
}

// GetReceivedBytesOk returns a tuple with the ReceivedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedDataUsageStats) GetReceivedBytesOk() (*int32, bool) {
	if o == nil || isNil(o.ReceivedBytes) {
		return nil, false
	}
	return o.ReceivedBytes, true
}

// HasReceivedBytes returns a boolean if a field has been set.
func (o *AggregatedDataUsageStats) HasReceivedBytes() bool {
	if o != nil && !isNil(o.ReceivedBytes) {
		return true
	}

	return false
}

// SetReceivedBytes gets a reference to the given int32 and assigns it to the ReceivedBytes field.
func (o *AggregatedDataUsageStats) SetReceivedBytes(v int32) {
	o.ReceivedBytes = &v
}

// GetTransmittedBytes returns the TransmittedBytes field value if set, zero value otherwise.
func (o *AggregatedDataUsageStats) GetTransmittedBytes() int32 {
	if o == nil || isNil(o.TransmittedBytes) {
		var ret int32
		return ret
	}
	return *o.TransmittedBytes
}

// GetTransmittedBytesOk returns a tuple with the TransmittedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedDataUsageStats) GetTransmittedBytesOk() (*int32, bool) {
	if o == nil || isNil(o.TransmittedBytes) {
		return nil, false
	}
	return o.TransmittedBytes, true
}

// HasTransmittedBytes returns a boolean if a field has been set.
func (o *AggregatedDataUsageStats) HasTransmittedBytes() bool {
	if o != nil && !isNil(o.TransmittedBytes) {
		return true
	}

	return false
}

// SetTransmittedBytes gets a reference to the given int32 and assigns it to the TransmittedBytes field.
func (o *AggregatedDataUsageStats) SetTransmittedBytes(v int32) {
	o.TransmittedBytes = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *AggregatedDataUsageStats) GetTotalBytes() int32 {
	if o == nil || isNil(o.TotalBytes) {
		var ret int32
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedDataUsageStats) GetTotalBytesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalBytes) {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *AggregatedDataUsageStats) HasTotalBytes() bool {
	if o != nil && !isNil(o.TotalBytes) {
		return true
	}

	return false
}

// SetTotalBytes gets a reference to the given int32 and assigns it to the TotalBytes field.
func (o *AggregatedDataUsageStats) SetTotalBytes(v int32) {
	o.TotalBytes = &v
}

// GetUsageTime returns the UsageTime field value if set, zero value otherwise.
func (o *AggregatedDataUsageStats) GetUsageTime() int32 {
	if o == nil || isNil(o.UsageTime) {
		var ret int32
		return ret
	}
	return *o.UsageTime
}

// GetUsageTimeOk returns a tuple with the UsageTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedDataUsageStats) GetUsageTimeOk() (*int32, bool) {
	if o == nil || isNil(o.UsageTime) {
		return nil, false
	}
	return o.UsageTime, true
}

// HasUsageTime returns a boolean if a field has been set.
func (o *AggregatedDataUsageStats) HasUsageTime() bool {
	if o != nil && !isNil(o.UsageTime) {
		return true
	}

	return false
}

// SetUsageTime gets a reference to the given int32 and assigns it to the UsageTime field.
func (o *AggregatedDataUsageStats) SetUsageTime(v int32) {
	o.UsageTime = &v
}

func (o AggregatedDataUsageStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregatedDataUsageStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AggrInterval) {
		toSerialize["aggr_interval"] = o.AggrInterval
	}
	if !isNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !isNil(o.DeviceName) {
		toSerialize["device_name"] = o.DeviceName
	}
	if !isNil(o.ReceivedBytes) {
		toSerialize["received_bytes"] = o.ReceivedBytes
	}
	if !isNil(o.TransmittedBytes) {
		toSerialize["transmitted_bytes"] = o.TransmittedBytes
	}
	if !isNil(o.TotalBytes) {
		toSerialize["total_bytes"] = o.TotalBytes
	}
	if !isNil(o.UsageTime) {
		toSerialize["usage_time"] = o.UsageTime
	}
	return toSerialize, nil
}

type NullableAggregatedDataUsageStats struct {
	value *AggregatedDataUsageStats
	isSet bool
}

func (v NullableAggregatedDataUsageStats) Get() *AggregatedDataUsageStats {
	return v.value
}

func (v *NullableAggregatedDataUsageStats) Set(val *AggregatedDataUsageStats) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregatedDataUsageStats) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregatedDataUsageStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregatedDataUsageStats(val *AggregatedDataUsageStats) *NullableAggregatedDataUsageStats {
	return &NullableAggregatedDataUsageStats{value: val, isSet: true}
}

func (v NullableAggregatedDataUsageStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregatedDataUsageStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



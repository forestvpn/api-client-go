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
	"time"
)

// Device struct for Device
type Device struct {
	Id string `json:"id"`
	ExternalKey *string `json:"external_key,omitempty"`
	Name *string `json:"name,omitempty"`
	Ips []string `json:"ips,omitempty"`
	Dns []string `json:"dns,omitempty"`
	TorOver *bool `json:"tor_over,omitempty"`
	ConnectionMode *ConnectionMode `json:"connection_mode,omitempty"`
	Wireguard *WireGuard `json:"wireguard,omitempty"`
	Location *Location `json:"location,omitempty"`
	Servers []Server `json:"servers,omitempty"`
	LastActiveAt *time.Time `json:"last_active_at,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice(id string) *Device {
	this := Device{}
	this.Id = id
	var torOver bool = false
	this.TorOver = &torOver
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	var torOver bool = false
	this.TorOver = &torOver
	return &this
}

// GetId returns the Id field value
func (o *Device) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Device) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Device) SetId(v string) {
	o.Id = v
}

// GetExternalKey returns the ExternalKey field value if set, zero value otherwise.
func (o *Device) GetExternalKey() string {
	if o == nil || o.ExternalKey == nil {
		var ret string
		return ret
	}
	return *o.ExternalKey
}

// GetExternalKeyOk returns a tuple with the ExternalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetExternalKeyOk() (*string, bool) {
	if o == nil || o.ExternalKey == nil {
		return nil, false
	}
	return o.ExternalKey, true
}

// HasExternalKey returns a boolean if a field has been set.
func (o *Device) HasExternalKey() bool {
	if o != nil && o.ExternalKey != nil {
		return true
	}

	return false
}

// SetExternalKey gets a reference to the given string and assigns it to the ExternalKey field.
func (o *Device) SetExternalKey(v string) {
	o.ExternalKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Device) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Device) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Device) SetName(v string) {
	o.Name = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *Device) GetIps() []string {
	if o == nil || o.Ips == nil {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIpsOk() ([]string, bool) {
	if o == nil || o.Ips == nil {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *Device) HasIps() bool {
	if o != nil && o.Ips != nil {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *Device) SetIps(v []string) {
	o.Ips = v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *Device) GetDns() []string {
	if o == nil || o.Dns == nil {
		var ret []string
		return ret
	}
	return o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDnsOk() ([]string, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *Device) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given []string and assigns it to the Dns field.
func (o *Device) SetDns(v []string) {
	o.Dns = v
}

// GetTorOver returns the TorOver field value if set, zero value otherwise.
func (o *Device) GetTorOver() bool {
	if o == nil || o.TorOver == nil {
		var ret bool
		return ret
	}
	return *o.TorOver
}

// GetTorOverOk returns a tuple with the TorOver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTorOverOk() (*bool, bool) {
	if o == nil || o.TorOver == nil {
		return nil, false
	}
	return o.TorOver, true
}

// HasTorOver returns a boolean if a field has been set.
func (o *Device) HasTorOver() bool {
	if o != nil && o.TorOver != nil {
		return true
	}

	return false
}

// SetTorOver gets a reference to the given bool and assigns it to the TorOver field.
func (o *Device) SetTorOver(v bool) {
	o.TorOver = &v
}

// GetConnectionMode returns the ConnectionMode field value if set, zero value otherwise.
func (o *Device) GetConnectionMode() ConnectionMode {
	if o == nil || o.ConnectionMode == nil {
		var ret ConnectionMode
		return ret
	}
	return *o.ConnectionMode
}

// GetConnectionModeOk returns a tuple with the ConnectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetConnectionModeOk() (*ConnectionMode, bool) {
	if o == nil || o.ConnectionMode == nil {
		return nil, false
	}
	return o.ConnectionMode, true
}

// HasConnectionMode returns a boolean if a field has been set.
func (o *Device) HasConnectionMode() bool {
	if o != nil && o.ConnectionMode != nil {
		return true
	}

	return false
}

// SetConnectionMode gets a reference to the given ConnectionMode and assigns it to the ConnectionMode field.
func (o *Device) SetConnectionMode(v ConnectionMode) {
	o.ConnectionMode = &v
}

// GetWireguard returns the Wireguard field value if set, zero value otherwise.
func (o *Device) GetWireguard() WireGuard {
	if o == nil || o.Wireguard == nil {
		var ret WireGuard
		return ret
	}
	return *o.Wireguard
}

// GetWireguardOk returns a tuple with the Wireguard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetWireguardOk() (*WireGuard, bool) {
	if o == nil || o.Wireguard == nil {
		return nil, false
	}
	return o.Wireguard, true
}

// HasWireguard returns a boolean if a field has been set.
func (o *Device) HasWireguard() bool {
	if o != nil && o.Wireguard != nil {
		return true
	}

	return false
}

// SetWireguard gets a reference to the given WireGuard and assigns it to the Wireguard field.
func (o *Device) SetWireguard(v WireGuard) {
	o.Wireguard = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Device) GetLocation() Location {
	if o == nil || o.Location == nil {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLocationOk() (*Location, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Device) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *Device) SetLocation(v Location) {
	o.Location = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *Device) GetServers() []Server {
	if o == nil || o.Servers == nil {
		var ret []Server
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetServersOk() ([]Server, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *Device) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Server and assigns it to the Servers field.
func (o *Device) SetServers(v []Server) {
	o.Servers = v
}

// GetLastActiveAt returns the LastActiveAt field value if set, zero value otherwise.
func (o *Device) GetLastActiveAt() time.Time {
	if o == nil || o.LastActiveAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActiveAt
}

// GetLastActiveAtOk returns a tuple with the LastActiveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastActiveAtOk() (*time.Time, bool) {
	if o == nil || o.LastActiveAt == nil {
		return nil, false
	}
	return o.LastActiveAt, true
}

// HasLastActiveAt returns a boolean if a field has been set.
func (o *Device) HasLastActiveAt() bool {
	if o != nil && o.LastActiveAt != nil {
		return true
	}

	return false
}

// SetLastActiveAt gets a reference to the given time.Time and assigns it to the LastActiveAt field.
func (o *Device) SetLastActiveAt(v time.Time) {
	o.LastActiveAt = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ExternalKey != nil {
		toSerialize["external_key"] = o.ExternalKey
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Ips != nil {
		toSerialize["ips"] = o.Ips
	}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.TorOver != nil {
		toSerialize["tor_over"] = o.TorOver
	}
	if o.ConnectionMode != nil {
		toSerialize["connection_mode"] = o.ConnectionMode
	}
	if o.Wireguard != nil {
		toSerialize["wireguard"] = o.Wireguard
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.LastActiveAt != nil {
		toSerialize["last_active_at"] = o.LastActiveAt
	}
	return json.Marshal(toSerialize)
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



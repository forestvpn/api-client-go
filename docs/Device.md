# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExternalKey** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 
**Dns** | Pointer to **[]string** |  | [optional] 
**TorOver** | Pointer to **bool** |  | [optional] [default to false]
**ConnectionMode** | Pointer to [**ConnectionMode**](ConnectionMode.md) |  | [optional] 
**Wireguard** | Pointer to [**WireGuard**](WireGuard.md) |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Servers** | Pointer to [**[]Server**](Server.md) |  | [optional] 
**LastActiveAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDevice

`func NewDevice(id string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetExternalKey

`func (o *Device) GetExternalKey() string`

GetExternalKey returns the ExternalKey field if non-nil, zero value otherwise.

### GetExternalKeyOk

`func (o *Device) GetExternalKeyOk() (*string, bool)`

GetExternalKeyOk returns a tuple with the ExternalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKey

`func (o *Device) SetExternalKey(v string)`

SetExternalKey sets ExternalKey field to given value.

### HasExternalKey

`func (o *Device) HasExternalKey() bool`

HasExternalKey returns a boolean if a field has been set.

### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIps

`func (o *Device) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *Device) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *Device) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *Device) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetDns

`func (o *Device) GetDns() []string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *Device) GetDnsOk() (*[]string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *Device) SetDns(v []string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *Device) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetTorOver

`func (o *Device) GetTorOver() bool`

GetTorOver returns the TorOver field if non-nil, zero value otherwise.

### GetTorOverOk

`func (o *Device) GetTorOverOk() (*bool, bool)`

GetTorOverOk returns a tuple with the TorOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorOver

`func (o *Device) SetTorOver(v bool)`

SetTorOver sets TorOver field to given value.

### HasTorOver

`func (o *Device) HasTorOver() bool`

HasTorOver returns a boolean if a field has been set.

### GetConnectionMode

`func (o *Device) GetConnectionMode() ConnectionMode`

GetConnectionMode returns the ConnectionMode field if non-nil, zero value otherwise.

### GetConnectionModeOk

`func (o *Device) GetConnectionModeOk() (*ConnectionMode, bool)`

GetConnectionModeOk returns a tuple with the ConnectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMode

`func (o *Device) SetConnectionMode(v ConnectionMode)`

SetConnectionMode sets ConnectionMode field to given value.

### HasConnectionMode

`func (o *Device) HasConnectionMode() bool`

HasConnectionMode returns a boolean if a field has been set.

### GetWireguard

`func (o *Device) GetWireguard() WireGuard`

GetWireguard returns the Wireguard field if non-nil, zero value otherwise.

### GetWireguardOk

`func (o *Device) GetWireguardOk() (*WireGuard, bool)`

GetWireguardOk returns a tuple with the Wireguard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguard

`func (o *Device) SetWireguard(v WireGuard)`

SetWireguard sets Wireguard field to given value.

### HasWireguard

`func (o *Device) HasWireguard() bool`

HasWireguard returns a boolean if a field has been set.

### GetLocation

`func (o *Device) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Device) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Device) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Device) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetServers

`func (o *Device) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Device) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Device) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Device) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetLastActiveAt

`func (o *Device) GetLastActiveAt() time.Time`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *Device) GetLastActiveAtOk() (*time.Time, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *Device) SetLastActiveAt(v time.Time)`

SetLastActiveAt sets LastActiveAt field to given value.

### HasLastActiveAt

`func (o *Device) HasLastActiveAt() bool`

HasLastActiveAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WireGuardPeerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PubKey** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**WireGuardPeerUser**](WireGuardPeerUser.md) |  | [optional] 
**Device** | Pointer to [**WireGuardPeerDevice**](WireGuardPeerDevice.md) |  | [optional] 

## Methods

### NewWireGuardPeerInfo

`func NewWireGuardPeerInfo() *WireGuardPeerInfo`

NewWireGuardPeerInfo instantiates a new WireGuardPeerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireGuardPeerInfoWithDefaults

`func NewWireGuardPeerInfoWithDefaults() *WireGuardPeerInfo`

NewWireGuardPeerInfoWithDefaults instantiates a new WireGuardPeerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubKey

`func (o *WireGuardPeerInfo) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *WireGuardPeerInfo) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *WireGuardPeerInfo) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.

### HasPubKey

`func (o *WireGuardPeerInfo) HasPubKey() bool`

HasPubKey returns a boolean if a field has been set.

### GetUser

`func (o *WireGuardPeerInfo) GetUser() WireGuardPeerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WireGuardPeerInfo) GetUserOk() (*WireGuardPeerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WireGuardPeerInfo) SetUser(v WireGuardPeerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *WireGuardPeerInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDevice

`func (o *WireGuardPeerInfo) GetDevice() WireGuardPeerDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WireGuardPeerInfo) GetDeviceOk() (*WireGuardPeerDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WireGuardPeerInfo) SetDevice(v WireGuardPeerDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *WireGuardPeerInfo) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



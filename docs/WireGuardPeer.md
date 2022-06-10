# WireGuardPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PubKey** | **string** |  | 
**PsKey** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**AllowedIps** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWireGuardPeer

`func NewWireGuardPeer(pubKey string, ) *WireGuardPeer`

NewWireGuardPeer instantiates a new WireGuardPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireGuardPeerWithDefaults

`func NewWireGuardPeerWithDefaults() *WireGuardPeer`

NewWireGuardPeerWithDefaults instantiates a new WireGuardPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubKey

`func (o *WireGuardPeer) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *WireGuardPeer) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *WireGuardPeer) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetPsKey

`func (o *WireGuardPeer) GetPsKey() string`

GetPsKey returns the PsKey field if non-nil, zero value otherwise.

### GetPsKeyOk

`func (o *WireGuardPeer) GetPsKeyOk() (*string, bool)`

GetPsKeyOk returns a tuple with the PsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsKey

`func (o *WireGuardPeer) SetPsKey(v string)`

SetPsKey sets PsKey field to given value.

### HasPsKey

`func (o *WireGuardPeer) HasPsKey() bool`

HasPsKey returns a boolean if a field has been set.

### GetEndpoint

`func (o *WireGuardPeer) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WireGuardPeer) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WireGuardPeer) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WireGuardPeer) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAllowedIps

`func (o *WireGuardPeer) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *WireGuardPeer) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *WireGuardPeer) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *WireGuardPeer) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



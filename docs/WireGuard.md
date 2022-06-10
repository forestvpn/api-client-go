# WireGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PrivKey** | **string** |  | 
**PubKey** | **string** |  | 
**Peers** | [**[]WireGuardPeer**](WireGuardPeer.md) |  | 

## Methods

### NewWireGuard

`func NewWireGuard(id string, privKey string, pubKey string, peers []WireGuardPeer, ) *WireGuard`

NewWireGuard instantiates a new WireGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireGuardWithDefaults

`func NewWireGuardWithDefaults() *WireGuard`

NewWireGuardWithDefaults instantiates a new WireGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireGuard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireGuard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireGuard) SetId(v string)`

SetId sets Id field to given value.


### GetPrivKey

`func (o *WireGuard) GetPrivKey() string`

GetPrivKey returns the PrivKey field if non-nil, zero value otherwise.

### GetPrivKeyOk

`func (o *WireGuard) GetPrivKeyOk() (*string, bool)`

GetPrivKeyOk returns a tuple with the PrivKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivKey

`func (o *WireGuard) SetPrivKey(v string)`

SetPrivKey sets PrivKey field to given value.


### GetPubKey

`func (o *WireGuard) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *WireGuard) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *WireGuard) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.


### GetPeers

`func (o *WireGuard) GetPeers() []WireGuardPeer`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *WireGuard) GetPeersOk() (*[]WireGuardPeer, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *WireGuard) SetPeers(v []WireGuardPeer)`

SetPeers sets Peers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



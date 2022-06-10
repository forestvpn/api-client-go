# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**NetworkServices** | [**[]NetworkService**](NetworkService.md) |  | 

## Methods

### NewServer

`func NewServer(host string, networkServices []NetworkService, ) *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *Server) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Server) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Server) SetHost(v string)`

SetHost sets Host field to given value.


### GetNetworkServices

`func (o *Server) GetNetworkServices() []NetworkService`

GetNetworkServices returns the NetworkServices field if non-nil, zero value otherwise.

### GetNetworkServicesOk

`func (o *Server) GetNetworkServicesOk() (*[]NetworkService, bool)`

GetNetworkServicesOk returns a tuple with the NetworkServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServices

`func (o *Server) SetNetworkServices(v []NetworkService)`

SetNetworkServices sets NetworkServices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



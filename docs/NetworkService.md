# NetworkService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proto** | **string** |  | 
**ConnectionUri** | **string** |  | 

## Methods

### NewNetworkService

`func NewNetworkService(proto string, connectionUri string, ) *NetworkService`

NewNetworkService instantiates a new NetworkService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServiceWithDefaults

`func NewNetworkServiceWithDefaults() *NetworkService`

NewNetworkServiceWithDefaults instantiates a new NetworkService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProto

`func (o *NetworkService) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *NetworkService) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *NetworkService) SetProto(v string)`

SetProto sets Proto field to given value.


### GetConnectionUri

`func (o *NetworkService) GetConnectionUri() string`

GetConnectionUri returns the ConnectionUri field if non-nil, zero value otherwise.

### GetConnectionUriOk

`func (o *NetworkService) GetConnectionUriOk() (*string, bool)`

GetConnectionUriOk returns a tuple with the ConnectionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUri

`func (o *NetworkService) SetConnectionUri(v string)`

SetConnectionUri sets ConnectionUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



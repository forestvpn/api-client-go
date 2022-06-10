# CloudPaymentsAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **int32** |  | [optional] 
**Secure3d** | Pointer to [**CloudPaymentsSecure3d**](CloudPaymentsSecure3d.md) |  | [optional] 

## Methods

### NewCloudPaymentsAuth

`func NewCloudPaymentsAuth() *CloudPaymentsAuth`

NewCloudPaymentsAuth instantiates a new CloudPaymentsAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPaymentsAuthWithDefaults

`func NewCloudPaymentsAuthWithDefaults() *CloudPaymentsAuth`

NewCloudPaymentsAuthWithDefaults instantiates a new CloudPaymentsAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *CloudPaymentsAuth) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CloudPaymentsAuth) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CloudPaymentsAuth) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CloudPaymentsAuth) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetSecure3d

`func (o *CloudPaymentsAuth) GetSecure3d() CloudPaymentsSecure3d`

GetSecure3d returns the Secure3d field if non-nil, zero value otherwise.

### GetSecure3dOk

`func (o *CloudPaymentsAuth) GetSecure3dOk() (*CloudPaymentsSecure3d, bool)`

GetSecure3dOk returns a tuple with the Secure3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure3d

`func (o *CloudPaymentsAuth) SetSecure3d(v CloudPaymentsSecure3d)`

SetSecure3d sets Secure3d field to given value.

### HasSecure3d

`func (o *CloudPaymentsAuth) HasSecure3d() bool`

HasSecure3d returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



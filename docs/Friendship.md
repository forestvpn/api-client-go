# Friendship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewFriendship

`func NewFriendship(id string, createdAt time.Time, ) *Friendship`

NewFriendship instantiates a new Friendship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFriendshipWithDefaults

`func NewFriendshipWithDefaults() *Friendship`

NewFriendshipWithDefaults instantiates a new Friendship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Friendship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Friendship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Friendship) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *Friendship) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Friendship) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Friendship) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Friendship) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Friendship) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Friendship) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Friendship) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



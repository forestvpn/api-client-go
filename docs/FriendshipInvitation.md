# FriendshipInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**User** | [**User**](User.md) |  | 
**ShareText** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewFriendshipInvitation

`func NewFriendshipInvitation(code string, user User, shareText string, createdAt time.Time, ) *FriendshipInvitation`

NewFriendshipInvitation instantiates a new FriendshipInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFriendshipInvitationWithDefaults

`func NewFriendshipInvitationWithDefaults() *FriendshipInvitation`

NewFriendshipInvitationWithDefaults instantiates a new FriendshipInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *FriendshipInvitation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FriendshipInvitation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FriendshipInvitation) SetCode(v string)`

SetCode sets Code field to given value.


### GetUser

`func (o *FriendshipInvitation) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FriendshipInvitation) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FriendshipInvitation) SetUser(v User)`

SetUser sets User field to given value.


### GetShareText

`func (o *FriendshipInvitation) GetShareText() string`

GetShareText returns the ShareText field if non-nil, zero value otherwise.

### GetShareTextOk

`func (o *FriendshipInvitation) GetShareTextOk() (*string, bool)`

GetShareTextOk returns a tuple with the ShareText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareText

`func (o *FriendshipInvitation) SetShareText(v string)`

SetShareText sets ShareText field to given value.


### GetCreatedAt

`func (o *FriendshipInvitation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FriendshipInvitation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FriendshipInvitation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



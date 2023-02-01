# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Slug** | **int32** |  | 
**Title** | **string** |  | 
**Description** | **string** |  | 
**Unread** | **bool** |  | 
**Type** | **string** |  | 
**Level** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**ActorContentType** | Pointer to **int32** |  | [optional] 
**ActorObjectId** | Pointer to **string** |  | [optional] 
**Verb** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 

## Methods

### NewNotification

`func NewNotification(id int32, slug int32, title string, description string, unread bool, type_ string, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Notification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notification) SetId(v int32)`

SetId sets Id field to given value.


### GetSlug

`func (o *Notification) GetSlug() int32`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Notification) GetSlugOk() (*int32, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Notification) SetSlug(v int32)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *Notification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Notification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Notification) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Notification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Notification) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Notification) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnread

`func (o *Notification) GetUnread() bool`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *Notification) GetUnreadOk() (*bool, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *Notification) SetUnread(v bool)`

SetUnread sets Unread field to given value.


### GetType

`func (o *Notification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v string)`

SetType sets Type field to given value.


### GetLevel

`func (o *Notification) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Notification) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Notification) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Notification) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetRecipient

`func (o *Notification) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Notification) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Notification) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Notification) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetActorContentType

`func (o *Notification) GetActorContentType() int32`

GetActorContentType returns the ActorContentType field if non-nil, zero value otherwise.

### GetActorContentTypeOk

`func (o *Notification) GetActorContentTypeOk() (*int32, bool)`

GetActorContentTypeOk returns a tuple with the ActorContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorContentType

`func (o *Notification) SetActorContentType(v int32)`

SetActorContentType sets ActorContentType field to given value.

### HasActorContentType

`func (o *Notification) HasActorContentType() bool`

HasActorContentType returns a boolean if a field has been set.

### GetActorObjectId

`func (o *Notification) GetActorObjectId() string`

GetActorObjectId returns the ActorObjectId field if non-nil, zero value otherwise.

### GetActorObjectIdOk

`func (o *Notification) GetActorObjectIdOk() (*string, bool)`

GetActorObjectIdOk returns a tuple with the ActorObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorObjectId

`func (o *Notification) SetActorObjectId(v string)`

SetActorObjectId sets ActorObjectId field to given value.

### HasActorObjectId

`func (o *Notification) HasActorObjectId() bool`

HasActorObjectId returns a boolean if a field has been set.

### GetVerb

`func (o *Notification) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *Notification) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *Notification) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *Notification) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Notification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Notification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Notification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Notification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPublic

`func (o *Notification) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Notification) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Notification) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *Notification) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetDeleted

`func (o *Notification) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Notification) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Notification) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Notification) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetData

`func (o *Notification) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Notification) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Notification) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Notification) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



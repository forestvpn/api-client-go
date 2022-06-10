# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Title** | **string** |  | 
**Summary** | **string** |  | 
**FeaturedImage** | [**FeaturedImage**](FeaturedImage.md) |  | 
**IsUnread** | **bool** |  | 
**IsFeatured** | **bool** |  | 

## Methods

### NewNotification

`func NewNotification(id int32, createdAt time.Time, title string, summary string, featuredImage FeaturedImage, isUnread bool, isFeatured bool, ) *Notification`

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


### GetSummary

`func (o *Notification) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Notification) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Notification) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetFeaturedImage

`func (o *Notification) GetFeaturedImage() FeaturedImage`

GetFeaturedImage returns the FeaturedImage field if non-nil, zero value otherwise.

### GetFeaturedImageOk

`func (o *Notification) GetFeaturedImageOk() (*FeaturedImage, bool)`

GetFeaturedImageOk returns a tuple with the FeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImage

`func (o *Notification) SetFeaturedImage(v FeaturedImage)`

SetFeaturedImage sets FeaturedImage field to given value.


### GetIsUnread

`func (o *Notification) GetIsUnread() bool`

GetIsUnread returns the IsUnread field if non-nil, zero value otherwise.

### GetIsUnreadOk

`func (o *Notification) GetIsUnreadOk() (*bool, bool)`

GetIsUnreadOk returns a tuple with the IsUnread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnread

`func (o *Notification) SetIsUnread(v bool)`

SetIsUnread sets IsUnread field to given value.


### GetIsFeatured

`func (o *Notification) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *Notification) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *Notification) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



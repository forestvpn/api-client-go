# NotificationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Title** | **string** |  | 
**Summary** | **string** |  | 
**FeaturedImage** | [**FeaturedImage**](FeaturedImage.md) |  | 
**Content** | **string** |  | 
**IsUnread** | **bool** |  | 
**IsFeatured** | **bool** |  | 

## Methods

### NewNotificationDetail

`func NewNotificationDetail(id int32, createdAt time.Time, title string, summary string, featuredImage FeaturedImage, content string, isUnread bool, isFeatured bool, ) *NotificationDetail`

NewNotificationDetail instantiates a new NotificationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDetailWithDefaults

`func NewNotificationDetailWithDefaults() *NotificationDetail`

NewNotificationDetailWithDefaults instantiates a new NotificationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationDetail) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *NotificationDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTitle

`func (o *NotificationDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationDetail) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *NotificationDetail) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *NotificationDetail) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *NotificationDetail) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetFeaturedImage

`func (o *NotificationDetail) GetFeaturedImage() FeaturedImage`

GetFeaturedImage returns the FeaturedImage field if non-nil, zero value otherwise.

### GetFeaturedImageOk

`func (o *NotificationDetail) GetFeaturedImageOk() (*FeaturedImage, bool)`

GetFeaturedImageOk returns a tuple with the FeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImage

`func (o *NotificationDetail) SetFeaturedImage(v FeaturedImage)`

SetFeaturedImage sets FeaturedImage field to given value.


### GetContent

`func (o *NotificationDetail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NotificationDetail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NotificationDetail) SetContent(v string)`

SetContent sets Content field to given value.


### GetIsUnread

`func (o *NotificationDetail) GetIsUnread() bool`

GetIsUnread returns the IsUnread field if non-nil, zero value otherwise.

### GetIsUnreadOk

`func (o *NotificationDetail) GetIsUnreadOk() (*bool, bool)`

GetIsUnreadOk returns a tuple with the IsUnread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnread

`func (o *NotificationDetail) SetIsUnread(v bool)`

SetIsUnread sets IsUnread field to given value.


### GetIsFeatured

`func (o *NotificationDetail) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *NotificationDetail) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *NotificationDetail) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



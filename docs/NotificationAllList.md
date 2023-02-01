# NotificationAllList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnreadCount** | Pointer to **int32** |  | [optional] 
**AllList** | [**[]Notification**](Notification.md) |  | 

## Methods

### NewNotificationAllList

`func NewNotificationAllList(allList []Notification, ) *NotificationAllList`

NewNotificationAllList instantiates a new NotificationAllList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationAllListWithDefaults

`func NewNotificationAllListWithDefaults() *NotificationAllList`

NewNotificationAllListWithDefaults instantiates a new NotificationAllList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnreadCount

`func (o *NotificationAllList) GetUnreadCount() int32`

GetUnreadCount returns the UnreadCount field if non-nil, zero value otherwise.

### GetUnreadCountOk

`func (o *NotificationAllList) GetUnreadCountOk() (*int32, bool)`

GetUnreadCountOk returns a tuple with the UnreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadCount

`func (o *NotificationAllList) SetUnreadCount(v int32)`

SetUnreadCount sets UnreadCount field to given value.

### HasUnreadCount

`func (o *NotificationAllList) HasUnreadCount() bool`

HasUnreadCount returns a boolean if a field has been set.

### GetAllList

`func (o *NotificationAllList) GetAllList() []Notification`

GetAllList returns the AllList field if non-nil, zero value otherwise.

### GetAllListOk

`func (o *NotificationAllList) GetAllListOk() (*[]Notification, bool)`

GetAllListOk returns a tuple with the AllList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllList

`func (o *NotificationAllList) SetAllList(v []Notification)`

SetAllList sets AllList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



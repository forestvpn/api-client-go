# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**PromotionalCode** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**SubscriptionSource**](SubscriptionSource.md) |  | [optional] [readonly] 
**StartDate** | **time.Time** |  | [readonly] 
**EndedAt** | Pointer to **time.Time** | If the subscription has ended, the date the subscription ended. | [optional] [readonly] 
**CurrentPeriodStart** | **time.Time** | Start of the current period that the subscription has been invoiced for. | [readonly] 
**CurrentPeriodEnd** | **time.Time** | End of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created. | [readonly] 
**TrialStart** | Pointer to **time.Time** | If the subscription has a trial, the beginning of that trial. | [optional] [readonly] 
**TrialEnd** | Pointer to **time.Time** | If the subscription has a trial, the end of that trial. | [optional] [readonly] 
**CancelAt** | Pointer to **time.Time** | A date in the future at which the subscription will automatically get canceled. | [optional] [readonly] 
**CanceledAt** | Pointer to **time.Time** | If the subscription has been canceled, the date of that cancellation. | [optional] [readonly] 
**Status** | [**SubscriptionStatus**](SubscriptionStatus.md) |  | [readonly] 
**Description** | Pointer to **string** | The subscription’s description, meant to be displayable to the customer. | [optional] [readonly] 
**Items** | Pointer to [**[]SubscriptionItem**](SubscriptionItem.md) |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription(id string, startDate time.Time, currentPeriodStart time.Time, currentPeriodEnd time.Time, status SubscriptionStatus, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.


### GetPromotionalCode

`func (o *Subscription) GetPromotionalCode() string`

GetPromotionalCode returns the PromotionalCode field if non-nil, zero value otherwise.

### GetPromotionalCodeOk

`func (o *Subscription) GetPromotionalCodeOk() (*string, bool)`

GetPromotionalCodeOk returns a tuple with the PromotionalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionalCode

`func (o *Subscription) SetPromotionalCode(v string)`

SetPromotionalCode sets PromotionalCode field to given value.

### HasPromotionalCode

`func (o *Subscription) HasPromotionalCode() bool`

HasPromotionalCode returns a boolean if a field has been set.

### GetSource

`func (o *Subscription) GetSource() SubscriptionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Subscription) GetSourceOk() (*SubscriptionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Subscription) SetSource(v SubscriptionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Subscription) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStartDate

`func (o *Subscription) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Subscription) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Subscription) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndedAt

`func (o *Subscription) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Subscription) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Subscription) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Subscription) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *Subscription) GetCurrentPeriodStart() time.Time`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *Subscription) GetCurrentPeriodStartOk() (*time.Time, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *Subscription) SetCurrentPeriodStart(v time.Time)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.


### GetCurrentPeriodEnd

`func (o *Subscription) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *Subscription) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *Subscription) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### GetTrialStart

`func (o *Subscription) GetTrialStart() time.Time`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *Subscription) GetTrialStartOk() (*time.Time, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *Subscription) SetTrialStart(v time.Time)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *Subscription) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### GetTrialEnd

`func (o *Subscription) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *Subscription) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *Subscription) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *Subscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetCancelAt

`func (o *Subscription) GetCancelAt() time.Time`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *Subscription) GetCancelAtOk() (*time.Time, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *Subscription) SetCancelAt(v time.Time)`

SetCancelAt sets CancelAt field to given value.

### HasCancelAt

`func (o *Subscription) HasCancelAt() bool`

HasCancelAt returns a boolean if a field has been set.

### GetCanceledAt

`func (o *Subscription) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Subscription) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Subscription) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Subscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *Subscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetItems

`func (o *Subscription) GetItems() []SubscriptionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Subscription) GetItemsOk() (*[]SubscriptionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Subscription) SetItems(v []SubscriptionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Subscription) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



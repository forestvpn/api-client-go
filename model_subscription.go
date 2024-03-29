/*
ForestVPN API

ForestVPN - Fast, secure, and modern VPN. It offers Distributed Computing, Crypto Mining, P2P, Ad Blocking, TOR over VPN, 30+ locations, and a free version with unlimited data. 

API version: 2.0
Contact: support@forestvpn.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forestvpn_api

import (
	"encoding/json"
	"time"
)

// checks if the Subscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscription{}

// Subscription struct for Subscription
type Subscription struct {
	Id string `json:"id"`
	PromotionalCode *string `json:"promotional_code,omitempty"`
	Source *SubscriptionSource `json:"source,omitempty"`
	StartDate time.Time `json:"start_date"`
	// If the subscription has ended, the date the subscription ended.
	EndedAt *time.Time `json:"ended_at,omitempty"`
	// Start of the current period that the subscription has been invoiced for.
	CurrentPeriodStart time.Time `json:"current_period_start"`
	// End of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created.
	CurrentPeriodEnd time.Time `json:"current_period_end"`
	// If the subscription has a trial, the beginning of that trial.
	TrialStart *time.Time `json:"trial_start,omitempty"`
	// If the subscription has a trial, the end of that trial.
	TrialEnd *time.Time `json:"trial_end,omitempty"`
	// A date in the future at which the subscription will automatically get canceled.
	CancelAt *time.Time `json:"cancel_at,omitempty"`
	// If the subscription has been canceled, the date of that cancellation.
	CanceledAt *time.Time `json:"canceled_at,omitempty"`
	Status SubscriptionStatus `json:"status"`
	// The subscription’s description, meant to be displayable to the customer.
	Description *string `json:"description,omitempty"`
	Items []SubscriptionItem `json:"items,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(id string, startDate time.Time, currentPeriodStart time.Time, currentPeriodEnd time.Time, status SubscriptionStatus) *Subscription {
	this := Subscription{}
	this.Id = id
	this.StartDate = startDate
	this.CurrentPeriodStart = currentPeriodStart
	this.CurrentPeriodEnd = currentPeriodEnd
	this.Status = status
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetId returns the Id field value
func (o *Subscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Subscription) SetId(v string) {
	o.Id = v
}

// GetPromotionalCode returns the PromotionalCode field value if set, zero value otherwise.
func (o *Subscription) GetPromotionalCode() string {
	if o == nil || isNil(o.PromotionalCode) {
		var ret string
		return ret
	}
	return *o.PromotionalCode
}

// GetPromotionalCodeOk returns a tuple with the PromotionalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPromotionalCodeOk() (*string, bool) {
	if o == nil || isNil(o.PromotionalCode) {
		return nil, false
	}
	return o.PromotionalCode, true
}

// HasPromotionalCode returns a boolean if a field has been set.
func (o *Subscription) HasPromotionalCode() bool {
	if o != nil && !isNil(o.PromotionalCode) {
		return true
	}

	return false
}

// SetPromotionalCode gets a reference to the given string and assigns it to the PromotionalCode field.
func (o *Subscription) SetPromotionalCode(v string) {
	o.PromotionalCode = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Subscription) GetSource() SubscriptionSource {
	if o == nil || isNil(o.Source) {
		var ret SubscriptionSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSourceOk() (*SubscriptionSource, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Subscription) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given SubscriptionSource and assigns it to the Source field.
func (o *Subscription) SetSource(v SubscriptionSource) {
	o.Source = &v
}

// GetStartDate returns the StartDate field value
func (o *Subscription) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *Subscription) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *Subscription) GetEndedAt() time.Time {
	if o == nil || isNil(o.EndedAt) {
		var ret time.Time
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetEndedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndedAt) {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *Subscription) HasEndedAt() bool {
	if o != nil && !isNil(o.EndedAt) {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given time.Time and assigns it to the EndedAt field.
func (o *Subscription) SetEndedAt(v time.Time) {
	o.EndedAt = &v
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value
func (o *Subscription) GetCurrentPeriodStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CurrentPeriodStart
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrentPeriodStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodStart, true
}

// SetCurrentPeriodStart sets field value
func (o *Subscription) SetCurrentPeriodStart(v time.Time) {
	o.CurrentPeriodStart = v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value
func (o *Subscription) GetCurrentPeriodEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrentPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodEnd, true
}

// SetCurrentPeriodEnd sets field value
func (o *Subscription) SetCurrentPeriodEnd(v time.Time) {
	o.CurrentPeriodEnd = v
}

// GetTrialStart returns the TrialStart field value if set, zero value otherwise.
func (o *Subscription) GetTrialStart() time.Time {
	if o == nil || isNil(o.TrialStart) {
		var ret time.Time
		return ret
	}
	return *o.TrialStart
}

// GetTrialStartOk returns a tuple with the TrialStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTrialStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.TrialStart) {
		return nil, false
	}
	return o.TrialStart, true
}

// HasTrialStart returns a boolean if a field has been set.
func (o *Subscription) HasTrialStart() bool {
	if o != nil && !isNil(o.TrialStart) {
		return true
	}

	return false
}

// SetTrialStart gets a reference to the given time.Time and assigns it to the TrialStart field.
func (o *Subscription) SetTrialStart(v time.Time) {
	o.TrialStart = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *Subscription) GetTrialEnd() time.Time {
	if o == nil || isNil(o.TrialEnd) {
		var ret time.Time
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTrialEndOk() (*time.Time, bool) {
	if o == nil || isNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *Subscription) HasTrialEnd() bool {
	if o != nil && !isNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given time.Time and assigns it to the TrialEnd field.
func (o *Subscription) SetTrialEnd(v time.Time) {
	o.TrialEnd = &v
}

// GetCancelAt returns the CancelAt field value if set, zero value otherwise.
func (o *Subscription) GetCancelAt() time.Time {
	if o == nil || isNil(o.CancelAt) {
		var ret time.Time
		return ret
	}
	return *o.CancelAt
}

// GetCancelAtOk returns a tuple with the CancelAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCancelAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CancelAt) {
		return nil, false
	}
	return o.CancelAt, true
}

// HasCancelAt returns a boolean if a field has been set.
func (o *Subscription) HasCancelAt() bool {
	if o != nil && !isNil(o.CancelAt) {
		return true
	}

	return false
}

// SetCancelAt gets a reference to the given time.Time and assigns it to the CancelAt field.
func (o *Subscription) SetCancelAt(v time.Time) {
	o.CancelAt = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *Subscription) GetCanceledAt() time.Time {
	if o == nil || isNil(o.CanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CanceledAt) {
		return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *Subscription) HasCanceledAt() bool {
	if o != nil && !isNil(o.CanceledAt) {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given time.Time and assigns it to the CanceledAt field.
func (o *Subscription) SetCanceledAt(v time.Time) {
	o.CanceledAt = &v
}

// GetStatus returns the Status field value
func (o *Subscription) GetStatus() SubscriptionStatus {
	if o == nil {
		var ret SubscriptionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetStatusOk() (*SubscriptionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Subscription) SetStatus(v SubscriptionStatus) {
	o.Status = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Subscription) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Subscription) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Subscription) SetDescription(v string) {
	o.Description = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Subscription) GetItems() []SubscriptionItem {
	if o == nil || isNil(o.Items) {
		var ret []SubscriptionItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetItemsOk() ([]SubscriptionItem, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Subscription) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SubscriptionItem and assigns it to the Items field.
func (o *Subscription) SetItems(v []SubscriptionItem) {
	o.Items = v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !isNil(o.PromotionalCode) {
		toSerialize["promotional_code"] = o.PromotionalCode
	}
	// skip: source is readOnly
	// skip: start_date is readOnly
	// skip: ended_at is readOnly
	// skip: current_period_start is readOnly
	// skip: current_period_end is readOnly
	// skip: trial_start is readOnly
	// skip: trial_end is readOnly
	// skip: cancel_at is readOnly
	// skip: canceled_at is readOnly
	// skip: status is readOnly
	// skip: description is readOnly
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



# BillingFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | **string** | Billing feature&#39;s slug | 
**ExpiryDate** | Pointer to **time.Time** | Billing feature&#39;s expiry date | [optional] 
**Constraints** | Pointer to [**[]Constraint**](Constraint.md) |  | [optional] 

## Methods

### NewBillingFeature

`func NewBillingFeature(bundleId string, ) *BillingFeature`

NewBillingFeature instantiates a new BillingFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingFeatureWithDefaults

`func NewBillingFeatureWithDefaults() *BillingFeature`

NewBillingFeatureWithDefaults instantiates a new BillingFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *BillingFeature) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *BillingFeature) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *BillingFeature) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetExpiryDate

`func (o *BillingFeature) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *BillingFeature) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *BillingFeature) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *BillingFeature) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetConstraints

`func (o *BillingFeature) GetConstraints() []Constraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *BillingFeature) GetConstraintsOk() (*[]Constraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *BillingFeature) SetConstraints(v []Constraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *BillingFeature) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



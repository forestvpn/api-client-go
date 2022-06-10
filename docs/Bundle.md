# Bundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TrialPeriod** | Pointer to **string** | Trial period duration in ISO 8601 format. | [optional] 
**Products** | Pointer to [**[]Product**](Product.md) |  | [optional] 

## Methods

### NewBundle

`func NewBundle(id string, name string, ) *Bundle`

NewBundle instantiates a new Bundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleWithDefaults

`func NewBundleWithDefaults() *Bundle`

NewBundleWithDefaults instantiates a new Bundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bundle) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Bundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bundle) SetName(v string)`

SetName sets Name field to given value.


### GetTrialPeriod

`func (o *Bundle) GetTrialPeriod() string`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *Bundle) GetTrialPeriodOk() (*string, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *Bundle) SetTrialPeriod(v string)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *Bundle) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetProducts

`func (o *Bundle) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Bundle) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Bundle) SetProducts(v []Product)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *Bundle) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



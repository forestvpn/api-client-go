# CheckoutSessionProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**Product**](Product.md) |  | 
**Quantity** | **int32** |  | 

## Methods

### NewCheckoutSessionProduct

`func NewCheckoutSessionProduct(product Product, quantity int32, ) *CheckoutSessionProduct`

NewCheckoutSessionProduct instantiates a new CheckoutSessionProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionProductWithDefaults

`func NewCheckoutSessionProductWithDefaults() *CheckoutSessionProduct`

NewCheckoutSessionProductWithDefaults instantiates a new CheckoutSessionProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *CheckoutSessionProduct) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CheckoutSessionProduct) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CheckoutSessionProduct) SetProduct(v Product)`

SetProduct sets Product field to given value.


### GetQuantity

`func (o *CheckoutSessionProduct) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CheckoutSessionProduct) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CheckoutSessionProduct) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



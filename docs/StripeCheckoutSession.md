# StripeCheckoutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** |  | 
**SuccessUrl** | **string** |  | 
**CancelUrl** | **string** |  | 

## Methods

### NewStripeCheckoutSession

`func NewStripeCheckoutSession(sessionId string, successUrl string, cancelUrl string, ) *StripeCheckoutSession`

NewStripeCheckoutSession instantiates a new StripeCheckoutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeCheckoutSessionWithDefaults

`func NewStripeCheckoutSessionWithDefaults() *StripeCheckoutSession`

NewStripeCheckoutSessionWithDefaults instantiates a new StripeCheckoutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *StripeCheckoutSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *StripeCheckoutSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *StripeCheckoutSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetSuccessUrl

`func (o *StripeCheckoutSession) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *StripeCheckoutSession) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *StripeCheckoutSession) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### GetCancelUrl

`func (o *StripeCheckoutSession) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *StripeCheckoutSession) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *StripeCheckoutSession) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



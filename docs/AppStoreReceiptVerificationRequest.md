# AppStoreReceiptVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receipt** | **string** | The latest Base64-encoded transaction receipt. | 
**Price** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewAppStoreReceiptVerificationRequest

`func NewAppStoreReceiptVerificationRequest(receipt string, ) *AppStoreReceiptVerificationRequest`

NewAppStoreReceiptVerificationRequest instantiates a new AppStoreReceiptVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreReceiptVerificationRequestWithDefaults

`func NewAppStoreReceiptVerificationRequestWithDefaults() *AppStoreReceiptVerificationRequest`

NewAppStoreReceiptVerificationRequestWithDefaults instantiates a new AppStoreReceiptVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceipt

`func (o *AppStoreReceiptVerificationRequest) GetReceipt() string`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *AppStoreReceiptVerificationRequest) GetReceiptOk() (*string, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *AppStoreReceiptVerificationRequest) SetReceipt(v string)`

SetReceipt sets Receipt field to given value.


### GetPrice

`func (o *AppStoreReceiptVerificationRequest) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AppStoreReceiptVerificationRequest) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AppStoreReceiptVerificationRequest) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AppStoreReceiptVerificationRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *AppStoreReceiptVerificationRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AppStoreReceiptVerificationRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AppStoreReceiptVerificationRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AppStoreReceiptVerificationRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



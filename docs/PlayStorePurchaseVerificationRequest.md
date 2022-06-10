# PlayStorePurchaseVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductSku** | **string** | The inapp product SKU (for example, &#39;com.some.thing.inapp1&#39;). | 
**PurchaseToken** | **string** | The token provided to the user&#39;s device when the inapp product was purchased. | 

## Methods

### NewPlayStorePurchaseVerificationRequest

`func NewPlayStorePurchaseVerificationRequest(productSku string, purchaseToken string, ) *PlayStorePurchaseVerificationRequest`

NewPlayStorePurchaseVerificationRequest instantiates a new PlayStorePurchaseVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayStorePurchaseVerificationRequestWithDefaults

`func NewPlayStorePurchaseVerificationRequestWithDefaults() *PlayStorePurchaseVerificationRequest`

NewPlayStorePurchaseVerificationRequestWithDefaults instantiates a new PlayStorePurchaseVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductSku

`func (o *PlayStorePurchaseVerificationRequest) GetProductSku() string`

GetProductSku returns the ProductSku field if non-nil, zero value otherwise.

### GetProductSkuOk

`func (o *PlayStorePurchaseVerificationRequest) GetProductSkuOk() (*string, bool)`

GetProductSkuOk returns a tuple with the ProductSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSku

`func (o *PlayStorePurchaseVerificationRequest) SetProductSku(v string)`

SetProductSku sets ProductSku field to given value.


### GetPurchaseToken

`func (o *PlayStorePurchaseVerificationRequest) GetPurchaseToken() string`

GetPurchaseToken returns the PurchaseToken field if non-nil, zero value otherwise.

### GetPurchaseTokenOk

`func (o *PlayStorePurchaseVerificationRequest) GetPurchaseTokenOk() (*string, bool)`

GetPurchaseTokenOk returns a tuple with the PurchaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseToken

`func (o *PlayStorePurchaseVerificationRequest) SetPurchaseToken(v string)`

SetPurchaseToken sets PurchaseToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



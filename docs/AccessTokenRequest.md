# AccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**UserAgent** | **string** |  | [readonly] 
**Token** | Pointer to **string** |  | [optional] [readonly] 
**Username** | **string** |  | [readonly] 
**Status** | **string** |  | 

## Methods

### NewAccessTokenRequest

`func NewAccessTokenRequest(id string, userAgent string, username string, status string, ) *AccessTokenRequest`

NewAccessTokenRequest instantiates a new AccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenRequestWithDefaults

`func NewAccessTokenRequestWithDefaults() *AccessTokenRequest`

NewAccessTokenRequestWithDefaults instantiates a new AccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessTokenRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokenRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokenRequest) SetId(v string)`

SetId sets Id field to given value.


### GetUserAgent

`func (o *AccessTokenRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AccessTokenRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AccessTokenRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetToken

`func (o *AccessTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccessTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccessTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccessTokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *AccessTokenRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccessTokenRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccessTokenRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetStatus

`func (o *AccessTokenRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessTokenRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessTokenRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



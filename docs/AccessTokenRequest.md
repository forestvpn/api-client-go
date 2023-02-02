# AccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** | It might be empty string | [optional] 
**UserAgent** | [**UserAgent**](UserAgent.md) |  | [readonly] 
**AccessToken** | Pointer to **string** |  | [optional] [readonly] 
**Status** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ExpiresAt** | **time.Time** |  | [readonly] 

## Methods

### NewAccessTokenRequest

`func NewAccessTokenRequest(id string, userAgent UserAgent, status string, createdAt time.Time, expiresAt time.Time, ) *AccessTokenRequest`

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


### GetName

`func (o *AccessTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessTokenRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AccessTokenRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccessTokenRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserAgent

`func (o *AccessTokenRequest) GetUserAgent() UserAgent`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AccessTokenRequest) GetUserAgentOk() (*UserAgent, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AccessTokenRequest) SetUserAgent(v UserAgent)`

SetUserAgent sets UserAgent field to given value.


### GetAccessToken

`func (o *AccessTokenRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccessTokenRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccessTokenRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AccessTokenRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

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


### GetCreatedAt

`func (o *AccessTokenRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessTokenRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessTokenRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *AccessTokenRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AccessTokenRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AccessTokenRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



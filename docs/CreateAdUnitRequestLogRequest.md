# CreateAdUnitRequestLogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** |  | 
**RequestDate** | **time.Time** |  | 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Duration** | **string** | ISO-8601 duration format | 

## Methods

### NewCreateAdUnitRequestLogRequest

`func NewCreateAdUnitRequestLogRequest(unit string, requestDate time.Time, duration string, ) *CreateAdUnitRequestLogRequest`

NewCreateAdUnitRequestLogRequest instantiates a new CreateAdUnitRequestLogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdUnitRequestLogRequestWithDefaults

`func NewCreateAdUnitRequestLogRequestWithDefaults() *CreateAdUnitRequestLogRequest`

NewCreateAdUnitRequestLogRequestWithDefaults instantiates a new CreateAdUnitRequestLogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *CreateAdUnitRequestLogRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CreateAdUnitRequestLogRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CreateAdUnitRequestLogRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetRequestDate

`func (o *CreateAdUnitRequestLogRequest) GetRequestDate() time.Time`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *CreateAdUnitRequestLogRequest) GetRequestDateOk() (*time.Time, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *CreateAdUnitRequestLogRequest) SetRequestDate(v time.Time)`

SetRequestDate sets RequestDate field to given value.


### GetErrorCode

`func (o *CreateAdUnitRequestLogRequest) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateAdUnitRequestLogRequest) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateAdUnitRequestLogRequest) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CreateAdUnitRequestLogRequest) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CreateAdUnitRequestLogRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CreateAdUnitRequestLogRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CreateAdUnitRequestLogRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CreateAdUnitRequestLogRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDuration

`func (o *CreateAdUnitRequestLogRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CreateAdUnitRequestLogRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CreateAdUnitRequestLogRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



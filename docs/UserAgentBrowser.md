# UserAgentBrowser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** | It might be \&quot;Other\&quot; in case if it can&#39;t be recognized | [optional] 
**Version** | Pointer to **string** | It might be empty string | [optional] 

## Methods

### NewUserAgentBrowser

`func NewUserAgentBrowser() *UserAgentBrowser`

NewUserAgentBrowser instantiates a new UserAgentBrowser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentBrowserWithDefaults

`func NewUserAgentBrowserWithDefaults() *UserAgentBrowser`

NewUserAgentBrowserWithDefaults instantiates a new UserAgentBrowser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *UserAgentBrowser) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *UserAgentBrowser) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *UserAgentBrowser) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *UserAgentBrowser) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetVersion

`func (o *UserAgentBrowser) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserAgentBrowser) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserAgentBrowser) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserAgentBrowser) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



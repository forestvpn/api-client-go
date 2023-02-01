# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Latitude** | **float64** |  | 
**Longitude** | **float64** |  | 
**Country** | [**Country**](Country.md) |  | 
**Distance** | Pointer to **float64** | it&#39;s distance in kilometers between an user (longitude and latitude from request object) and a location object | [optional] 
**LatencyRate** | Pointer to **float64** | it&#39;s a rate from 0 to 1 which shows a connection quality. where 1 is good and 0 is bad | [optional] 
**AlternativeNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLocation

`func NewLocation(id string, name string, latitude float64, longitude float64, country Country, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.


### GetLatitude

`func (o *Location) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Location) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Location) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *Location) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Location) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Location) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetCountry

`func (o *Location) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Location) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Location) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetDistance

`func (o *Location) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *Location) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *Location) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *Location) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetLatencyRate

`func (o *Location) GetLatencyRate() float64`

GetLatencyRate returns the LatencyRate field if non-nil, zero value otherwise.

### GetLatencyRateOk

`func (o *Location) GetLatencyRateOk() (*float64, bool)`

GetLatencyRateOk returns a tuple with the LatencyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyRate

`func (o *Location) SetLatencyRate(v float64)`

SetLatencyRate sets LatencyRate field to given value.

### HasLatencyRate

`func (o *Location) HasLatencyRate() bool`

HasLatencyRate returns a boolean if a field has been set.

### GetAlternativeNames

`func (o *Location) GetAlternativeNames() []string`

GetAlternativeNames returns the AlternativeNames field if non-nil, zero value otherwise.

### GetAlternativeNamesOk

`func (o *Location) GetAlternativeNamesOk() (*[]string, bool)`

GetAlternativeNamesOk returns a tuple with the AlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeNames

`func (o *Location) SetAlternativeNames(v []string)`

SetAlternativeNames sets AlternativeNames field to given value.

### HasAlternativeNames

`func (o *Location) HasAlternativeNames() bool`

HasAlternativeNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



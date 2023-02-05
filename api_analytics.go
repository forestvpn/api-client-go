/*
ForestVPN API

ForestVPN - Fast, secure, and modern VPN. It offers Distributed Computing, Crypto Mining, P2P, Ad Blocking, TOR over VPN, 30+ locations, and a free version with unlimited data. 

API version: 2.0
Contact: support@forestvpn.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forestvpn_api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)


type AnalyticsApi interface {

	/*
	GetDataUsageStats Data Usage Stats

	Users data usage dtatistics for each device


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetDataUsageStatsRequest
	*/
	GetDataUsageStats(ctx context.Context) ApiGetDataUsageStatsRequest

	// GetDataUsageStatsExecute executes the request
	//  @return []AggregatedDataUsageStats
	GetDataUsageStatsExecute(r ApiGetDataUsageStatsRequest) ([]AggregatedDataUsageStats, *http.Response, error)
}

// AnalyticsApiService AnalyticsApi service
type AnalyticsApiService service

type ApiGetDataUsageStatsRequest struct {
	ctx context.Context
	ApiService AnalyticsApi
	aggrInterval *string
	dateAfter *string
	dateBefore *string
	deviceTypeIn *[]string
	deviceIn *[]string
	sortBy *string
}

// No more than 3 days for hour aggregation type
func (r ApiGetDataUsageStatsRequest) AggrInterval(aggrInterval string) ApiGetDataUsageStatsRequest {
	r.aggrInterval = &aggrInterval
	return r
}

// No more than 3 months between date_after and date_before
func (r ApiGetDataUsageStatsRequest) DateAfter(dateAfter string) ApiGetDataUsageStatsRequest {
	r.dateAfter = &dateAfter
	return r
}

// No more than 3 months between date_after and date_before
func (r ApiGetDataUsageStatsRequest) DateBefore(dateBefore string) ApiGetDataUsageStatsRequest {
	r.dateBefore = &dateBefore
	return r
}

// Filter by device type
func (r ApiGetDataUsageStatsRequest) DeviceTypeIn(deviceTypeIn []string) ApiGetDataUsageStatsRequest {
	r.deviceTypeIn = &deviceTypeIn
	return r
}

// Filter by device
func (r ApiGetDataUsageStatsRequest) DeviceIn(deviceIn []string) ApiGetDataUsageStatsRequest {
	r.deviceIn = &deviceIn
	return r
}

// Sort output by
func (r ApiGetDataUsageStatsRequest) SortBy(sortBy string) ApiGetDataUsageStatsRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiGetDataUsageStatsRequest) Execute() ([]AggregatedDataUsageStats, *http.Response, error) {
	return r.ApiService.GetDataUsageStatsExecute(r)
}

/*
GetDataUsageStats Data Usage Stats

Users data usage dtatistics for each device


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetDataUsageStatsRequest
*/
func (a *AnalyticsApiService) GetDataUsageStats(ctx context.Context) ApiGetDataUsageStatsRequest {
	return ApiGetDataUsageStatsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AggregatedDataUsageStats
func (a *AnalyticsApiService) GetDataUsageStatsExecute(r ApiGetDataUsageStatsRequest) ([]AggregatedDataUsageStats, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AggregatedDataUsageStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsApiService.GetDataUsageStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/analytics/data-usage/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aggrInterval == nil {
		return localVarReturnValue, nil, reportError("aggrInterval is required and must be specified")
	}

	parameterAddToQuery(localVarQueryParams, "aggr_interval", r.aggrInterval, "")
	if r.dateAfter != nil {
		parameterAddToQuery(localVarQueryParams, "date_after", r.dateAfter, "")
	}
	if r.dateBefore != nil {
		parameterAddToQuery(localVarQueryParams, "date_before", r.dateBefore, "")
	}
	if r.deviceTypeIn != nil {
		t := *r.deviceTypeIn
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "device_type__in", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "device_type__in", t, "multi")
		}
	}
	if r.deviceIn != nil {
		t := *r.deviceIn
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "device__in", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "device__in", t, "multi")
		}
	}
	if r.sortBy != nil {
		parameterAddToQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

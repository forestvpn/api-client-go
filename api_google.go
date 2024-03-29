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
)


type GoogleApi interface {

	/*
	VerifyPlayStorePurchase Play store purchase verification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiVerifyPlayStorePurchaseRequest
	*/
	VerifyPlayStorePurchase(ctx context.Context) ApiVerifyPlayStorePurchaseRequest

	// VerifyPlayStorePurchaseExecute executes the request
	VerifyPlayStorePurchaseExecute(r ApiVerifyPlayStorePurchaseRequest) (*http.Response, error)
}

// GoogleApiService GoogleApi service
type GoogleApiService service

type ApiVerifyPlayStorePurchaseRequest struct {
	ctx context.Context
	ApiService GoogleApi
	playStorePurchaseVerificationRequest *PlayStorePurchaseVerificationRequest
}

func (r ApiVerifyPlayStorePurchaseRequest) PlayStorePurchaseVerificationRequest(playStorePurchaseVerificationRequest PlayStorePurchaseVerificationRequest) ApiVerifyPlayStorePurchaseRequest {
	r.playStorePurchaseVerificationRequest = &playStorePurchaseVerificationRequest
	return r
}

func (r ApiVerifyPlayStorePurchaseRequest) Execute() (*http.Response, error) {
	return r.ApiService.VerifyPlayStorePurchaseExecute(r)
}

/*
VerifyPlayStorePurchase Play store purchase verification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifyPlayStorePurchaseRequest
*/
func (a *GoogleApiService) VerifyPlayStorePurchase(ctx context.Context) ApiVerifyPlayStorePurchaseRequest {
	return ApiVerifyPlayStorePurchaseRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *GoogleApiService) VerifyPlayStorePurchaseExecute(r ApiVerifyPlayStorePurchaseRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleApiService.VerifyPlayStorePurchase")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase/google/verify/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playStorePurchaseVerificationRequest == nil {
		return nil, reportError("playStorePurchaseVerificationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.playStorePurchaseVerificationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

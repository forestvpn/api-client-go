/*
ForestVPN API

ForestVPN defeats content restrictions and censorship to deliver unlimited access to video, music, social media, and more, from anywhere in the world. 

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


type AppleApi interface {

	/*
	VerifyAppStoreReceipt App store receipt verification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiVerifyAppStoreReceiptRequest
	*/
	VerifyAppStoreReceipt(ctx context.Context) ApiVerifyAppStoreReceiptRequest

	// VerifyAppStoreReceiptExecute executes the request
	VerifyAppStoreReceiptExecute(r ApiVerifyAppStoreReceiptRequest) (*http.Response, error)
}

// AppleApiService AppleApi service
type AppleApiService service

type ApiVerifyAppStoreReceiptRequest struct {
	ctx context.Context
	ApiService AppleApi
	appStoreReceiptVerificationRequest *AppStoreReceiptVerificationRequest
}

func (r ApiVerifyAppStoreReceiptRequest) AppStoreReceiptVerificationRequest(appStoreReceiptVerificationRequest AppStoreReceiptVerificationRequest) ApiVerifyAppStoreReceiptRequest {
	r.appStoreReceiptVerificationRequest = &appStoreReceiptVerificationRequest
	return r
}

func (r ApiVerifyAppStoreReceiptRequest) Execute() (*http.Response, error) {
	return r.ApiService.VerifyAppStoreReceiptExecute(r)
}

/*
VerifyAppStoreReceipt App store receipt verification

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifyAppStoreReceiptRequest
*/
func (a *AppleApiService) VerifyAppStoreReceipt(ctx context.Context) ApiVerifyAppStoreReceiptRequest {
	return ApiVerifyAppStoreReceiptRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AppleApiService) VerifyAppStoreReceiptExecute(r ApiVerifyAppStoreReceiptRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppleApiService.VerifyAppStoreReceipt")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/purchase/apple/verify/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appStoreReceiptVerificationRequest == nil {
		return nil, reportError("appStoreReceiptVerificationRequest is required and must be specified")
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
	localVarPostBody = r.appStoreReceiptVerificationRequest
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
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

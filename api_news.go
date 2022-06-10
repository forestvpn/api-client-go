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
	"strings"
)


type NewsApi interface {

	/*
	GetNotification Get notification content

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationID
	@return ApiGetNotificationRequest
	*/
	GetNotification(ctx context.Context, notificationID int32) ApiGetNotificationRequest

	// GetNotificationExecute executes the request
	//  @return NotificationDetail
	GetNotificationExecute(r ApiGetNotificationRequest) (*NotificationDetail, *http.Response, error)

	/*
	GetNotificationsUnreadCount Get unread notifications count

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetNotificationsUnreadCountRequest
	*/
	GetNotificationsUnreadCount(ctx context.Context) ApiGetNotificationsUnreadCountRequest

	// GetNotificationsUnreadCountExecute executes the request
	//  @return NotificationUnreadCount
	GetNotificationsUnreadCountExecute(r ApiGetNotificationsUnreadCountRequest) (*NotificationUnreadCount, *http.Response, error)

	/*
	ListNotifications Get notifications list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListNotificationsRequest
	*/
	ListNotifications(ctx context.Context) ApiListNotificationsRequest

	// ListNotificationsExecute executes the request
	//  @return []Notification
	ListNotificationsExecute(r ApiListNotificationsRequest) ([]Notification, *http.Response, error)

	/*
	UpdateNotificationMarkRead Mark notification as read by user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param notificationID
	@return ApiUpdateNotificationMarkReadRequest
	*/
	UpdateNotificationMarkRead(ctx context.Context, notificationID int32) ApiUpdateNotificationMarkReadRequest

	// UpdateNotificationMarkReadExecute executes the request
	UpdateNotificationMarkReadExecute(r ApiUpdateNotificationMarkReadRequest) (*http.Response, error)

	/*
	UpdateNotificationMarkReadAll Mark all notifications as read by user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateNotificationMarkReadAllRequest
	*/
	UpdateNotificationMarkReadAll(ctx context.Context) ApiUpdateNotificationMarkReadAllRequest

	// UpdateNotificationMarkReadAllExecute executes the request
	UpdateNotificationMarkReadAllExecute(r ApiUpdateNotificationMarkReadAllRequest) (*http.Response, error)
}

// NewsApiService NewsApi service
type NewsApiService service

type ApiGetNotificationRequest struct {
	ctx context.Context
	ApiService NewsApi
	notificationID int32
}

func (r ApiGetNotificationRequest) Execute() (*NotificationDetail, *http.Response, error) {
	return r.ApiService.GetNotificationExecute(r)
}

/*
GetNotification Get notification content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param notificationID
 @return ApiGetNotificationRequest
*/
func (a *NewsApiService) GetNotification(ctx context.Context, notificationID int32) ApiGetNotificationRequest {
	return ApiGetNotificationRequest{
		ApiService: a,
		ctx: ctx,
		notificationID: notificationID,
	}
}

// Execute executes the request
//  @return NotificationDetail
func (a *NewsApiService) GetNotificationExecute(r ApiGetNotificationRequest) (*NotificationDetail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotificationDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewsApiService.GetNotification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/news/notifications/{notificationID}/"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationID"+"}", url.PathEscape(parameterToString(r.notificationID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetNotificationsUnreadCountRequest struct {
	ctx context.Context
	ApiService NewsApi
}

func (r ApiGetNotificationsUnreadCountRequest) Execute() (*NotificationUnreadCount, *http.Response, error) {
	return r.ApiService.GetNotificationsUnreadCountExecute(r)
}

/*
GetNotificationsUnreadCount Get unread notifications count

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNotificationsUnreadCountRequest
*/
func (a *NewsApiService) GetNotificationsUnreadCount(ctx context.Context) ApiGetNotificationsUnreadCountRequest {
	return ApiGetNotificationsUnreadCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NotificationUnreadCount
func (a *NewsApiService) GetNotificationsUnreadCountExecute(r ApiGetNotificationsUnreadCountRequest) (*NotificationUnreadCount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotificationUnreadCount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewsApiService.GetNotificationsUnreadCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/news/unread_count/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListNotificationsRequest struct {
	ctx context.Context
	ApiService NewsApi
	isPublished *bool
}

func (r ApiListNotificationsRequest) IsPublished(isPublished bool) ApiListNotificationsRequest {
	r.isPublished = &isPublished
	return r
}

func (r ApiListNotificationsRequest) Execute() ([]Notification, *http.Response, error) {
	return r.ApiService.ListNotificationsExecute(r)
}

/*
ListNotifications Get notifications list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNotificationsRequest
*/
func (a *NewsApiService) ListNotifications(ctx context.Context) ApiListNotificationsRequest {
	return ApiListNotificationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Notification
func (a *NewsApiService) ListNotificationsExecute(r ApiListNotificationsRequest) ([]Notification, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Notification
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewsApiService.ListNotifications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/news/notifications/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isPublished != nil {
		localVarQueryParams.Add("is_published", parameterToString(*r.isPublished, ""))
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

type ApiUpdateNotificationMarkReadRequest struct {
	ctx context.Context
	ApiService NewsApi
	notificationID int32
}

func (r ApiUpdateNotificationMarkReadRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateNotificationMarkReadExecute(r)
}

/*
UpdateNotificationMarkRead Mark notification as read by user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param notificationID
 @return ApiUpdateNotificationMarkReadRequest
*/
func (a *NewsApiService) UpdateNotificationMarkRead(ctx context.Context, notificationID int32) ApiUpdateNotificationMarkReadRequest {
	return ApiUpdateNotificationMarkReadRequest{
		ApiService: a,
		ctx: ctx,
		notificationID: notificationID,
	}
}

// Execute executes the request
func (a *NewsApiService) UpdateNotificationMarkReadExecute(r ApiUpdateNotificationMarkReadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewsApiService.UpdateNotificationMarkRead")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/news/notifications/{notificationID}/mark_read/"
	localVarPath = strings.Replace(localVarPath, "{"+"notificationID"+"}", url.PathEscape(parameterToString(r.notificationID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiUpdateNotificationMarkReadAllRequest struct {
	ctx context.Context
	ApiService NewsApi
}

func (r ApiUpdateNotificationMarkReadAllRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateNotificationMarkReadAllExecute(r)
}

/*
UpdateNotificationMarkReadAll Mark all notifications as read by user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateNotificationMarkReadAllRequest
*/
func (a *NewsApiService) UpdateNotificationMarkReadAll(ctx context.Context) ApiUpdateNotificationMarkReadAllRequest {
	return ApiUpdateNotificationMarkReadAllRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NewsApiService) UpdateNotificationMarkReadAllExecute(r ApiUpdateNotificationMarkReadAllRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NewsApiService.UpdateNotificationMarkReadAll")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/news/notifications/mark_read_all/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

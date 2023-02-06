// Code generated by MockGen. DO NOT EDIT.
// Source: ../api-client-go/api_notifications.go

// Package forestvpn_api_test is a generated GoMock package.
package forestvpn_api_test

import (
	context "context"
	http "net/http"
	reflect "reflect"

	api_client_go "github.com/forestvpn/api-client-go"
	gomock "github.com/golang/mock/gomock"
)

// MockNotificationsApi is a mock of NotificationsApi interface.
type MockNotificationsApi struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationsApiMockRecorder
}

// MockNotificationsApiMockRecorder is the mock recorder for MockNotificationsApi.
type MockNotificationsApiMockRecorder struct {
	mock *MockNotificationsApi
}

// NewMockNotificationsApi creates a new mock instance.
func NewMockNotificationsApi(ctrl *gomock.Controller) *MockNotificationsApi {
	mock := &MockNotificationsApi{ctrl: ctrl}
	mock.recorder = &MockNotificationsApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationsApi) EXPECT() *MockNotificationsApiMockRecorder {
	return m.recorder
}

// GetNotificationsUnreadCount mocks base method.
func (m *MockNotificationsApi) GetNotificationsUnreadCount(ctx context.Context) api_client_go.ApiGetNotificationsUnreadCountRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotificationsUnreadCount", ctx)
	ret0, _ := ret[0].(api_client_go.ApiGetNotificationsUnreadCountRequest)
	return ret0
}

// GetNotificationsUnreadCount indicates an expected call of GetNotificationsUnreadCount.
func (mr *MockNotificationsApiMockRecorder) GetNotificationsUnreadCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotificationsUnreadCount", reflect.TypeOf((*MockNotificationsApi)(nil).GetNotificationsUnreadCount), ctx)
}

// GetNotificationsUnreadCountExecute mocks base method.
func (m *MockNotificationsApi) GetNotificationsUnreadCountExecute(r api_client_go.ApiGetNotificationsUnreadCountRequest) (*api_client_go.NotificationUnreadCount, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotificationsUnreadCountExecute", r)
	ret0, _ := ret[0].(*api_client_go.NotificationUnreadCount)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNotificationsUnreadCountExecute indicates an expected call of GetNotificationsUnreadCountExecute.
func (mr *MockNotificationsApiMockRecorder) GetNotificationsUnreadCountExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotificationsUnreadCountExecute", reflect.TypeOf((*MockNotificationsApi)(nil).GetNotificationsUnreadCountExecute), r)
}

// ListNotifications mocks base method.
func (m *MockNotificationsApi) ListNotifications(ctx context.Context) api_client_go.ApiListNotificationsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotifications", ctx)
	ret0, _ := ret[0].(api_client_go.ApiListNotificationsRequest)
	return ret0
}

// ListNotifications indicates an expected call of ListNotifications.
func (mr *MockNotificationsApiMockRecorder) ListNotifications(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotifications", reflect.TypeOf((*MockNotificationsApi)(nil).ListNotifications), ctx)
}

// ListNotificationsExecute mocks base method.
func (m *MockNotificationsApi) ListNotificationsExecute(r api_client_go.ApiListNotificationsRequest) ([]api_client_go.NotificationAllList, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotificationsExecute", r)
	ret0, _ := ret[0].([]api_client_go.NotificationAllList)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListNotificationsExecute indicates an expected call of ListNotificationsExecute.
func (mr *MockNotificationsApiMockRecorder) ListNotificationsExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotificationsExecute", reflect.TypeOf((*MockNotificationsApi)(nil).ListNotificationsExecute), r)
}

// UpdateNotificationMarkRead mocks base method.
func (m *MockNotificationsApi) UpdateNotificationMarkRead(ctx context.Context, slug int32) api_client_go.ApiUpdateNotificationMarkReadRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotificationMarkRead", ctx, slug)
	ret0, _ := ret[0].(api_client_go.ApiUpdateNotificationMarkReadRequest)
	return ret0
}

// UpdateNotificationMarkRead indicates an expected call of UpdateNotificationMarkRead.
func (mr *MockNotificationsApiMockRecorder) UpdateNotificationMarkRead(ctx, slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotificationMarkRead", reflect.TypeOf((*MockNotificationsApi)(nil).UpdateNotificationMarkRead), ctx, slug)
}

// UpdateNotificationMarkReadAll mocks base method.
func (m *MockNotificationsApi) UpdateNotificationMarkReadAll(ctx context.Context) api_client_go.ApiUpdateNotificationMarkReadAllRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotificationMarkReadAll", ctx)
	ret0, _ := ret[0].(api_client_go.ApiUpdateNotificationMarkReadAllRequest)
	return ret0
}

// UpdateNotificationMarkReadAll indicates an expected call of UpdateNotificationMarkReadAll.
func (mr *MockNotificationsApiMockRecorder) UpdateNotificationMarkReadAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotificationMarkReadAll", reflect.TypeOf((*MockNotificationsApi)(nil).UpdateNotificationMarkReadAll), ctx)
}

// UpdateNotificationMarkReadAllExecute mocks base method.
func (m *MockNotificationsApi) UpdateNotificationMarkReadAllExecute(r api_client_go.ApiUpdateNotificationMarkReadAllRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotificationMarkReadAllExecute", r)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNotificationMarkReadAllExecute indicates an expected call of UpdateNotificationMarkReadAllExecute.
func (mr *MockNotificationsApiMockRecorder) UpdateNotificationMarkReadAllExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotificationMarkReadAllExecute", reflect.TypeOf((*MockNotificationsApi)(nil).UpdateNotificationMarkReadAllExecute), r)
}

// UpdateNotificationMarkReadExecute mocks base method.
func (m *MockNotificationsApi) UpdateNotificationMarkReadExecute(r api_client_go.ApiUpdateNotificationMarkReadRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotificationMarkReadExecute", r)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNotificationMarkReadExecute indicates an expected call of UpdateNotificationMarkReadExecute.
func (mr *MockNotificationsApiMockRecorder) UpdateNotificationMarkReadExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotificationMarkReadExecute", reflect.TypeOf((*MockNotificationsApi)(nil).UpdateNotificationMarkReadExecute), r)
}
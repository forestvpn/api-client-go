// Code generated by MockGen. DO NOT EDIT.
// Source: ../api-client-go/api_app.go

// Package forestvpn_api_mock is a generated GoMock package.
package forestvpn_api_mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	api_client_go "github.com/forestvpn/api-client-go"
	gomock "github.com/golang/mock/gomock"
)

// MockAppApi is a mock of AppApi interface.
type MockAppApi struct {
	ctrl     *gomock.Controller
	recorder *MockAppApiMockRecorder
}

// MockAppApiMockRecorder is the mock recorder for MockAppApi.
type MockAppApiMockRecorder struct {
	mock *MockAppApi
}

// NewMockAppApi creates a new mock instance.
func NewMockAppApi(ctrl *gomock.Controller) *MockAppApi {
	mock := &MockAppApi{ctrl: ctrl}
	mock.recorder = &MockAppApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppApi) EXPECT() *MockAppApiMockRecorder {
	return m.recorder
}

// GetCurrentUserDevice mocks base method.
func (m *MockAppApi) GetCurrentUserDevice(ctx context.Context) api_client_go.ApiGetCurrentUserDeviceRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUserDevice", ctx)
	ret0, _ := ret[0].(api_client_go.ApiGetCurrentUserDeviceRequest)
	return ret0
}

// GetCurrentUserDevice indicates an expected call of GetCurrentUserDevice.
func (mr *MockAppApiMockRecorder) GetCurrentUserDevice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUserDevice", reflect.TypeOf((*MockAppApi)(nil).GetCurrentUserDevice), ctx)
}

// GetCurrentUserDeviceExecute mocks base method.
func (m *MockAppApi) GetCurrentUserDeviceExecute(r api_client_go.ApiGetCurrentUserDeviceRequest) (*api_client_go.UserDevice, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUserDeviceExecute", r)
	ret0, _ := ret[0].(*api_client_go.UserDevice)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCurrentUserDeviceExecute indicates an expected call of GetCurrentUserDeviceExecute.
func (mr *MockAppApiMockRecorder) GetCurrentUserDeviceExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUserDeviceExecute", reflect.TypeOf((*MockAppApi)(nil).GetCurrentUserDeviceExecute), r)
}

// UpdateCurrentUserDevice mocks base method.
func (m *MockAppApi) UpdateCurrentUserDevice(ctx context.Context) api_client_go.ApiUpdateCurrentUserDeviceRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCurrentUserDevice", ctx)
	ret0, _ := ret[0].(api_client_go.ApiUpdateCurrentUserDeviceRequest)
	return ret0
}

// UpdateCurrentUserDevice indicates an expected call of UpdateCurrentUserDevice.
func (mr *MockAppApiMockRecorder) UpdateCurrentUserDevice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCurrentUserDevice", reflect.TypeOf((*MockAppApi)(nil).UpdateCurrentUserDevice), ctx)
}

// UpdateCurrentUserDeviceExecute mocks base method.
func (m *MockAppApi) UpdateCurrentUserDeviceExecute(r api_client_go.ApiUpdateCurrentUserDeviceRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCurrentUserDeviceExecute", r)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCurrentUserDeviceExecute indicates an expected call of UpdateCurrentUserDeviceExecute.
func (mr *MockAppApiMockRecorder) UpdateCurrentUserDeviceExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCurrentUserDeviceExecute", reflect.TypeOf((*MockAppApi)(nil).UpdateCurrentUserDeviceExecute), r)
}
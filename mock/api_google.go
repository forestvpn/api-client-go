// Code generated by MockGen. DO NOT EDIT.
// Source: ../api-client-go/api_google.go

// Package forestvpn_api_mock is a generated GoMock package.
package forestvpn_api_mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	api_client_go "github.com/forestvpn/api-client-go"
	gomock "github.com/golang/mock/gomock"
)

// MockGoogleApi is a mock of GoogleApi interface.
type MockGoogleApi struct {
	ctrl     *gomock.Controller
	recorder *MockGoogleApiMockRecorder
}

// MockGoogleApiMockRecorder is the mock recorder for MockGoogleApi.
type MockGoogleApiMockRecorder struct {
	mock *MockGoogleApi
}

// NewMockGoogleApi creates a new mock instance.
func NewMockGoogleApi(ctrl *gomock.Controller) *MockGoogleApi {
	mock := &MockGoogleApi{ctrl: ctrl}
	mock.recorder = &MockGoogleApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGoogleApi) EXPECT() *MockGoogleApiMockRecorder {
	return m.recorder
}

// VerifyPlayStorePurchase mocks base method.
func (m *MockGoogleApi) VerifyPlayStorePurchase(ctx context.Context) api_client_go.ApiVerifyPlayStorePurchaseRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPlayStorePurchase", ctx)
	ret0, _ := ret[0].(api_client_go.ApiVerifyPlayStorePurchaseRequest)
	return ret0
}

// VerifyPlayStorePurchase indicates an expected call of VerifyPlayStorePurchase.
func (mr *MockGoogleApiMockRecorder) VerifyPlayStorePurchase(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPlayStorePurchase", reflect.TypeOf((*MockGoogleApi)(nil).VerifyPlayStorePurchase), ctx)
}

// VerifyPlayStorePurchaseExecute mocks base method.
func (m *MockGoogleApi) VerifyPlayStorePurchaseExecute(r api_client_go.ApiVerifyPlayStorePurchaseRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPlayStorePurchaseExecute", r)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyPlayStorePurchaseExecute indicates an expected call of VerifyPlayStorePurchaseExecute.
func (mr *MockGoogleApiMockRecorder) VerifyPlayStorePurchaseExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPlayStorePurchaseExecute", reflect.TypeOf((*MockGoogleApi)(nil).VerifyPlayStorePurchaseExecute), r)
}

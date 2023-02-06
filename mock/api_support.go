// Code generated by MockGen. DO NOT EDIT.
// Source: ../api-client-go/api_support.go

// Package forestvpn_api_mock is a generated GoMock package.
package forestvpn_api_mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	api_client_go "github.com/forestvpn/api-client-go"
	gomock "github.com/golang/mock/gomock"
)

// MockSupportApi is a mock of SupportApi interface.
type MockSupportApi struct {
	ctrl     *gomock.Controller
	recorder *MockSupportApiMockRecorder
}

// MockSupportApiMockRecorder is the mock recorder for MockSupportApi.
type MockSupportApiMockRecorder struct {
	mock *MockSupportApi
}

// NewMockSupportApi creates a new mock instance.
func NewMockSupportApi(ctrl *gomock.Controller) *MockSupportApi {
	mock := &MockSupportApi{ctrl: ctrl}
	mock.recorder = &MockSupportApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSupportApi) EXPECT() *MockSupportApiMockRecorder {
	return m.recorder
}

// CreateSupportTicket mocks base method.
func (m *MockSupportApi) CreateSupportTicket(ctx context.Context) api_client_go.ApiCreateSupportTicketRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSupportTicket", ctx)
	ret0, _ := ret[0].(api_client_go.ApiCreateSupportTicketRequest)
	return ret0
}

// CreateSupportTicket indicates an expected call of CreateSupportTicket.
func (mr *MockSupportApiMockRecorder) CreateSupportTicket(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSupportTicket", reflect.TypeOf((*MockSupportApi)(nil).CreateSupportTicket), ctx)
}

// CreateSupportTicketExecute mocks base method.
func (m *MockSupportApi) CreateSupportTicketExecute(r api_client_go.ApiCreateSupportTicketRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSupportTicketExecute", r)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSupportTicketExecute indicates an expected call of CreateSupportTicketExecute.
func (mr *MockSupportApiMockRecorder) CreateSupportTicketExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSupportTicketExecute", reflect.TypeOf((*MockSupportApi)(nil).CreateSupportTicketExecute), r)
}

// GetSupportTicketCategory mocks base method.
func (m *MockSupportApi) GetSupportTicketCategory(ctx context.Context) api_client_go.ApiGetSupportTicketCategoryRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportTicketCategory", ctx)
	ret0, _ := ret[0].(api_client_go.ApiGetSupportTicketCategoryRequest)
	return ret0
}

// GetSupportTicketCategory indicates an expected call of GetSupportTicketCategory.
func (mr *MockSupportApiMockRecorder) GetSupportTicketCategory(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportTicketCategory", reflect.TypeOf((*MockSupportApi)(nil).GetSupportTicketCategory), ctx)
}

// GetSupportTicketCategoryExecute mocks base method.
func (m *MockSupportApi) GetSupportTicketCategoryExecute(r api_client_go.ApiGetSupportTicketCategoryRequest) ([]api_client_go.TicketCategory, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportTicketCategoryExecute", r)
	ret0, _ := ret[0].([]api_client_go.TicketCategory)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSupportTicketCategoryExecute indicates an expected call of GetSupportTicketCategoryExecute.
func (mr *MockSupportApiMockRecorder) GetSupportTicketCategoryExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportTicketCategoryExecute", reflect.TypeOf((*MockSupportApi)(nil).GetSupportTicketCategoryExecute), r)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ../api-client-go/api_checkout.go

// Package forestvpn_api_mock is a generated GoMock package.
package forestvpn_api_mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	api_client_go "github.com/forestvpn/api-client-go"
	gomock "github.com/golang/mock/gomock"
)

// MockCheckoutApi is a mock of CheckoutApi interface.
type MockCheckoutApi struct {
	ctrl     *gomock.Controller
	recorder *MockCheckoutApiMockRecorder
}

// MockCheckoutApiMockRecorder is the mock recorder for MockCheckoutApi.
type MockCheckoutApiMockRecorder struct {
	mock *MockCheckoutApi
}

// NewMockCheckoutApi creates a new mock instance.
func NewMockCheckoutApi(ctrl *gomock.Controller) *MockCheckoutApi {
	mock := &MockCheckoutApi{ctrl: ctrl}
	mock.recorder = &MockCheckoutApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckoutApi) EXPECT() *MockCheckoutApiMockRecorder {
	return m.recorder
}

// ApplyCouponCheckoutSession mocks base method.
func (m *MockCheckoutApi) ApplyCouponCheckoutSession(ctx context.Context, sessionID string) api_client_go.ApiApplyCouponCheckoutSessionRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyCouponCheckoutSession", ctx, sessionID)
	ret0, _ := ret[0].(api_client_go.ApiApplyCouponCheckoutSessionRequest)
	return ret0
}

// ApplyCouponCheckoutSession indicates an expected call of ApplyCouponCheckoutSession.
func (mr *MockCheckoutApiMockRecorder) ApplyCouponCheckoutSession(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyCouponCheckoutSession", reflect.TypeOf((*MockCheckoutApi)(nil).ApplyCouponCheckoutSession), ctx, sessionID)
}

// ApplyCouponCheckoutSessionExecute mocks base method.
func (m *MockCheckoutApi) ApplyCouponCheckoutSessionExecute(r api_client_go.ApiApplyCouponCheckoutSessionRequest) (*api_client_go.CouponCheckoutSession, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyCouponCheckoutSessionExecute", r)
	ret0, _ := ret[0].(*api_client_go.CouponCheckoutSession)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ApplyCouponCheckoutSessionExecute indicates an expected call of ApplyCouponCheckoutSessionExecute.
func (mr *MockCheckoutApiMockRecorder) ApplyCouponCheckoutSessionExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyCouponCheckoutSessionExecute", reflect.TypeOf((*MockCheckoutApi)(nil).ApplyCouponCheckoutSessionExecute), r)
}

// CreateCheckoutSession mocks base method.
func (m *MockCheckoutApi) CreateCheckoutSession(ctx context.Context) api_client_go.ApiCreateCheckoutSessionRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCheckoutSession", ctx)
	ret0, _ := ret[0].(api_client_go.ApiCreateCheckoutSessionRequest)
	return ret0
}

// CreateCheckoutSession indicates an expected call of CreateCheckoutSession.
func (mr *MockCheckoutApiMockRecorder) CreateCheckoutSession(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCheckoutSession", reflect.TypeOf((*MockCheckoutApi)(nil).CreateCheckoutSession), ctx)
}

// CreateCheckoutSessionExecute mocks base method.
func (m *MockCheckoutApi) CreateCheckoutSessionExecute(r api_client_go.ApiCreateCheckoutSessionRequest) (*api_client_go.CheckoutSession, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCheckoutSessionExecute", r)
	ret0, _ := ret[0].(*api_client_go.CheckoutSession)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCheckoutSessionExecute indicates an expected call of CreateCheckoutSessionExecute.
func (mr *MockCheckoutApiMockRecorder) CreateCheckoutSessionExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCheckoutSessionExecute", reflect.TypeOf((*MockCheckoutApi)(nil).CreateCheckoutSessionExecute), r)
}

// CreateWaitListRequest mocks base method.
func (m *MockCheckoutApi) CreateWaitListRequest(ctx context.Context) api_client_go.ApiCreateWaitListRequestRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWaitListRequest", ctx)
	ret0, _ := ret[0].(api_client_go.ApiCreateWaitListRequestRequest)
	return ret0
}

// CreateWaitListRequest indicates an expected call of CreateWaitListRequest.
func (mr *MockCheckoutApiMockRecorder) CreateWaitListRequest(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWaitListRequest", reflect.TypeOf((*MockCheckoutApi)(nil).CreateWaitListRequest), ctx)
}

// CreateWaitListRequestExecute mocks base method.
func (m *MockCheckoutApi) CreateWaitListRequestExecute(r api_client_go.ApiCreateWaitListRequestRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWaitListRequestExecute", r)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWaitListRequestExecute indicates an expected call of CreateWaitListRequestExecute.
func (mr *MockCheckoutApiMockRecorder) CreateWaitListRequestExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWaitListRequestExecute", reflect.TypeOf((*MockCheckoutApi)(nil).CreateWaitListRequestExecute), r)
}

// ExpireCheckoutSession mocks base method.
func (m *MockCheckoutApi) ExpireCheckoutSession(ctx context.Context, sessionID string) api_client_go.ApiExpireCheckoutSessionRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireCheckoutSession", ctx, sessionID)
	ret0, _ := ret[0].(api_client_go.ApiExpireCheckoutSessionRequest)
	return ret0
}

// ExpireCheckoutSession indicates an expected call of ExpireCheckoutSession.
func (mr *MockCheckoutApiMockRecorder) ExpireCheckoutSession(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireCheckoutSession", reflect.TypeOf((*MockCheckoutApi)(nil).ExpireCheckoutSession), ctx, sessionID)
}

// ExpireCheckoutSessionExecute mocks base method.
func (m *MockCheckoutApi) ExpireCheckoutSessionExecute(r api_client_go.ApiExpireCheckoutSessionRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireCheckoutSessionExecute", r)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpireCheckoutSessionExecute indicates an expected call of ExpireCheckoutSessionExecute.
func (mr *MockCheckoutApiMockRecorder) ExpireCheckoutSessionExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireCheckoutSessionExecute", reflect.TypeOf((*MockCheckoutApi)(nil).ExpireCheckoutSessionExecute), r)
}

// GetCheckoutSession mocks base method.
func (m *MockCheckoutApi) GetCheckoutSession(ctx context.Context, sessionID string) api_client_go.ApiGetCheckoutSessionRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckoutSession", ctx, sessionID)
	ret0, _ := ret[0].(api_client_go.ApiGetCheckoutSessionRequest)
	return ret0
}

// GetCheckoutSession indicates an expected call of GetCheckoutSession.
func (mr *MockCheckoutApiMockRecorder) GetCheckoutSession(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckoutSession", reflect.TypeOf((*MockCheckoutApi)(nil).GetCheckoutSession), ctx, sessionID)
}

// GetCheckoutSessionExecute mocks base method.
func (m *MockCheckoutApi) GetCheckoutSessionExecute(r api_client_go.ApiGetCheckoutSessionRequest) (*api_client_go.CheckoutSession, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckoutSessionExecute", r)
	ret0, _ := ret[0].(*api_client_go.CheckoutSession)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCheckoutSessionExecute indicates an expected call of GetCheckoutSessionExecute.
func (mr *MockCheckoutApiMockRecorder) GetCheckoutSessionExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckoutSessionExecute", reflect.TypeOf((*MockCheckoutApi)(nil).GetCheckoutSessionExecute), r)
}

// GetStripeCheckoutSession mocks base method.
func (m *MockCheckoutApi) GetStripeCheckoutSession(ctx context.Context, sessionID string) api_client_go.ApiGetStripeCheckoutSessionRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStripeCheckoutSession", ctx, sessionID)
	ret0, _ := ret[0].(api_client_go.ApiGetStripeCheckoutSessionRequest)
	return ret0
}

// GetStripeCheckoutSession indicates an expected call of GetStripeCheckoutSession.
func (mr *MockCheckoutApiMockRecorder) GetStripeCheckoutSession(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStripeCheckoutSession", reflect.TypeOf((*MockCheckoutApi)(nil).GetStripeCheckoutSession), ctx, sessionID)
}

// GetStripeCheckoutSessionExecute mocks base method.
func (m *MockCheckoutApi) GetStripeCheckoutSessionExecute(r api_client_go.ApiGetStripeCheckoutSessionRequest) (*api_client_go.StripeCheckoutSession, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStripeCheckoutSessionExecute", r)
	ret0, _ := ret[0].(*api_client_go.StripeCheckoutSession)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStripeCheckoutSessionExecute indicates an expected call of GetStripeCheckoutSessionExecute.
func (mr *MockCheckoutApiMockRecorder) GetStripeCheckoutSessionExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStripeCheckoutSessionExecute", reflect.TypeOf((*MockCheckoutApi)(nil).GetStripeCheckoutSessionExecute), r)
}

// GetStripePaymentIntent mocks base method.
func (m *MockCheckoutApi) GetStripePaymentIntent(ctx context.Context, sessionID string) api_client_go.ApiGetStripePaymentIntentRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStripePaymentIntent", ctx, sessionID)
	ret0, _ := ret[0].(api_client_go.ApiGetStripePaymentIntentRequest)
	return ret0
}

// GetStripePaymentIntent indicates an expected call of GetStripePaymentIntent.
func (mr *MockCheckoutApiMockRecorder) GetStripePaymentIntent(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStripePaymentIntent", reflect.TypeOf((*MockCheckoutApi)(nil).GetStripePaymentIntent), ctx, sessionID)
}

// GetStripePaymentIntentExecute mocks base method.
func (m *MockCheckoutApi) GetStripePaymentIntentExecute(r api_client_go.ApiGetStripePaymentIntentRequest) (*api_client_go.StripePaymentIntent, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStripePaymentIntentExecute", r)
	ret0, _ := ret[0].(*api_client_go.StripePaymentIntent)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStripePaymentIntentExecute indicates an expected call of GetStripePaymentIntentExecute.
func (mr *MockCheckoutApiMockRecorder) GetStripePaymentIntentExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStripePaymentIntentExecute", reflect.TypeOf((*MockCheckoutApi)(nil).GetStripePaymentIntentExecute), r)
}

// ProcessCloudPaymentsAuth mocks base method.
func (m *MockCheckoutApi) ProcessCloudPaymentsAuth(ctx context.Context, sessionID string) api_client_go.ApiProcessCloudPaymentsAuthRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessCloudPaymentsAuth", ctx, sessionID)
	ret0, _ := ret[0].(api_client_go.ApiProcessCloudPaymentsAuthRequest)
	return ret0
}

// ProcessCloudPaymentsAuth indicates an expected call of ProcessCloudPaymentsAuth.
func (mr *MockCheckoutApiMockRecorder) ProcessCloudPaymentsAuth(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessCloudPaymentsAuth", reflect.TypeOf((*MockCheckoutApi)(nil).ProcessCloudPaymentsAuth), ctx, sessionID)
}

// ProcessCloudPaymentsAuthExecute mocks base method.
func (m *MockCheckoutApi) ProcessCloudPaymentsAuthExecute(r api_client_go.ApiProcessCloudPaymentsAuthRequest) (*api_client_go.CloudPaymentsAuth, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessCloudPaymentsAuthExecute", r)
	ret0, _ := ret[0].(*api_client_go.CloudPaymentsAuth)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProcessCloudPaymentsAuthExecute indicates an expected call of ProcessCloudPaymentsAuthExecute.
func (mr *MockCheckoutApiMockRecorder) ProcessCloudPaymentsAuthExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessCloudPaymentsAuthExecute", reflect.TypeOf((*MockCheckoutApi)(nil).ProcessCloudPaymentsAuthExecute), r)
}

// ProcessCloudPaymentsPost3ds mocks base method.
func (m *MockCheckoutApi) ProcessCloudPaymentsPost3ds(ctx context.Context, sessionID string) api_client_go.ApiProcessCloudPaymentsPost3dsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessCloudPaymentsPost3ds", ctx, sessionID)
	ret0, _ := ret[0].(api_client_go.ApiProcessCloudPaymentsPost3dsRequest)
	return ret0
}

// ProcessCloudPaymentsPost3ds indicates an expected call of ProcessCloudPaymentsPost3ds.
func (mr *MockCheckoutApiMockRecorder) ProcessCloudPaymentsPost3ds(ctx, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessCloudPaymentsPost3ds", reflect.TypeOf((*MockCheckoutApi)(nil).ProcessCloudPaymentsPost3ds), ctx, sessionID)
}

// ProcessCloudPaymentsPost3dsExecute mocks base method.
func (m *MockCheckoutApi) ProcessCloudPaymentsPost3dsExecute(r api_client_go.ApiProcessCloudPaymentsPost3dsRequest) (*api_client_go.CloudPaymentsPost3ds, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessCloudPaymentsPost3dsExecute", r)
	ret0, _ := ret[0].(*api_client_go.CloudPaymentsPost3ds)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProcessCloudPaymentsPost3dsExecute indicates an expected call of ProcessCloudPaymentsPost3dsExecute.
func (mr *MockCheckoutApiMockRecorder) ProcessCloudPaymentsPost3dsExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessCloudPaymentsPost3dsExecute", reflect.TypeOf((*MockCheckoutApi)(nil).ProcessCloudPaymentsPost3dsExecute), r)
}

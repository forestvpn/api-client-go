// Code generated by MockGen. DO NOT EDIT.
// Source: ../api-client-go/api_analytics.go

// Package forestvpn_api_test is a generated GoMock package.
package forestvpn_api_test

import (
	context "context"
	http "net/http"
	reflect "reflect"

	api_client_go "github.com/forestvpn/api-client-go"
	gomock "github.com/golang/mock/gomock"
)

// MockAnalyticsApi is a mock of AnalyticsApi interface.
type MockAnalyticsApi struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyticsApiMockRecorder
}

// MockAnalyticsApiMockRecorder is the mock recorder for MockAnalyticsApi.
type MockAnalyticsApiMockRecorder struct {
	mock *MockAnalyticsApi
}

// NewMockAnalyticsApi creates a new mock instance.
func NewMockAnalyticsApi(ctrl *gomock.Controller) *MockAnalyticsApi {
	mock := &MockAnalyticsApi{ctrl: ctrl}
	mock.recorder = &MockAnalyticsApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyticsApi) EXPECT() *MockAnalyticsApiMockRecorder {
	return m.recorder
}

// GetDataUsageStats mocks base method.
func (m *MockAnalyticsApi) GetDataUsageStats(ctx context.Context) api_client_go.ApiGetDataUsageStatsRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataUsageStats", ctx)
	ret0, _ := ret[0].(api_client_go.ApiGetDataUsageStatsRequest)
	return ret0
}

// GetDataUsageStats indicates an expected call of GetDataUsageStats.
func (mr *MockAnalyticsApiMockRecorder) GetDataUsageStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUsageStats", reflect.TypeOf((*MockAnalyticsApi)(nil).GetDataUsageStats), ctx)
}

// GetDataUsageStatsExecute mocks base method.
func (m *MockAnalyticsApi) GetDataUsageStatsExecute(r api_client_go.ApiGetDataUsageStatsRequest) ([]api_client_go.AggregatedDataUsageStats, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataUsageStatsExecute", r)
	ret0, _ := ret[0].([]api_client_go.AggregatedDataUsageStats)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDataUsageStatsExecute indicates an expected call of GetDataUsageStatsExecute.
func (mr *MockAnalyticsApiMockRecorder) GetDataUsageStatsExecute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUsageStatsExecute", reflect.TypeOf((*MockAnalyticsApi)(nil).GetDataUsageStatsExecute), r)
}
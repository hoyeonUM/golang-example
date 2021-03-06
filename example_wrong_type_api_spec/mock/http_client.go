// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hoyeonUM/golang-example/example_wrong_type_api_spec (interfaces: HttpClient)

// Package mock_example_wrong_type_api_spec is a generated GoMock package.
package mock_example_wrong_type_api_spec

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHttpClient is a mock of HttpClient interface
type MockHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientMockRecorder
}

// MockHttpClientMockRecorder is the mock recorder for MockHttpClient
type MockHttpClientMockRecorder struct {
	mock *MockHttpClient
}

// NewMockHttpClient creates a new mock instance
func NewMockHttpClient(ctrl *gomock.Controller) *MockHttpClient {
	mock := &MockHttpClient{ctrl: ctrl}
	mock.recorder = &MockHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpClient) EXPECT() *MockHttpClientMockRecorder {
	return m.recorder
}

// Request mocks base method
func (m *MockHttpClient) Request() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(string)
	return ret0
}

// Request indicates an expected call of Request
func (mr *MockHttpClientMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockHttpClient)(nil).Request))
}

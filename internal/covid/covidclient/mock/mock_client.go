// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	covidclient "covid-19-project/internal/covid/covidclient"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCovidClient is a mock of CovidClient interface.
type MockCovidClient struct {
	ctrl     *gomock.Controller
	recorder *MockCovidClientMockRecorder
}

// MockCovidClientMockRecorder is the mock recorder for MockCovidClient.
type MockCovidClientMockRecorder struct {
	mock *MockCovidClient
}

// NewMockCovidClient creates a new mock instance.
func NewMockCovidClient(ctrl *gomock.Controller) *MockCovidClient {
	mock := &MockCovidClient{ctrl: ctrl}
	mock.recorder = &MockCovidClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCovidClient) EXPECT() *MockCovidClientMockRecorder {
	return m.recorder
}

// GetCovidCases mocks base method.
func (m *MockCovidClient) GetCovidCases() (*covidclient.Covid19, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCovidCases")
	ret0, _ := ret[0].(*covidclient.Covid19)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCovidCases indicates an expected call of GetCovidCases.
func (mr *MockCovidClientMockRecorder) GetCovidCases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCovidCases", reflect.TypeOf((*MockCovidClient)(nil).GetCovidCases))
}

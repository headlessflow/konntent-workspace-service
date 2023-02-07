// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/nrclient/nrclient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

// MockNewRelicInstance is a mock of NewRelicInstance interface.
type MockNewRelicInstance struct {
	ctrl     *gomock.Controller
	recorder *MockNewRelicInstanceMockRecorder
}

// MockNewRelicInstanceMockRecorder is the mock recorder for MockNewRelicInstance.
type MockNewRelicInstanceMockRecorder struct {
	mock *MockNewRelicInstance
}

// NewMockNewRelicInstance creates a new mock instance.
func NewMockNewRelicInstance(ctrl *gomock.Controller) *MockNewRelicInstance {
	mock := &MockNewRelicInstance{ctrl: ctrl}
	mock.recorder = &MockNewRelicInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNewRelicInstance) EXPECT() *MockNewRelicInstanceMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockNewRelicInstance) Application() *newrelic.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(*newrelic.Application)
	return ret0
}

// Application indicates an expected call of Application.
func (mr *MockNewRelicInstanceMockRecorder) Application() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockNewRelicInstance)(nil).Application))
}

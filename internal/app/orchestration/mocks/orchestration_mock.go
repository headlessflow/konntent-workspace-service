// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/orchestration/workspace.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	request "konntent-workspace-service/internal/app/dto/request"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWorkspaceOrchestrator is a mock of WorkspaceOrchestrator interface.
type MockWorkspaceOrchestrator struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceOrchestratorMockRecorder
}

// MockWorkspaceOrchestratorMockRecorder is the mock recorder for MockWorkspaceOrchestrator.
type MockWorkspaceOrchestratorMockRecorder struct {
	mock *MockWorkspaceOrchestrator
}

// NewMockWorkspaceOrchestrator creates a new mock instance.
func NewMockWorkspaceOrchestrator(ctrl *gomock.Controller) *MockWorkspaceOrchestrator {
	mock := &MockWorkspaceOrchestrator{ctrl: ctrl}
	mock.recorder = &MockWorkspaceOrchestratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceOrchestrator) EXPECT() *MockWorkspaceOrchestratorMockRecorder {
	return m.recorder
}

// AddWorkspace mocks base method.
func (m *MockWorkspaceOrchestrator) AddWorkspace(c context.Context, req request.AddWorkspaceRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkspace", c, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWorkspace indicates an expected call of AddWorkspace.
func (mr *MockWorkspaceOrchestratorMockRecorder) AddWorkspace(c, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkspace", reflect.TypeOf((*MockWorkspaceOrchestrator)(nil).AddWorkspace), c, req)
}

// GetWorkspace mocks base method.
func (m *MockWorkspaceOrchestrator) GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", c, req)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockWorkspaceOrchestratorMockRecorder) GetWorkspace(c, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockWorkspaceOrchestrator)(nil).GetWorkspace), c, req)
}

// GetWorkspaces mocks base method.
func (m *MockWorkspaceOrchestrator) GetWorkspaces(c context.Context, uid int) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaces", c, uid)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaces indicates an expected call of GetWorkspaces.
func (mr *MockWorkspaceOrchestratorMockRecorder) GetWorkspaces(c, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaces", reflect.TypeOf((*MockWorkspaceOrchestrator)(nil).GetWorkspaces), c, uid)
}

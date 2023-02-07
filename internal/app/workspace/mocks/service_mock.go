// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/workspace/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	request "konntent-workspace-service/internal/app/dto/request"
	resource "konntent-workspace-service/internal/app/dto/resource"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddWorkspace mocks base method.
func (m *MockService) AddWorkspace(c context.Context, req request.AddWorkspaceRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkspace", c, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWorkspace indicates an expected call of AddWorkspace.
func (mr *MockServiceMockRecorder) AddWorkspace(c, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkspace", reflect.TypeOf((*MockService)(nil).AddWorkspace), c, req)
}

// GetWorkspace mocks base method.
func (m *MockService) GetWorkspace(c context.Context, req request.GetWorkspaceRequest) (*resource.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", c, req)
	ret0, _ := ret[0].(*resource.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockServiceMockRecorder) GetWorkspace(c, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockService)(nil).GetWorkspace), c, req)
}

// GetWorkspaces mocks base method.
func (m *MockService) GetWorkspaces(c context.Context, uid int) ([]resource.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaces", c, uid)
	ret0, _ := ret[0].([]resource.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkspaces indicates an expected call of GetWorkspaces.
func (mr *MockServiceMockRecorder) GetWorkspaces(c, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaces", reflect.TypeOf((*MockService)(nil).GetWorkspaces), c, uid)
}

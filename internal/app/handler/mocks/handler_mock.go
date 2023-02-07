// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/handler/workspace.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockWorkspaceHandler is a mock of WorkspaceHandler interface.
type MockWorkspaceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceHandlerMockRecorder
}

// MockWorkspaceHandlerMockRecorder is the mock recorder for MockWorkspaceHandler.
type MockWorkspaceHandlerMockRecorder struct {
	mock *MockWorkspaceHandler
}

// NewMockWorkspaceHandler creates a new mock instance.
func NewMockWorkspaceHandler(ctrl *gomock.Controller) *MockWorkspaceHandler {
	mock := &MockWorkspaceHandler{ctrl: ctrl}
	mock.recorder = &MockWorkspaceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaceHandler) EXPECT() *MockWorkspaceHandlerMockRecorder {
	return m.recorder
}

// AddWorkspace mocks base method.
func (m *MockWorkspaceHandler) AddWorkspace(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkspace", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWorkspace indicates an expected call of AddWorkspace.
func (mr *MockWorkspaceHandlerMockRecorder) AddWorkspace(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkspace", reflect.TypeOf((*MockWorkspaceHandler)(nil).AddWorkspace), c)
}

// GetWorkspace mocks base method.
func (m *MockWorkspaceHandler) GetWorkspace(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspace", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetWorkspace indicates an expected call of GetWorkspace.
func (mr *MockWorkspaceHandlerMockRecorder) GetWorkspace(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspace", reflect.TypeOf((*MockWorkspaceHandler)(nil).GetWorkspace), c)
}

// GetWorkspaces mocks base method.
func (m *MockWorkspaceHandler) GetWorkspaces(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaces", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetWorkspaces indicates an expected call of GetWorkspaces.
func (mr *MockWorkspaceHandlerMockRecorder) GetWorkspaces(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaces", reflect.TypeOf((*MockWorkspaceHandler)(nil).GetWorkspaces), c)
}
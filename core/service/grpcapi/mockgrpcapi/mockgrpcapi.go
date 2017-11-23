// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/service/grpcapi (interfaces: Manager,Registrable)

// Package mockgrpcapi is a generated GoMock package.
package mockgrpcapi

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	manager "github.com/stratumn/alice/core/manager"
	grpc "google.golang.org/grpc"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockManager) Find(arg0 string) (manager.Service, error) {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(manager.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockManagerMockRecorder) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockManager)(nil).Find), arg0)
}

// List mocks base method
func (m *MockManager) List() []string {
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	return ret0
}

// List indicates an expected call of List
func (mr *MockManagerMockRecorder) List() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManager)(nil).List))
}

// MockRegistrable is a mock of Registrable interface
type MockRegistrable struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrableMockRecorder
}

// MockRegistrableMockRecorder is the mock recorder for MockRegistrable
type MockRegistrableMockRecorder struct {
	mock *MockRegistrable
}

// NewMockRegistrable creates a new mock instance
func NewMockRegistrable(ctrl *gomock.Controller) *MockRegistrable {
	mock := &MockRegistrable{ctrl: ctrl}
	mock.recorder = &MockRegistrableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistrable) EXPECT() *MockRegistrableMockRecorder {
	return m.recorder
}

// AddToGRPCServer mocks base method
func (m *MockRegistrable) AddToGRPCServer(arg0 *grpc.Server) {
	m.ctrl.Call(m, "AddToGRPCServer", arg0)
}

// AddToGRPCServer indicates an expected call of AddToGRPCServer
func (mr *MockRegistrableMockRecorder) AddToGRPCServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToGRPCServer", reflect.TypeOf((*MockRegistrable)(nil).AddToGRPCServer), arg0)
}

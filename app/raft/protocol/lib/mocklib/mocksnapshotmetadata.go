// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/app/raft/protocol/lib (interfaces: SnapshotMetadata)

// Package mocklib is a generated GoMock package.
package mocklib

import (
	gomock "github.com/golang/mock/gomock"
	lib "github.com/stratumn/alice/app/raft/protocol/lib"
	reflect "reflect"
)

// MockSnapshotMetadata is a mock of SnapshotMetadata interface
type MockSnapshotMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotMetadataMockRecorder
}

// MockSnapshotMetadataMockRecorder is the mock recorder for MockSnapshotMetadata
type MockSnapshotMetadataMockRecorder struct {
	mock *MockSnapshotMetadata
}

// NewMockSnapshotMetadata creates a new mock instance
func NewMockSnapshotMetadata(ctrl *gomock.Controller) *MockSnapshotMetadata {
	mock := &MockSnapshotMetadata{ctrl: ctrl}
	mock.recorder = &MockSnapshotMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnapshotMetadata) EXPECT() *MockSnapshotMetadataMockRecorder {
	return m.recorder
}

// ConfState mocks base method
func (m *MockSnapshotMetadata) ConfState() lib.ConfState {
	ret := m.ctrl.Call(m, "ConfState")
	ret0, _ := ret[0].(lib.ConfState)
	return ret0
}

// ConfState indicates an expected call of ConfState
func (mr *MockSnapshotMetadataMockRecorder) ConfState() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfState", reflect.TypeOf((*MockSnapshotMetadata)(nil).ConfState))
}

// Index mocks base method
func (m *MockSnapshotMetadata) Index() uint64 {
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Index indicates an expected call of Index
func (mr *MockSnapshotMetadataMockRecorder) Index() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockSnapshotMetadata)(nil).Index))
}

// Term mocks base method
func (m *MockSnapshotMetadata) Term() uint64 {
	ret := m.ctrl.Call(m, "Term")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Term indicates an expected call of Term
func (mr *MockSnapshotMetadataMockRecorder) Term() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Term", reflect.TypeOf((*MockSnapshotMetadata)(nil).Term))
}
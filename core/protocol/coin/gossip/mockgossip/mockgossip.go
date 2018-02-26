// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/coin/gossip (interfaces: Gossip)

// Package mockgossip is a generated GoMock package.
package mockgossip

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	coin "github.com/stratumn/alice/pb/coin"
	reflect "reflect"
)

// MockGossip is a mock of Gossip interface
type MockGossip struct {
	ctrl     *gomock.Controller
	recorder *MockGossipMockRecorder
}

// MockGossipMockRecorder is the mock recorder for MockGossip
type MockGossipMockRecorder struct {
	mock *MockGossip
}

// NewMockGossip creates a new mock instance
func NewMockGossip(ctrl *gomock.Controller) *MockGossip {
	mock := &MockGossip{ctrl: ctrl}
	mock.recorder = &MockGossipMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGossip) EXPECT() *MockGossipMockRecorder {
	return m.recorder
}

// AddBlockListener mocks base method
func (m *MockGossip) AddBlockListener() chan *coin.Header {
	ret := m.ctrl.Call(m, "AddBlockListener")
	ret0, _ := ret[0].(chan *coin.Header)
	return ret0
}

// AddBlockListener indicates an expected call of AddBlockListener
func (mr *MockGossipMockRecorder) AddBlockListener() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlockListener", reflect.TypeOf((*MockGossip)(nil).AddBlockListener))
}

// Close mocks base method
func (m *MockGossip) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockGossipMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGossip)(nil).Close))
}

// ListenBlock mocks base method
func (m *MockGossip) ListenBlock(arg0 context.Context, arg1 func(*coin.Block) error, arg2 func([]byte) error) error {
	ret := m.ctrl.Call(m, "ListenBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenBlock indicates an expected call of ListenBlock
func (mr *MockGossipMockRecorder) ListenBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenBlock", reflect.TypeOf((*MockGossip)(nil).ListenBlock), arg0, arg1, arg2)
}

// ListenTx mocks base method
func (m *MockGossip) ListenTx(arg0 context.Context, arg1 func(*coin.Transaction) error) error {
	ret := m.ctrl.Call(m, "ListenTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenTx indicates an expected call of ListenTx
func (mr *MockGossipMockRecorder) ListenTx(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenTx", reflect.TypeOf((*MockGossip)(nil).ListenTx), arg0, arg1)
}

// PublishBlock mocks base method
func (m *MockGossip) PublishBlock(arg0 *coin.Block) error {
	ret := m.ctrl.Call(m, "PublishBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishBlock indicates an expected call of PublishBlock
func (mr *MockGossipMockRecorder) PublishBlock(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishBlock", reflect.TypeOf((*MockGossip)(nil).PublishBlock), arg0)
}

// PublishTx mocks base method
func (m *MockGossip) PublishTx(arg0 *coin.Transaction) error {
	ret := m.ctrl.Call(m, "PublishTx", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishTx indicates an expected call of PublishTx
func (mr *MockGossipMockRecorder) PublishTx(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTx", reflect.TypeOf((*MockGossip)(nil).PublishTx), arg0)
}

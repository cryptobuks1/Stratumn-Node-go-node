// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/streamutil (interfaces: Stream)

// Package mockstream is a generated GoMock package.
package mockstream

import (
	gomock "github.com/golang/mock/gomock"
	streamutil "github.com/stratumn/alice/core/streamutil"
	go_libp2p_net "gx/ipfs/QmXoz9o2PT3tEzf7hicegwex5UgVP54n3k82K7jrWFyN86/go-libp2p-net"
	reflect "reflect"
)

// MockStream is a mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStream) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockStreamMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close))
}

// Codec mocks base method
func (m *MockStream) Codec() streamutil.Codec {
	ret := m.ctrl.Call(m, "Codec")
	ret0, _ := ret[0].(streamutil.Codec)
	return ret0
}

// Codec indicates an expected call of Codec
func (mr *MockStreamMockRecorder) Codec() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Codec", reflect.TypeOf((*MockStream)(nil).Codec))
}

// Conn mocks base method
func (m *MockStream) Conn() go_libp2p_net.Conn {
	ret := m.ctrl.Call(m, "Conn")
	ret0, _ := ret[0].(go_libp2p_net.Conn)
	return ret0
}

// Conn indicates an expected call of Conn
func (mr *MockStreamMockRecorder) Conn() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockStream)(nil).Conn))
}
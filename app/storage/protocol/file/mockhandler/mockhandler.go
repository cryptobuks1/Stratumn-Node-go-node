// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/app/storage/protocol/file (interfaces: Handler,Reader)

// Package mockhandler is a generated GoMock package.
package mockhandler

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	file "github.com/stratumn/alice/app/storage/protocol/file"
	go_uuid "github.com/satori/go.uuid"
	reflect "reflect"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// AbortWrite mocks base method
func (m *MockHandler) AbortWrite(arg0 context.Context, arg1 go_uuid.UUID) error {
	ret := m.ctrl.Call(m, "AbortWrite", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortWrite indicates an expected call of AbortWrite
func (mr *MockHandlerMockRecorder) AbortWrite(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortWrite", reflect.TypeOf((*MockHandler)(nil).AbortWrite), arg0, arg1)
}

// BeginWrite mocks base method
func (m *MockHandler) BeginWrite(arg0 context.Context, arg1 string) (go_uuid.UUID, error) {
	ret := m.ctrl.Call(m, "BeginWrite", arg0, arg1)
	ret0, _ := ret[0].(go_uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginWrite indicates an expected call of BeginWrite
func (mr *MockHandlerMockRecorder) BeginWrite(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginWrite", reflect.TypeOf((*MockHandler)(nil).BeginWrite), arg0, arg1)
}

// EndWrite mocks base method
func (m *MockHandler) EndWrite(arg0 context.Context, arg1 go_uuid.UUID) ([]byte, error) {
	ret := m.ctrl.Call(m, "EndWrite", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndWrite indicates an expected call of EndWrite
func (mr *MockHandlerMockRecorder) EndWrite(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndWrite", reflect.TypeOf((*MockHandler)(nil).EndWrite), arg0, arg1)
}

// Exists mocks base method
func (m *MockHandler) Exists(arg0 context.Context, arg1 []byte) (bool, error) {
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockHandlerMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockHandler)(nil).Exists), arg0, arg1)
}

// Read mocks base method
func (m *MockHandler) Read(arg0 context.Context, arg1 []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockHandlerMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockHandler)(nil).Read), arg0, arg1)
}

// ReadChunks mocks base method
func (m *MockHandler) ReadChunks(arg0 context.Context, arg1 []byte, arg2 int, arg3 file.Reader) error {
	ret := m.ctrl.Call(m, "ReadChunks", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadChunks indicates an expected call of ReadChunks
func (mr *MockHandlerMockRecorder) ReadChunks(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadChunks", reflect.TypeOf((*MockHandler)(nil).ReadChunks), arg0, arg1, arg2, arg3)
}

// WriteChunk mocks base method
func (m *MockHandler) WriteChunk(arg0 context.Context, arg1 go_uuid.UUID, arg2 []byte) error {
	ret := m.ctrl.Call(m, "WriteChunk", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteChunk indicates an expected call of WriteChunk
func (mr *MockHandlerMockRecorder) WriteChunk(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteChunk", reflect.TypeOf((*MockHandler)(nil).WriteChunk), arg0, arg1, arg2)
}

// MockReader is a mock of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// OnChunk mocks base method
func (m *MockReader) OnChunk(arg0 []byte, arg1 string) error {
	ret := m.ctrl.Call(m, "OnChunk", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnChunk indicates an expected call of OnChunk
func (mr *MockReaderMockRecorder) OnChunk(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChunk", reflect.TypeOf((*MockReader)(nil).OnChunk), arg0, arg1)
}
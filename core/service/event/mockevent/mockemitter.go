// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/service/event (interfaces: Emitter)

package mockevent

import (
	gomock "github.com/golang/mock/gomock"
	event "github.com/stratumn/alice/grpc/event"
	reflect "reflect"
)

// MockEmitter is a mock of Emitter interface
type MockEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockEmitterMockRecorder
}

// MockEmitterMockRecorder is the mock recorder for MockEmitter
type MockEmitterMockRecorder struct {
	mock *MockEmitter
}

// NewMockEmitter creates a new mock instance
func NewMockEmitter(ctrl *gomock.Controller) *MockEmitter {
	mock := &MockEmitter{ctrl: ctrl}
	mock.recorder = &MockEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockEmitter) EXPECT() *MockEmitterMockRecorder {
	return _m.recorder
}

// AddListener mocks base method
func (_m *MockEmitter) AddListener() <-chan *event.Event {
	ret := _m.ctrl.Call(_m, "AddListener")
	ret0, _ := ret[0].(<-chan *event.Event)
	return ret0
}

// AddListener indicates an expected call of AddListener
func (_mr *MockEmitterMockRecorder) AddListener() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AddListener", reflect.TypeOf((*MockEmitter)(nil).AddListener))
}

// Close mocks base method
func (_m *MockEmitter) Close() {
	_m.ctrl.Call(_m, "Close")
}

// Close indicates an expected call of Close
func (_mr *MockEmitterMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockEmitter)(nil).Close))
}

// Emit mocks base method
func (_m *MockEmitter) Emit(_param0 *event.Event) {
	_m.ctrl.Call(_m, "Emit", _param0)
}

// Emit indicates an expected call of Emit
func (_mr *MockEmitterMockRecorder) Emit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Emit", reflect.TypeOf((*MockEmitter)(nil).Emit), arg0)
}

// GetListenersCount mocks base method
func (_m *MockEmitter) GetListenersCount() int {
	ret := _m.ctrl.Call(_m, "GetListenersCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetListenersCount indicates an expected call of GetListenersCount
func (_mr *MockEmitterMockRecorder) GetListenersCount() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetListenersCount", reflect.TypeOf((*MockEmitter)(nil).GetListenersCount))
}

// RemoveListener mocks base method
func (_m *MockEmitter) RemoveListener(_param0 <-chan *event.Event) {
	_m.ctrl.Call(_m, "RemoveListener", _param0)
}

// RemoveListener indicates an expected call of RemoveListener
func (_mr *MockEmitterMockRecorder) RemoveListener(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveListener", reflect.TypeOf((*MockEmitter)(nil).RemoveListener), arg0)
}
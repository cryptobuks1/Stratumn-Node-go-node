// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/protocol/coin/validator (interfaces: Validator)

package mockvalidator

import (
	gomock "github.com/golang/mock/gomock"
	state "github.com/stratumn/alice/core/protocol/coin/state"
	coin "github.com/stratumn/alice/pb/coin"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return _m.recorder
}

// ValidateBlock mocks base method
func (_m *MockValidator) ValidateBlock(_param0 *coin.Block, _param1 state.Reader) error {
	ret := _m.ctrl.Call(_m, "ValidateBlock", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateBlock indicates an expected call of ValidateBlock
func (_mr *MockValidatorMockRecorder) ValidateBlock(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ValidateBlock", reflect.TypeOf((*MockValidator)(nil).ValidateBlock), arg0, arg1)
}

// ValidateTx mocks base method
func (_m *MockValidator) ValidateTx(_param0 *coin.Transaction, _param1 state.Reader) error {
	ret := _m.ctrl.Call(_m, "ValidateTx", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTx indicates an expected call of ValidateTx
func (_mr *MockValidatorMockRecorder) ValidateTx(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ValidateTx", reflect.TypeOf((*MockValidator)(nil).ValidateTx), arg0, arg1)
}

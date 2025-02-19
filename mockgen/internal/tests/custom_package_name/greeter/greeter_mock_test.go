// Code generated by MockGen. DO NOT EDIT.
// Source: greeter.go
//
// Generated by this command:
//
//	mockgen -source greeter.go -destination greeter_mock_test.go -package greeter
//
// Package greeter is a generated GoMock package.
package greeter

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	client "go.uber.org/mock/mockgen/internal/tests/custom_package_name/client/v1"
)

// MockInputMaker is a mock of InputMaker interface.
type MockInputMaker struct {
	ctrl     *gomock.Controller
	recorder *MockInputMakerMockRecorder
}

// MockInputMakerMockRecorder is the mock recorder for MockInputMaker.
type MockInputMakerMockRecorder struct {
	mock *MockInputMaker
}

// NewMockInputMaker creates a new mock instance.
func NewMockInputMaker(ctrl *gomock.Controller) *MockInputMaker {
	mock := &MockInputMaker{ctrl: ctrl}
	mock.recorder = &MockInputMakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInputMaker) EXPECT() *MockInputMakerMockRecorder {
	return m.recorder
}

// MakeInput mocks base method.
func (m *MockInputMaker) MakeInput() client.GreetInput {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeInput")
	ret0, _ := ret[0].(client.GreetInput)
	return ret0
}

// MakeInput indicates an expected call of MakeInput.
func (mr *MockInputMakerMockRecorder) MakeInput() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeInput", reflect.TypeOf((*MockInputMaker)(nil).MakeInput))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/RealWorldGoSolution/sec11/mockgen (interfaces: GetSetter)

// Package internal is a generated GoMock package.
package internal

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGetSetter is a mock of GetSetter interface.
type MockGetSetter struct {
	ctrl     *gomock.Controller
	recorder *MockGetSetterMockRecorder
}

// MockGetSetterMockRecorder is the mock recorder for MockGetSetter.
type MockGetSetterMockRecorder struct {
	mock *MockGetSetter
}

// NewMockGetSetter creates a new mock instance.
func NewMockGetSetter(ctrl *gomock.Controller) *MockGetSetter {
	mock := &MockGetSetter{ctrl: ctrl}
	mock.recorder = &MockGetSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetSetter) EXPECT() *MockGetSetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetSetter) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetSetterMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetSetter)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockGetSetter) Set(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockGetSetterMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockGetSetter)(nil).Set), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/sensor/service/pipeline (interfaces: Fragment)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	common "github.com/stackrox/rox/central/sensor/service/common"
	central "github.com/stackrox/rox/generated/internalapi/central"
	reflect "reflect"
)

// MockFragment is a mock of Fragment interface
type MockFragment struct {
	ctrl     *gomock.Controller
	recorder *MockFragmentMockRecorder
}

// MockFragmentMockRecorder is the mock recorder for MockFragment
type MockFragmentMockRecorder struct {
	mock *MockFragment
}

// NewMockFragment creates a new mock instance
func NewMockFragment(ctrl *gomock.Controller) *MockFragment {
	mock := &MockFragment{ctrl: ctrl}
	mock.recorder = &MockFragmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFragment) EXPECT() *MockFragmentMockRecorder {
	return m.recorder
}

// Match mocks base method
func (m *MockFragment) Match(arg0 *central.MsgFromSensor) bool {
	ret := m.ctrl.Call(m, "Match", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match
func (mr *MockFragmentMockRecorder) Match(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockFragment)(nil).Match), arg0)
}

// OnFinish mocks base method
func (m *MockFragment) OnFinish(arg0 string) {
	m.ctrl.Call(m, "OnFinish", arg0)
}

// OnFinish indicates an expected call of OnFinish
func (mr *MockFragmentMockRecorder) OnFinish(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFinish", reflect.TypeOf((*MockFragment)(nil).OnFinish), arg0)
}

// Reconcile mocks base method
func (m *MockFragment) Reconcile(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconcile indicates an expected call of Reconcile
func (mr *MockFragmentMockRecorder) Reconcile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockFragment)(nil).Reconcile), arg0, arg1)
}

// Run mocks base method
func (m *MockFragment) Run(arg0 context.Context, arg1 string, arg2 *central.MsgFromSensor, arg3 common.MessageInjector) error {
	ret := m.ctrl.Call(m, "Run", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockFragmentMockRecorder) Run(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockFragment)(nil).Run), arg0, arg1, arg2, arg3)
}

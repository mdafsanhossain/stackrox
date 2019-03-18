// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/alert/store (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAlert mocks base method
func (m *MockStore) AddAlert(arg0 *storage.Alert) error {
	ret := m.ctrl.Call(m, "AddAlert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAlert indicates an expected call of AddAlert
func (mr *MockStoreMockRecorder) AddAlert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAlert", reflect.TypeOf((*MockStore)(nil).AddAlert), arg0)
}

// GetAlert mocks base method
func (m *MockStore) GetAlert(arg0 string) (*storage.Alert, bool, error) {
	ret := m.ctrl.Call(m, "GetAlert", arg0)
	ret0, _ := ret[0].(*storage.Alert)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlert indicates an expected call of GetAlert
func (mr *MockStoreMockRecorder) GetAlert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlert", reflect.TypeOf((*MockStore)(nil).GetAlert), arg0)
}

// GetAlertStates mocks base method
func (m *MockStore) GetAlertStates() ([]*storage.AlertState, error) {
	ret := m.ctrl.Call(m, "GetAlertStates")
	ret0, _ := ret[0].([]*storage.AlertState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlertStates indicates an expected call of GetAlertStates
func (mr *MockStoreMockRecorder) GetAlertStates() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlertStates", reflect.TypeOf((*MockStore)(nil).GetAlertStates))
}

// GetAlerts mocks base method
func (m *MockStore) GetAlerts(arg0 ...string) ([]*storage.Alert, error) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAlerts", varargs...)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlerts indicates an expected call of GetAlerts
func (mr *MockStoreMockRecorder) GetAlerts(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlerts", reflect.TypeOf((*MockStore)(nil).GetAlerts), arg0...)
}

// ListAlert mocks base method
func (m *MockStore) ListAlert(arg0 string) (*storage.ListAlert, bool, error) {
	ret := m.ctrl.Call(m, "ListAlert", arg0)
	ret0, _ := ret[0].(*storage.ListAlert)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAlert indicates an expected call of ListAlert
func (mr *MockStoreMockRecorder) ListAlert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlert", reflect.TypeOf((*MockStore)(nil).ListAlert), arg0)
}

// ListAlerts mocks base method
func (m *MockStore) ListAlerts() ([]*storage.ListAlert, error) {
	ret := m.ctrl.Call(m, "ListAlerts")
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlerts indicates an expected call of ListAlerts
func (mr *MockStoreMockRecorder) ListAlerts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlerts", reflect.TypeOf((*MockStore)(nil).ListAlerts))
}

// UpdateAlert mocks base method
func (m *MockStore) UpdateAlert(arg0 *storage.Alert) error {
	ret := m.ctrl.Call(m, "UpdateAlert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAlert indicates an expected call of UpdateAlert
func (mr *MockStoreMockRecorder) UpdateAlert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlert", reflect.TypeOf((*MockStore)(nil).UpdateAlert), arg0)
}

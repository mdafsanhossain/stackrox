// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/networkflow/store (interfaces: ClusterStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	store "github.com/stackrox/rox/central/networkflow/store"
	reflect "reflect"
)

// MockClusterStore is a mock of ClusterStore interface
type MockClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStoreMockRecorder
}

// MockClusterStoreMockRecorder is the mock recorder for MockClusterStore
type MockClusterStoreMockRecorder struct {
	mock *MockClusterStore
}

// NewMockClusterStore creates a new mock instance
func NewMockClusterStore(ctrl *gomock.Controller) *MockClusterStore {
	mock := &MockClusterStore{ctrl: ctrl}
	mock.recorder = &MockClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterStore) EXPECT() *MockClusterStoreMockRecorder {
	return m.recorder
}

// CreateFlowStore mocks base method
func (m *MockClusterStore) CreateFlowStore(arg0 string) (store.FlowStore, error) {
	ret := m.ctrl.Call(m, "CreateFlowStore", arg0)
	ret0, _ := ret[0].(store.FlowStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlowStore indicates an expected call of CreateFlowStore
func (mr *MockClusterStoreMockRecorder) CreateFlowStore(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlowStore", reflect.TypeOf((*MockClusterStore)(nil).CreateFlowStore), arg0)
}

// GetFlowStore mocks base method
func (m *MockClusterStore) GetFlowStore(arg0 string) store.FlowStore {
	ret := m.ctrl.Call(m, "GetFlowStore", arg0)
	ret0, _ := ret[0].(store.FlowStore)
	return ret0
}

// GetFlowStore indicates an expected call of GetFlowStore
func (mr *MockClusterStoreMockRecorder) GetFlowStore(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowStore", reflect.TypeOf((*MockClusterStore)(nil).GetFlowStore), arg0)
}

// RemoveFlowStore mocks base method
func (m *MockClusterStore) RemoveFlowStore(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveFlowStore", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFlowStore indicates an expected call of RemoveFlowStore
func (mr *MockClusterStoreMockRecorder) RemoveFlowStore(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFlowStore", reflect.TypeOf((*MockClusterStore)(nil).RemoveFlowStore), arg0)
}

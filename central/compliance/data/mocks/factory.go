// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/compliance/data (interfaces: RepositoryFactory)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	framework "github.com/stackrox/rox/central/compliance/framework"
	compliance "github.com/stackrox/rox/generated/internalapi/compliance"
	reflect "reflect"
)

// MockRepositoryFactory is a mock of RepositoryFactory interface
type MockRepositoryFactory struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryFactoryMockRecorder
}

// MockRepositoryFactoryMockRecorder is the mock recorder for MockRepositoryFactory
type MockRepositoryFactoryMockRecorder struct {
	mock *MockRepositoryFactory
}

// NewMockRepositoryFactory creates a new mock instance
func NewMockRepositoryFactory(ctrl *gomock.Controller) *MockRepositoryFactory {
	mock := &MockRepositoryFactory{ctrl: ctrl}
	mock.recorder = &MockRepositoryFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryFactory) EXPECT() *MockRepositoryFactoryMockRecorder {
	return m.recorder
}

// CreateDataRepository mocks base method
func (m *MockRepositoryFactory) CreateDataRepository(arg0 context.Context, arg1 framework.ComplianceDomain, arg2 map[string]*compliance.ComplianceReturn) (framework.ComplianceDataRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataRepository", arg0, arg1, arg2)
	ret0, _ := ret[0].(framework.ComplianceDataRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataRepository indicates an expected call of CreateDataRepository
func (mr *MockRepositoryFactoryMockRecorder) CreateDataRepository(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDataRepository", reflect.TypeOf((*MockRepositoryFactory)(nil).CreateDataRepository), arg0, arg1, arg2)
}

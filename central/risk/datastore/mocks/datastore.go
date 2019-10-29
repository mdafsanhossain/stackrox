// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// Search mocks base method
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockDataStoreMockRecorder) Search(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchRawRisks mocks base method
func (m *MockDataStore) SearchRawRisks(ctx context.Context, q *v1.Query) ([]*storage.Risk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawRisks", ctx, q)
	ret0, _ := ret[0].([]*storage.Risk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawRisks indicates an expected call of SearchRawRisks
func (mr *MockDataStoreMockRecorder) SearchRawRisks(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawRisks", reflect.TypeOf((*MockDataStore)(nil).SearchRawRisks), ctx, q)
}

// GetRisk mocks base method
func (m *MockDataStore) GetRisk(ctx context.Context, subjectID string, subjectType storage.RiskSubjectType) (*storage.Risk, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRisk", ctx, subjectID, subjectType)
	ret0, _ := ret[0].(*storage.Risk)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRisk indicates an expected call of GetRisk
func (mr *MockDataStoreMockRecorder) GetRisk(ctx, subjectID, subjectType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRisk", reflect.TypeOf((*MockDataStore)(nil).GetRisk), ctx, subjectID, subjectType)
}

// GetRiskByIndicators mocks base method
func (m *MockDataStore) GetRiskByIndicators(ctx context.Context, subjectID string, subjectType storage.RiskSubjectType, riskIndicatorNames []string) (*storage.Risk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRiskByIndicators", ctx, subjectID, subjectType, riskIndicatorNames)
	ret0, _ := ret[0].(*storage.Risk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRiskByIndicators indicates an expected call of GetRiskByIndicators
func (mr *MockDataStoreMockRecorder) GetRiskByIndicators(ctx, subjectID, subjectType, riskIndicatorNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRiskByIndicators", reflect.TypeOf((*MockDataStore)(nil).GetRiskByIndicators), ctx, subjectID, subjectType, riskIndicatorNames)
}

// UpsertRisk mocks base method
func (m *MockDataStore) UpsertRisk(ctx context.Context, risk *storage.Risk) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRisk", ctx, risk)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRisk indicates an expected call of UpsertRisk
func (mr *MockDataStoreMockRecorder) UpsertRisk(ctx, risk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRisk", reflect.TypeOf((*MockDataStore)(nil).UpsertRisk), ctx, risk)
}

// RemoveRisk mocks base method
func (m *MockDataStore) RemoveRisk(ctx context.Context, subjectID string, subjectType storage.RiskSubjectType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRisk", ctx, subjectID, subjectType)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRisk indicates an expected call of RemoveRisk
func (mr *MockDataStoreMockRecorder) RemoveRisk(ctx, subjectID, subjectType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRisk", reflect.TypeOf((*MockDataStore)(nil).RemoveRisk), ctx, subjectID, subjectType)
}

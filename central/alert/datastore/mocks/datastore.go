// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockDataStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDataStoreMockRecorder) Count(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// CountAlerts mocks base method.
func (m *MockDataStore) CountAlerts(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAlerts", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAlerts indicates an expected call of CountAlerts.
func (mr *MockDataStoreMockRecorder) CountAlerts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAlerts", reflect.TypeOf((*MockDataStore)(nil).CountAlerts), ctx)
}

// DeleteAlerts mocks base method.
func (m *MockDataStore) DeleteAlerts(ctx context.Context, ids ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAlerts", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlerts indicates an expected call of DeleteAlerts.
func (mr *MockDataStoreMockRecorder) DeleteAlerts(ctx interface{}, ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlerts", reflect.TypeOf((*MockDataStore)(nil).DeleteAlerts), varargs...)
}

// GetAlert mocks base method.
func (m *MockDataStore) GetAlert(ctx context.Context, id string) (*storage.Alert, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlert", ctx, id)
	ret0, _ := ret[0].(*storage.Alert)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAlert indicates an expected call of GetAlert.
func (mr *MockDataStoreMockRecorder) GetAlert(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlert", reflect.TypeOf((*MockDataStore)(nil).GetAlert), ctx, id)
}

// MarkAlertStaleBatch mocks base method.
func (m *MockDataStore) MarkAlertStaleBatch(ctx context.Context, id ...string) ([]*storage.Alert, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range id {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MarkAlertStaleBatch", varargs...)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkAlertStaleBatch indicates an expected call of MarkAlertStaleBatch.
func (mr *MockDataStoreMockRecorder) MarkAlertStaleBatch(ctx interface{}, id ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, id...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAlertStaleBatch", reflect.TypeOf((*MockDataStore)(nil).MarkAlertStaleBatch), varargs...)
}

// Search mocks base method.
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDataStoreMockRecorder) Search(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchAlerts mocks base method.
func (m *MockDataStore) SearchAlerts(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAlerts", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAlerts indicates an expected call of SearchAlerts.
func (mr *MockDataStoreMockRecorder) SearchAlerts(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchAlerts), ctx, q)
}

// SearchListAlerts mocks base method.
func (m *MockDataStore) SearchListAlerts(ctx context.Context, q *v1.Query) ([]*storage.ListAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchListAlerts", ctx, q)
	ret0, _ := ret[0].([]*storage.ListAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchListAlerts indicates an expected call of SearchListAlerts.
func (mr *MockDataStoreMockRecorder) SearchListAlerts(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchListAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchListAlerts), ctx, q)
}

// SearchRawAlerts mocks base method.
func (m *MockDataStore) SearchRawAlerts(ctx context.Context, q *v1.Query) ([]*storage.Alert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawAlerts", ctx, q)
	ret0, _ := ret[0].([]*storage.Alert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawAlerts indicates an expected call of SearchRawAlerts.
func (mr *MockDataStoreMockRecorder) SearchRawAlerts(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawAlerts", reflect.TypeOf((*MockDataStore)(nil).SearchRawAlerts), ctx, q)
}

// UpsertAlert mocks base method.
func (m *MockDataStore) UpsertAlert(ctx context.Context, alert *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertAlert", ctx, alert)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAlert indicates an expected call of UpsertAlert.
func (mr *MockDataStoreMockRecorder) UpsertAlert(ctx, alert interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAlert", reflect.TypeOf((*MockDataStore)(nil).UpsertAlert), ctx, alert)
}

// UpsertAlerts mocks base method.
func (m *MockDataStore) UpsertAlerts(ctx context.Context, alerts []*storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertAlerts", ctx, alerts)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertAlerts indicates an expected call of UpsertAlerts.
func (mr *MockDataStoreMockRecorder) UpsertAlerts(ctx, alerts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAlerts", reflect.TypeOf((*MockDataStore)(nil).UpsertAlerts), ctx, alerts)
}

// WalkAll mocks base method.
func (m *MockDataStore) WalkAll(ctx context.Context, fn func(*storage.ListAlert) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkAll", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkAll indicates an expected call of WalkAll.
func (mr *MockDataStoreMockRecorder) WalkAll(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkAll", reflect.TypeOf((*MockDataStore)(nil).WalkAll), ctx, fn)
}

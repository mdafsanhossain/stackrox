// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/central/image/datastore (interfaces: DataStore)

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

// CountImages mocks base method
func (m *MockDataStore) CountImages(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountImages", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountImages indicates an expected call of CountImages
func (mr *MockDataStoreMockRecorder) CountImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountImages", reflect.TypeOf((*MockDataStore)(nil).CountImages), arg0)
}

// DeleteImages mocks base method
func (m *MockDataStore) DeleteImages(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteImages", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImages indicates an expected call of DeleteImages
func (mr *MockDataStoreMockRecorder) DeleteImages(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImages", reflect.TypeOf((*MockDataStore)(nil).DeleteImages), varargs...)
}

// Exists mocks base method
func (m *MockDataStore) Exists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockDataStoreMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDataStore)(nil).Exists), arg0, arg1)
}

// GetImage mocks base method
func (m *MockDataStore) GetImage(arg0 context.Context, arg1 string) (*storage.Image, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", arg0, arg1)
	ret0, _ := ret[0].(*storage.Image)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetImage indicates an expected call of GetImage
func (mr *MockDataStoreMockRecorder) GetImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockDataStore)(nil).GetImage), arg0, arg1)
}

// GetImagesBatch mocks base method
func (m *MockDataStore) GetImagesBatch(arg0 context.Context, arg1 []string) ([]*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagesBatch", arg0, arg1)
	ret0, _ := ret[0].([]*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImagesBatch indicates an expected call of GetImagesBatch
func (mr *MockDataStoreMockRecorder) GetImagesBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagesBatch", reflect.TypeOf((*MockDataStore)(nil).GetImagesBatch), arg0, arg1)
}

// ListImage mocks base method
func (m *MockDataStore) ListImage(arg0 context.Context, arg1 string) (*storage.ListImage, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImage", arg0, arg1)
	ret0, _ := ret[0].(*storage.ListImage)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListImage indicates an expected call of ListImage
func (mr *MockDataStoreMockRecorder) ListImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImage", reflect.TypeOf((*MockDataStore)(nil).ListImage), arg0, arg1)
}

// Search mocks base method
func (m *MockDataStore) Search(arg0 context.Context, arg1 *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockDataStoreMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), arg0, arg1)
}

// SearchImages mocks base method
func (m *MockDataStore) SearchImages(arg0 context.Context, arg1 *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchImages", arg0, arg1)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchImages indicates an expected call of SearchImages
func (mr *MockDataStoreMockRecorder) SearchImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchImages", reflect.TypeOf((*MockDataStore)(nil).SearchImages), arg0, arg1)
}

// SearchListImages mocks base method
func (m *MockDataStore) SearchListImages(arg0 context.Context, arg1 *v1.Query) ([]*storage.ListImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchListImages", arg0, arg1)
	ret0, _ := ret[0].([]*storage.ListImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchListImages indicates an expected call of SearchListImages
func (mr *MockDataStoreMockRecorder) SearchListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchListImages", reflect.TypeOf((*MockDataStore)(nil).SearchListImages), arg0, arg1)
}

// SearchRawImages mocks base method
func (m *MockDataStore) SearchRawImages(arg0 context.Context, arg1 *v1.Query) ([]*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawImages", arg0, arg1)
	ret0, _ := ret[0].([]*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawImages indicates an expected call of SearchRawImages
func (mr *MockDataStoreMockRecorder) SearchRawImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawImages", reflect.TypeOf((*MockDataStore)(nil).SearchRawImages), arg0, arg1)
}

// UpsertImage mocks base method
func (m *MockDataStore) UpsertImage(arg0 context.Context, arg1 *storage.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertImage indicates an expected call of UpsertImage
func (mr *MockDataStoreMockRecorder) UpsertImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertImage", reflect.TypeOf((*MockDataStore)(nil).UpsertImage), arg0, arg1)
}

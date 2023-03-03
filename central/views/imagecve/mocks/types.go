// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	imagecve "github.com/stackrox/rox/central/views/imagecve"
	v1 "github.com/stackrox/rox/generated/api/v1"
)

// MockCveCore is a mock of CveCore interface.
type MockCveCore struct {
	ctrl     *gomock.Controller
	recorder *MockCveCoreMockRecorder
}

// MockCveCoreMockRecorder is the mock recorder for MockCveCore.
type MockCveCoreMockRecorder struct {
	mock *MockCveCore
}

// NewMockCveCore creates a new mock instance.
func NewMockCveCore(ctrl *gomock.Controller) *MockCveCore {
	mock := &MockCveCore{ctrl: ctrl}
	mock.recorder = &MockCveCoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCveCore) EXPECT() *MockCveCoreMockRecorder {
	return m.recorder
}

// GetAffectedImages mocks base method.
func (m *MockCveCore) GetAffectedImages() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAffectedImages")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAffectedImages indicates an expected call of GetAffectedImages.
func (mr *MockCveCoreMockRecorder) GetAffectedImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAffectedImages", reflect.TypeOf((*MockCveCore)(nil).GetAffectedImages))
}

// GetCVE mocks base method.
func (m *MockCveCore) GetCVE() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCVE")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCVE indicates an expected call of GetCVE.
func (mr *MockCveCoreMockRecorder) GetCVE() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCVE", reflect.TypeOf((*MockCveCore)(nil).GetCVE))
}

// GetFirstDiscoveredInSystem mocks base method.
func (m *MockCveCore) GetFirstDiscoveredInSystem() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstDiscoveredInSystem")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetFirstDiscoveredInSystem indicates an expected call of GetFirstDiscoveredInSystem.
func (mr *MockCveCoreMockRecorder) GetFirstDiscoveredInSystem() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstDiscoveredInSystem", reflect.TypeOf((*MockCveCore)(nil).GetFirstDiscoveredInSystem))
}

// GetImagesBySeverity mocks base method.
func (m *MockCveCore) GetImagesBySeverity() *imagecve.ResourceCountByCVESeverity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagesBySeverity")
	ret0, _ := ret[0].(*imagecve.ResourceCountByCVESeverity)
	return ret0
}

// GetImagesBySeverity indicates an expected call of GetImagesBySeverity.
func (mr *MockCveCoreMockRecorder) GetImagesBySeverity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagesBySeverity", reflect.TypeOf((*MockCveCore)(nil).GetImagesBySeverity))
}

// GetTopCVSS mocks base method.
func (m *MockCveCore) GetTopCVSS() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopCVSS")
	ret0, _ := ret[0].(float32)
	return ret0
}

// GetTopCVSS indicates an expected call of GetTopCVSS.
func (mr *MockCveCoreMockRecorder) GetTopCVSS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopCVSS", reflect.TypeOf((*MockCveCore)(nil).GetTopCVSS))
}

// MockCveView is a mock of CveView interface.
type MockCveView struct {
	ctrl     *gomock.Controller
	recorder *MockCveViewMockRecorder
}

// MockCveViewMockRecorder is the mock recorder for MockCveView.
type MockCveViewMockRecorder struct {
	mock *MockCveView
}

// NewMockCveView creates a new mock instance.
func NewMockCveView(ctrl *gomock.Controller) *MockCveView {
	mock := &MockCveView{ctrl: ctrl}
	mock.recorder = &MockCveViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCveView) EXPECT() *MockCveViewMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockCveView) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockCveViewMockRecorder) Count(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCveView)(nil).Count), ctx, q)
}

// Get mocks base method.
func (m *MockCveView) Get(ctx context.Context, q *v1.Query) ([]imagecve.CveCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, q)
	ret0, _ := ret[0].([]imagecve.CveCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCveViewMockRecorder) Get(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCveView)(nil).Get), ctx, q)
}

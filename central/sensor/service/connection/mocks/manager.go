// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	connection "github.com/stackrox/rox/central/sensor/service/connection"
	pipeline "github.com/stackrox/rox/central/sensor/service/pipeline"
	central "github.com/stackrox/rox/generated/internalapi/central"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
	time "time"
)

// MockClusterManager is a mock of ClusterManager interface
type MockClusterManager struct {
	ctrl     *gomock.Controller
	recorder *MockClusterManagerMockRecorder
}

// MockClusterManagerMockRecorder is the mock recorder for MockClusterManager
type MockClusterManagerMockRecorder struct {
	mock *MockClusterManager
}

// NewMockClusterManager creates a new mock instance
func NewMockClusterManager(ctrl *gomock.Controller) *MockClusterManager {
	mock := &MockClusterManager{ctrl: ctrl}
	mock.recorder = &MockClusterManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterManager) EXPECT() *MockClusterManagerMockRecorder {
	return m.recorder
}

// UpdateClusterContactTimes mocks base method
func (m *MockClusterManager) UpdateClusterContactTimes(ctx context.Context, time time.Time, clusterID ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, time}
	for _, a := range clusterID {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateClusterContactTimes", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClusterContactTimes indicates an expected call of UpdateClusterContactTimes
func (mr *MockClusterManagerMockRecorder) UpdateClusterContactTimes(ctx, time interface{}, clusterID ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, time}, clusterID...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterContactTimes", reflect.TypeOf((*MockClusterManager)(nil).UpdateClusterContactTimes), varargs...)
}

// UpdateClusterUpgradeStatus mocks base method
func (m *MockClusterManager) UpdateClusterUpgradeStatus(ctx context.Context, clusterID string, status *storage.ClusterUpgradeStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClusterUpgradeStatus", ctx, clusterID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClusterUpgradeStatus indicates an expected call of UpdateClusterUpgradeStatus
func (mr *MockClusterManagerMockRecorder) UpdateClusterUpgradeStatus(ctx, clusterID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClusterUpgradeStatus", reflect.TypeOf((*MockClusterManager)(nil).UpdateClusterUpgradeStatus), ctx, clusterID, status)
}

// GetCluster mocks base method
func (m *MockClusterManager) GetCluster(ctx context.Context, id string) (*storage.Cluster, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", ctx, id)
	ret0, _ := ret[0].(*storage.Cluster)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockClusterManagerMockRecorder) GetCluster(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterManager)(nil).GetCluster), ctx, id)
}

// GetClusters mocks base method
func (m *MockClusterManager) GetClusters(ctx context.Context) ([]*storage.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusters", ctx)
	ret0, _ := ret[0].([]*storage.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusters indicates an expected call of GetClusters
func (mr *MockClusterManagerMockRecorder) GetClusters(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusters", reflect.TypeOf((*MockClusterManager)(nil).GetClusters), ctx)
}

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockManager) Start(mgr connection.ClusterManager) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", mgr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockManagerMockRecorder) Start(mgr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start), mgr)
}

// HandleConnection mocks base method
func (m *MockManager) HandleConnection(ctx context.Context, clusterID string, pf pipeline.Factory, server central.SensorService_CommunicateServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleConnection", ctx, clusterID, pf, server)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleConnection indicates an expected call of HandleConnection
func (mr *MockManagerMockRecorder) HandleConnection(ctx, clusterID, pf, server interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleConnection", reflect.TypeOf((*MockManager)(nil).HandleConnection), ctx, clusterID, pf, server)
}

// GetConnection mocks base method
func (m *MockManager) GetConnection(clusterID string) connection.SensorConnection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", clusterID)
	ret0, _ := ret[0].(connection.SensorConnection)
	return ret0
}

// GetConnection indicates an expected call of GetConnection
func (mr *MockManagerMockRecorder) GetConnection(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockManager)(nil).GetConnection), clusterID)
}

// GetActiveConnections mocks base method
func (m *MockManager) GetActiveConnections() []connection.SensorConnection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveConnections")
	ret0, _ := ret[0].([]connection.SensorConnection)
	return ret0
}

// GetActiveConnections indicates an expected call of GetActiveConnections
func (mr *MockManagerMockRecorder) GetActiveConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveConnections", reflect.TypeOf((*MockManager)(nil).GetActiveConnections))
}

// TriggerUpgrade mocks base method
func (m *MockManager) TriggerUpgrade(ctx context.Context, clusterID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerUpgrade", ctx, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// TriggerUpgrade indicates an expected call of TriggerUpgrade
func (mr *MockManagerMockRecorder) TriggerUpgrade(ctx, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerUpgrade", reflect.TypeOf((*MockManager)(nil).TriggerUpgrade), ctx, clusterID)
}

// ProcessCheckInFromUpgrader mocks base method
func (m *MockManager) ProcessCheckInFromUpgrader(ctx context.Context, clusterID string, req *central.UpgradeCheckInFromUpgraderRequest) (*central.UpgradeCheckInFromUpgraderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessCheckInFromUpgrader", ctx, clusterID, req)
	ret0, _ := ret[0].(*central.UpgradeCheckInFromUpgraderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessCheckInFromUpgrader indicates an expected call of ProcessCheckInFromUpgrader
func (mr *MockManagerMockRecorder) ProcessCheckInFromUpgrader(ctx, clusterID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessCheckInFromUpgrader", reflect.TypeOf((*MockManager)(nil).ProcessCheckInFromUpgrader), ctx, clusterID, req)
}

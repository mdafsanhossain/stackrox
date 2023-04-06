// Code generated by MockGen. DO NOT EDIT.
// Source: detector.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	central "github.com/stackrox/rox/generated/internalapi/central"
	storage "github.com/stackrox/rox/generated/storage"
	centralsensor "github.com/stackrox/rox/pkg/centralsensor"
	common "github.com/stackrox/rox/sensor/common"
	grpc "google.golang.org/grpc"
)

// MockDetector is a mock of Detector interface.
type MockDetector struct {
	ctrl     *gomock.Controller
	recorder *MockDetectorMockRecorder
}

// MockDetectorMockRecorder is the mock recorder for MockDetector.
type MockDetectorMockRecorder struct {
	mock *MockDetector
}

// NewMockDetector creates a new mock instance.
func NewMockDetector(ctrl *gomock.Controller) *MockDetector {
	mock := &MockDetector{ctrl: ctrl}
	mock.recorder = &MockDetectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetector) EXPECT() *MockDetectorMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockDetector) Capabilities() []centralsensor.SensorCapability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].([]centralsensor.SensorCapability)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockDetectorMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockDetector)(nil).Capabilities))
}

// Notify mocks base method.
func (m *MockDetector) Notify(e common.SensorComponentEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", e)
}

// Notify indicates an expected call of Notify.
func (mr *MockDetectorMockRecorder) Notify(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockDetector)(nil).Notify), e)
}

// ProcessDeployment mocks base method.
func (m *MockDetector) ProcessDeployment(deployment *storage.Deployment, action central.ResourceAction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessDeployment", deployment, action)
}

// ProcessDeployment indicates an expected call of ProcessDeployment.
func (mr *MockDetectorMockRecorder) ProcessDeployment(deployment, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDeployment", reflect.TypeOf((*MockDetector)(nil).ProcessDeployment), deployment, action)
}

// ProcessIndicator mocks base method.
func (m *MockDetector) ProcessIndicator(indicator *storage.ProcessIndicator) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessIndicator", indicator)
}

// ProcessIndicator indicates an expected call of ProcessIndicator.
func (mr *MockDetectorMockRecorder) ProcessIndicator(indicator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessIndicator", reflect.TypeOf((*MockDetector)(nil).ProcessIndicator), indicator)
}

// ProcessMessage mocks base method.
func (m *MockDetector) ProcessMessage(msg *central.MsgToSensor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage.
func (mr *MockDetectorMockRecorder) ProcessMessage(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockDetector)(nil).ProcessMessage), msg)
}

// ProcessNetworkFlow mocks base method.
func (m *MockDetector) ProcessNetworkFlow(flow *storage.NetworkFlow) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ProcessNetworkFlow", flow)
}

// ProcessNetworkFlow indicates an expected call of ProcessNetworkFlow.
func (mr *MockDetectorMockRecorder) ProcessNetworkFlow(flow interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNetworkFlow", reflect.TypeOf((*MockDetector)(nil).ProcessNetworkFlow), flow)
}

// ProcessPolicySync mocks base method.
func (m *MockDetector) ProcessPolicySync(sync *central.PolicySync) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPolicySync", sync)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPolicySync indicates an expected call of ProcessPolicySync.
func (mr *MockDetectorMockRecorder) ProcessPolicySync(sync interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPolicySync", reflect.TypeOf((*MockDetector)(nil).ProcessPolicySync), sync)
}

// ProcessReassessPolicies mocks base method.
func (m *MockDetector) ProcessReassessPolicies() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessReassessPolicies")
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessReassessPolicies indicates an expected call of ProcessReassessPolicies.
func (mr *MockDetectorMockRecorder) ProcessReassessPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessReassessPolicies", reflect.TypeOf((*MockDetector)(nil).ProcessReassessPolicies))
}

// ProcessReprocessDeployments mocks base method.
func (m *MockDetector) ProcessReprocessDeployments() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessReprocessDeployments")
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessReprocessDeployments indicates an expected call of ProcessReprocessDeployments.
func (mr *MockDetectorMockRecorder) ProcessReprocessDeployments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessReprocessDeployments", reflect.TypeOf((*MockDetector)(nil).ProcessReprocessDeployments))
}

// ProcessUpdatedImage mocks base method.
func (m *MockDetector) ProcessUpdatedImage(image *storage.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessUpdatedImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessUpdatedImage indicates an expected call of ProcessUpdatedImage.
func (mr *MockDetectorMockRecorder) ProcessUpdatedImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessUpdatedImage", reflect.TypeOf((*MockDetector)(nil).ProcessUpdatedImage), image)
}

// ReprocessDeployments mocks base method.
func (m *MockDetector) ReprocessDeployments(deploymentIDs ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range deploymentIDs {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "ReprocessDeployments", varargs...)
}

// ReprocessDeployments indicates an expected call of ReprocessDeployments.
func (mr *MockDetectorMockRecorder) ReprocessDeployments(deploymentIDs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReprocessDeployments", reflect.TypeOf((*MockDetector)(nil).ReprocessDeployments), deploymentIDs...)
}

// ResponsesC mocks base method.
func (m *MockDetector) ResponsesC() <-chan *central.MsgFromSensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponsesC")
	ret0, _ := ret[0].(<-chan *central.MsgFromSensor)
	return ret0
}

// ResponsesC indicates an expected call of ResponsesC.
func (mr *MockDetectorMockRecorder) ResponsesC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponsesC", reflect.TypeOf((*MockDetector)(nil).ResponsesC))
}

// SetCentralGRPCClient mocks base method.
func (m *MockDetector) SetCentralGRPCClient(cc grpc.ClientConnInterface) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCentralGRPCClient", cc)
}

// SetCentralGRPCClient indicates an expected call of SetCentralGRPCClient.
func (mr *MockDetectorMockRecorder) SetCentralGRPCClient(cc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCentralGRPCClient", reflect.TypeOf((*MockDetector)(nil).SetCentralGRPCClient), cc)
}

// Start mocks base method.
func (m *MockDetector) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockDetectorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockDetector)(nil).Start))
}

// Stop mocks base method.
func (m *MockDetector) Stop(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", err)
}

// Stop indicates an expected call of Stop.
func (mr *MockDetectorMockRecorder) Stop(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockDetector)(nil).Stop), err)
}

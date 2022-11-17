// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_ecs is a generated GoMock package.
package mock_ecs

import (
	reflect "reflect"

	aws "github.com/aws/aws-sdk-go-v2/aws"
	gomock "github.com/golang/mock/gomock"
)

// MockECSServicer is a mock of ECSServicer interface.
type MockECSServicer struct {
	ctrl     *gomock.Controller
	recorder *MockECSServicerMockRecorder
}

// MockECSServicerMockRecorder is the mock recorder for MockECSServicer.
type MockECSServicerMockRecorder struct {
	mock *MockECSServicer
}

// NewMockECSServicer creates a new mock instance.
func NewMockECSServicer(ctrl *gomock.Controller) *MockECSServicer {
	mock := &MockECSServicer{ctrl: ctrl}
	mock.recorder = &MockECSServicerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECSServicer) EXPECT() *MockECSServicerMockRecorder {
	return m.recorder
}

// EcsExec mocks base method.
func (m *MockECSServicer) EcsExec(arg0 aws.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EcsExec", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EcsExec indicates an expected call of EcsExec.
func (mr *MockECSServicerMockRecorder) EcsExec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EcsExec", reflect.TypeOf((*MockECSServicer)(nil).EcsExec), arg0)
}

// StartEcs mocks base method.
func (m *MockECSServicer) StartEcs(arg0 aws.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartEcs", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartEcs indicates an expected call of StartEcs.
func (mr *MockECSServicerMockRecorder) StartEcs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartEcs", reflect.TypeOf((*MockECSServicer)(nil).StartEcs), arg0)
}

// StopEcsTask mocks base method.
func (m *MockECSServicer) StopEcsTask(arg0 aws.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopEcsTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopEcsTask indicates an expected call of StopEcsTask.
func (mr *MockECSServicerMockRecorder) StopEcsTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopEcsTask", reflect.TypeOf((*MockECSServicer)(nil).StopEcsTask), arg0)
}

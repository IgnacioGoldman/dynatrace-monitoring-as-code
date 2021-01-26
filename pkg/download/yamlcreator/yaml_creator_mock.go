// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dynatrace-oss/dynatrace-monitoring-as-code/pkg/download/yamlcreator (interfaces: YamlCreator)

// Package mock_yamlcreator is a generated GoMock package.
package yamlcreator

import (
	files "github.com/dynatrace-oss/dynatrace-monitoring-as-code/pkg/util/files"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockYamlCreator is a mock of YamlCreator interface
type MockYamlCreator struct {
	ctrl     *gomock.Controller
	recorder *MockYamlCreatorMockRecorder
}

// MockYamlCreatorMockRecorder is the mock recorder for MockYamlCreator
type MockYamlCreatorMockRecorder struct {
	mock *MockYamlCreator
}

// NewMockYamlCreator creates a new mock instance
func NewMockYamlCreator(ctrl *gomock.Controller) *MockYamlCreator {
	mock := &MockYamlCreator{ctrl: ctrl}
	mock.recorder = &MockYamlCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockYamlCreator) EXPECT() *MockYamlCreatorMockRecorder {
	return m.recorder
}

// AddConfig mocks base method
func (m *MockYamlCreator) AddConfig(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddConfig", arg0, arg1)
}

// AddConfig indicates an expected call of AddConfig
func (mr *MockYamlCreatorMockRecorder) AddConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConfig", reflect.TypeOf((*MockYamlCreator)(nil).AddConfig), arg0, arg1)
}

// CreateYamlFile mocks base method
func (m *MockYamlCreator) CreateYamlFile(arg0 files.FileCreator, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateYamlFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateYamlFile indicates an expected call of CreateYamlFile
func (mr *MockYamlCreatorMockRecorder) CreateYamlFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateYamlFile", reflect.TypeOf((*MockYamlCreator)(nil).CreateYamlFile), arg0, arg1, arg2)
}

// NewYamlConfig mocks base method
func (m *MockYamlCreator) NewYamlConfig() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewYamlConfig")
}

// NewYamlConfig indicates an expected call of NewYamlConfig
func (mr *MockYamlCreatorMockRecorder) NewYamlConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewYamlConfig", reflect.TypeOf((*MockYamlCreator)(nil).NewYamlConfig))
}
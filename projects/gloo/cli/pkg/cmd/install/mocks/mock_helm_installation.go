// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install (interfaces: HelmInstallation)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	chart "helm.sh/helm/v3/pkg/chart"
	release "helm.sh/helm/v3/pkg/release"
)

// MockHelmInstallation is a mock of HelmInstallation interface.
type MockHelmInstallation struct {
	ctrl     *gomock.Controller
	recorder *MockHelmInstallationMockRecorder
}

// MockHelmInstallationMockRecorder is the mock recorder for MockHelmInstallation.
type MockHelmInstallationMockRecorder struct {
	mock *MockHelmInstallation
}

// NewMockHelmInstallation creates a new mock instance.
func NewMockHelmInstallation(ctrl *gomock.Controller) *MockHelmInstallation {
	mock := &MockHelmInstallation{ctrl: ctrl}
	mock.recorder = &MockHelmInstallationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelmInstallation) EXPECT() *MockHelmInstallationMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockHelmInstallation) Run(arg0 *chart.Chart, arg1 map[string]interface{}) (*release.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(*release.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockHelmInstallationMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockHelmInstallation)(nil).Run), arg0, arg1)
}

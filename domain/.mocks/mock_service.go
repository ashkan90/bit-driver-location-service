// Code generated by MockGen. DO NOT EDIT.
// Source: adapters/handler/match_handler.go

// Package mocks is a generated GoMock package.
package mocks

import (
	request "bit-driver-location-service/request"
	response "bit-driver-location-service/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDriverImplementations is a mock of DriverImplementations interface.
type MockDriverImplementations struct {
	ctrl     *gomock.Controller
	recorder *MockDriverImplementationsMockRecorder
}

// MockDriverImplementationsMockRecorder is the mock recorder for MockDriverImplementations.
type MockDriverImplementationsMockRecorder struct {
	mock *MockDriverImplementations
}

// NewMockDriverImplementations creates a new mock instance.
func NewMockDriverImplementations(ctrl *gomock.Controller) *MockDriverImplementations {
	mock := &MockDriverImplementations{ctrl: ctrl}
	mock.recorder = &MockDriverImplementationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriverImplementations) EXPECT() *MockDriverImplementationsMockRecorder {
	return m.recorder
}

// FindNearestDriverLocation mocks base method.
func (m *MockDriverImplementations) FindNearestDriverLocation(loc request.CustomerLocation) response.DriverLocation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNearestDriverLocation", loc)
	ret0, _ := ret[0].(response.DriverLocation)
	return ret0
}

// FindNearestDriverLocation indicates an expected call of FindNearestDriverLocation.
func (mr *MockDriverImplementationsMockRecorder) FindNearestDriverLocation(loc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNearestDriverLocation", reflect.TypeOf((*MockDriverImplementations)(nil).FindNearestDriverLocation), loc)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: string_api.go

// Package mock_apis is a generated GoMock package.
package mock_apis

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStringAPI is a mock of StringAPI interface
type MockStringAPI struct {
	ctrl     *gomock.Controller
	recorder *MockStringAPIMockRecorder
}

// MockStringAPIMockRecorder is the mock recorder for MockStringAPI
type MockStringAPIMockRecorder struct {
	mock *MockStringAPI
}

// NewMockStringAPI creates a new mock instance
func NewMockStringAPI(ctrl *gomock.Controller) *MockStringAPI {
	mock := &MockStringAPI{ctrl: ctrl}
	mock.recorder = &MockStringAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStringAPI) EXPECT() *MockStringAPIMockRecorder {
	return m.recorder
}

// GetAString mocks base method
func (m *MockStringAPI) GetAString() string {
	ret := m.ctrl.Call(m, "GetAString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAString indicates an expected call of GetAString
func (mr *MockStringAPIMockRecorder) GetAString() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAString", reflect.TypeOf((*MockStringAPI)(nil).GetAString))
}

// SetAString mocks base method
func (m *MockStringAPI) SetAString(str string) {
	m.ctrl.Call(m, "SetAString", str)
}

// SetAString indicates an expected call of SetAString
func (mr *MockStringAPIMockRecorder) SetAString(str interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAString", reflect.TypeOf((*MockStringAPI)(nil).SetAString), str)
}

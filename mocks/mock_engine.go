// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/drone/autoscaler (interfaces: Engine)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Pause mocks base method
func (m *MockEngine) Pause() {
	m.ctrl.Call(m, "Pause")
}

// Pause indicates an expected call of Pause
func (mr *MockEngineMockRecorder) Pause() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockEngine)(nil).Pause))
}

// Paused mocks base method
func (m *MockEngine) Paused() bool {
	ret := m.ctrl.Call(m, "Paused")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Paused indicates an expected call of Paused
func (mr *MockEngineMockRecorder) Paused() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paused", reflect.TypeOf((*MockEngine)(nil).Paused))
}

// Resume mocks base method
func (m *MockEngine) Resume() {
	m.ctrl.Call(m, "Resume")
}

// Resume indicates an expected call of Resume
func (mr *MockEngineMockRecorder) Resume() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resume", reflect.TypeOf((*MockEngine)(nil).Resume))
}

// Start mocks base method
func (m *MockEngine) Start(arg0 context.Context) {
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start
func (mr *MockEngineMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEngine)(nil).Start), arg0)
}
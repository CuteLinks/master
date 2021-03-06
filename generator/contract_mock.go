// Code generated by MockGen. DO NOT EDIT.
// Source: contract1.go

// Package campaigns_forecasts_metrics is a generated GoMock package.
package generator

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockrandomizer is a mock of randomizer interface.
type Mockrandomizer struct {
	ctrl     *gomock.Controller
	recorder *MockrandomizerMockRecorder
}

// MockrandomizerMockRecorder is the mock recorder for Mockrandomizer.
type MockrandomizerMockRecorder struct {
	mock *Mockrandomizer
}

// NewMockrandomizer creates a new mock instance.
func NewMockrandomizer(ctrl *gomock.Controller) *Mockrandomizer {
	mock := &Mockrandomizer{ctrl: ctrl}
	mock.recorder = &MockrandomizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrandomizer) EXPECT() *MockrandomizerMockRecorder {
	return m.recorder
}

// GetRandomInt mocks base method.
func (m *Mockrandomizer) GetRandomInt(n int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomInt", n)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetRandomInt indicates an expected call of GetRandomInt.
func (mr *MockrandomizerMockRecorder) GetRandomInt(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomInt", reflect.TypeOf((*Mockrandomizer)(nil).GetRandomInt), n)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/http/interfaces.go

// Package http is a generated GoMock package.
package http

import (
	produce "github.com/davidlick/supermarket-api/internal/produce"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProduceService is a mock of ProduceService interface
type MockProduceService struct {
	ctrl     *gomock.Controller
	recorder *MockProduceServiceMockRecorder
}

// MockProduceServiceMockRecorder is the mock recorder for MockProduceService
type MockProduceServiceMockRecorder struct {
	mock *MockProduceService
}

// NewMockProduceService creates a new mock instance
func NewMockProduceService(ctrl *gomock.Controller) *MockProduceService {
	mock := &MockProduceService{ctrl: ctrl}
	mock.recorder = &MockProduceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProduceService) EXPECT() *MockProduceServiceMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockProduceService) Add(items []produce.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", items)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockProduceServiceMockRecorder) Add(items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockProduceService)(nil).Add), items)
}

// Remove mocks base method
func (m *MockProduceService) Remove(item produce.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockProduceServiceMockRecorder) Remove(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockProduceService)(nil).Remove), item)
}

// Get mocks base method
func (m *MockProduceService) Get(produceCode string) (produce.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", produceCode)
	ret0, _ := ret[0].(produce.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockProduceServiceMockRecorder) Get(produceCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProduceService)(nil).Get), produceCode)
}

// All mocks base method
func (m *MockProduceService) All() ([]produce.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]produce.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockProduceServiceMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockProduceService)(nil).All))
}

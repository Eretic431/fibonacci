// Code generated by MockGen. DO NOT EDIT.
// Source: fibonacci_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFibonacciRepository is a mock of FibonacciRepository interface.
type MockFibonacciRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFibonacciRepositoryMockRecorder
}

// MockFibonacciRepositoryMockRecorder is the mock recorder for MockFibonacciRepository.
type MockFibonacciRepositoryMockRecorder struct {
	mock *MockFibonacciRepository
}

// NewMockFibonacciRepository creates a new mock instance.
func NewMockFibonacciRepository(ctrl *gomock.Controller) *MockFibonacciRepository {
	mock := &MockFibonacciRepository{ctrl: ctrl}
	mock.recorder = &MockFibonacciRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFibonacciRepository) EXPECT() *MockFibonacciRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockFibonacciRepository) Get(ctx context.Context, key int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFibonacciRepositoryMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFibonacciRepository)(nil).Get), ctx, key)
}

// GetInterval mocks base method.
func (m *MockFibonacciRepository) GetInterval(ctx context.Context, from, to int) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterval", ctx, from, to)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInterval indicates an expected call of GetInterval.
func (mr *MockFibonacciRepositoryMockRecorder) GetInterval(ctx, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterval", reflect.TypeOf((*MockFibonacciRepository)(nil).GetInterval), ctx, from, to)
}

// GetLastTwoNumbers mocks base method.
func (m *MockFibonacciRepository) GetLastTwoNumbers(ctx context.Context) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastTwoNumbers", ctx)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastTwoNumbers indicates an expected call of GetLastTwoNumbers.
func (mr *MockFibonacciRepositoryMockRecorder) GetLastTwoNumbers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastTwoNumbers", reflect.TypeOf((*MockFibonacciRepository)(nil).GetLastTwoNumbers), ctx)
}

// Set mocks base method.
func (m *MockFibonacciRepository) Set(ctx context.Context, key int, value int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockFibonacciRepositoryMockRecorder) Set(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockFibonacciRepository)(nil).Set), ctx, key, value)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/database/query/main.go

// Package query is a generated GoMock package.
package query

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQueryService is a mock of QueryService interface.
type MockQueryService[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockQueryServiceMockRecorder[T]
}

// MockQueryServiceMockRecorder is the mock recorder for MockQueryService.
type MockQueryServiceMockRecorder[T any] struct {
	mock *MockQueryService[T]
}

// NewMockQueryService creates a new mock instance.
func NewMockQueryService[T any](ctrl *gomock.Controller) *MockQueryService[T] {
	mock := &MockQueryService[T]{ctrl: ctrl}
	mock.recorder = &MockQueryServiceMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryService[T]) EXPECT() *MockQueryServiceMockRecorder[T] {
	return m.recorder
}

// Insert mocks base method.
func (m *MockQueryService[T]) Insert(query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockQueryServiceMockRecorder[T]) Insert(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockQueryService[T])(nil).Insert), varargs...)
}

// Query mocks base method.
func (m *MockQueryService[T]) Query(query string, args ...interface{}) ([]T, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockQueryServiceMockRecorder[T]) Query(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockQueryService[T])(nil).Query), varargs...)
}

// QueryOne mocks base method.
func (m *MockQueryService[T]) QueryOne(query string, args ...interface{}) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryOne", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOne indicates an expected call of QueryOne.
func (mr *MockQueryServiceMockRecorder[T]) QueryOne(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOne", reflect.TypeOf((*MockQueryService[T])(nil).QueryOne), varargs...)
}

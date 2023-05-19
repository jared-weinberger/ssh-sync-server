// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/database/query/transaction.go

// Package query is a generated GoMock package.
package query

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgx "github.com/jackc/pgx/v5"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTransactionService) Commit(tx pgx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionServiceMockRecorder) Commit(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransactionService)(nil).Commit), tx)
}

// Rollback mocks base method.
func (m *MockTransactionService) Rollback(tx pgx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionServiceMockRecorder) Rollback(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransactionService)(nil).Rollback), tx)
}

// StartTx mocks base method.
func (m *MockTransactionService) StartTx(options pgx.TxOptions) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTx", options)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartTx indicates an expected call of StartTx.
func (mr *MockTransactionServiceMockRecorder) StartTx(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTx", reflect.TypeOf((*MockTransactionService)(nil).StartTx), options)
}

// MockQueryServiceTx is a mock of QueryServiceTx interface.
type MockQueryServiceTx[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockQueryServiceTxMockRecorder[T]
}

// MockQueryServiceTxMockRecorder is the mock recorder for MockQueryServiceTx.
type MockQueryServiceTxMockRecorder[T any] struct {
	mock *MockQueryServiceTx[T]
}

// NewMockQueryServiceTx creates a new mock instance.
func NewMockQueryServiceTx[T any](ctrl *gomock.Controller) *MockQueryServiceTx[T] {
	mock := &MockQueryServiceTx[T]{ctrl: ctrl}
	mock.recorder = &MockQueryServiceTxMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryServiceTx[T]) EXPECT() *MockQueryServiceTxMockRecorder[T] {
	return m.recorder
}

// Insert mocks base method.
func (m *MockQueryServiceTx[T]) Insert(tx pgx.Tx, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockQueryServiceTxMockRecorder[T]) Insert(tx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockQueryServiceTx[T])(nil).Insert), varargs...)
}

// Query mocks base method.
func (m *MockQueryServiceTx[T]) Query(tx pgx.Tx, query string, args ...interface{}) ([]T, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockQueryServiceTxMockRecorder[T]) Query(tx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockQueryServiceTx[T])(nil).Query), varargs...)
}

// QueryOne mocks base method.
func (m *MockQueryServiceTx[T]) QueryOne(tx pgx.Tx, query string, args ...interface{}) (*T, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryOne", varargs...)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryOne indicates an expected call of QueryOne.
func (mr *MockQueryServiceTxMockRecorder[T]) QueryOne(tx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryOne", reflect.TypeOf((*MockQueryServiceTx[T])(nil).QueryOne), varargs...)
}

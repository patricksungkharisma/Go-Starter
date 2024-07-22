// Code generated by MockGen. DO NOT EDIT.
// Source: dependency.go

// Package app is a generated GoMock package.
package app

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockdatabaseResource is a mock of databaseResource interface.
type MockdatabaseResource struct {
	ctrl     *gomock.Controller
	recorder *MockdatabaseResourceMockRecorder
}

// MockdatabaseResourceMockRecorder is the mock recorder for MockdatabaseResource.
type MockdatabaseResourceMockRecorder struct {
	mock *MockdatabaseResource
}

// NewMockdatabaseResource creates a new mock instance.
func NewMockdatabaseResource(ctrl *gomock.Controller) *MockdatabaseResource {
	mock := &MockdatabaseResource{ctrl: ctrl}
	mock.recorder = &MockdatabaseResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdatabaseResource) EXPECT() *MockdatabaseResourceMockRecorder {
	return m.recorder
}

// BeginTxx mocks base method.
func (m *MockdatabaseResource) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTxx", ctx, opts)
	ret0, _ := ret[0].(*sqlx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTxx indicates an expected call of BeginTxx.
func (mr *MockdatabaseResourceMockRecorder) BeginTxx(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTxx", reflect.TypeOf((*MockdatabaseResource)(nil).BeginTxx), ctx, opts)
}

// ExecContext mocks base method.
func (m *MockdatabaseResource) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockdatabaseResourceMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockdatabaseResource)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockdatabaseResource) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockdatabaseResourceMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockdatabaseResource)(nil).GetContext), varargs...)
}

// QueryxContext mocks base method.
func (m *MockdatabaseResource) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryxContext", varargs...)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryxContext indicates an expected call of QueryxContext.
func (mr *MockdatabaseResourceMockRecorder) QueryxContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryxContext", reflect.TypeOf((*MockdatabaseResource)(nil).QueryxContext), varargs...)
}

// SelectContext mocks base method.
func (m *MockdatabaseResource) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockdatabaseResourceMockRecorder) SelectContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockdatabaseResource)(nil).SelectContext), varargs...)
}
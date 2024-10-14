// Code generated by MockGen. DO NOT EDIT.
// Source: dependency.go

// Package app_test is a generated GoMock package.
package app_test

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockDatabaseResource is a mock of DatabaseResource interface.
type MockDatabaseResource struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseResourceMockRecorder
}

// MockDatabaseResourceMockRecorder is the mock recorder for MockDatabaseResource.
type MockDatabaseResourceMockRecorder struct {
	mock *MockDatabaseResource
}

// NewMockDatabaseResource creates a new mock instance.
func NewMockDatabaseResource(ctrl *gomock.Controller) *MockDatabaseResource {
	mock := &MockDatabaseResource{ctrl: ctrl}
	mock.recorder = &MockDatabaseResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseResource) EXPECT() *MockDatabaseResourceMockRecorder {
	return m.recorder
}

// BeginTxx mocks base method.
func (m *MockDatabaseResource) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTxx", ctx, opts)
	ret0, _ := ret[0].(*sqlx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTxx indicates an expected call of BeginTxx.
func (mr *MockDatabaseResourceMockRecorder) BeginTxx(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTxx", reflect.TypeOf((*MockDatabaseResource)(nil).BeginTxx), ctx, opts)
}

// ExecContext mocks base method.
func (m *MockDatabaseResource) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
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
func (mr *MockDatabaseResourceMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockDatabaseResource)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockDatabaseResource) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
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
func (mr *MockDatabaseResourceMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockDatabaseResource)(nil).GetContext), varargs...)
}

// QueryxContext mocks base method.
func (m *MockDatabaseResource) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
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
func (mr *MockDatabaseResourceMockRecorder) QueryxContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryxContext", reflect.TypeOf((*MockDatabaseResource)(nil).QueryxContext), varargs...)
}

// SelectContext mocks base method.
func (m *MockDatabaseResource) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
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
func (mr *MockDatabaseResourceMockRecorder) SelectContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockDatabaseResource)(nil).SelectContext), varargs...)
}

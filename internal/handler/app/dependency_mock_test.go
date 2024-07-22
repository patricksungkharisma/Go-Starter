// Code generated by MockGen. DO NOT EDIT.
// Source: dependency.go

// Package app is a generated GoMock package.
package app

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/patricksungkharisma/go-starter/internal/entity"
)

// MockappUsecase is a mock of appUsecase interface.
type MockappUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockappUsecaseMockRecorder
}

// MockappUsecaseMockRecorder is the mock recorder for MockappUsecase.
type MockappUsecaseMockRecorder struct {
	mock *MockappUsecase
}

// NewMockappUsecase creates a new mock instance.
func NewMockappUsecase(ctrl *gomock.Controller) *MockappUsecase {
	mock := &MockappUsecase{ctrl: ctrl}
	mock.recorder = &MockappUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockappUsecase) EXPECT() *MockappUsecaseMockRecorder {
	return m.recorder
}

// GetExampleData mocks base method.
func (m *MockappUsecase) GetExampleData(ctx context.Context, req entity.GetExampleRequest) (entity.GetExampleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExampleData", ctx, req)
	ret0, _ := ret[0].(entity.GetExampleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExampleData indicates an expected call of GetExampleData.
func (mr *MockappUsecaseMockRecorder) GetExampleData(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExampleData", reflect.TypeOf((*MockappUsecase)(nil).GetExampleData), ctx, req)
}
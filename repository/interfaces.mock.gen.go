// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=repository/interfaces.go -destination=repository/interfaces.mock.gen.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

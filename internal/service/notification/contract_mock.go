// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package notification is a generated GoMock package.
package notification

import (
	context "context"
	reflect "reflect"

	gophkeeper "github.com/egor-zakharov/goph-keeper/pkg/proto/gophkeeper"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockService) Add(ctx context.Context, stream gophkeeper.GophKeeper_SubscribeToChangesServer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", ctx, stream)
}

// Add indicates an expected call of Add.
func (mr *MockServiceMockRecorder) Add(ctx, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockService)(nil).Add), ctx, stream)
}

// Send mocks base method.
func (m *MockService) Send(ctx context.Context, product, action, id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", ctx, product, action, id)
}

// Send indicates an expected call of Send.
func (mr *MockServiceMockRecorder) Send(ctx, product, action, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockService)(nil).Send), ctx, product, action, id)
}

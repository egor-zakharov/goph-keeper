// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package textdata is a generated GoMock package.
package textdata

import (
	context "context"
	reflect "reflect"

	models "github.com/egor-zakharov/goph-keeper/internal/models"
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

// Create mocks base method.
func (m *MockService) Create(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, textData, userID)
	ret0, _ := ret[0].(*models.TextData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(ctx, textData, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, textData, userID)
}

// Delete mocks base method.
func (m *MockService) Delete(ctx context.Context, id, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), ctx, id, userID)
}

// Read mocks base method.
func (m *MockService) Read(ctx context.Context, userID string) (*[]models.TextData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, userID)
	ret0, _ := ret[0].(*[]models.TextData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockServiceMockRecorder) Read(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockService)(nil).Read), ctx, userID)
}

// Update mocks base method.
func (m *MockService) Update(ctx context.Context, textData models.TextData, userID string) (*models.TextData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, textData, userID)
	ret0, _ := ret[0].(*models.TextData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(ctx, textData, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), ctx, textData, userID)
}

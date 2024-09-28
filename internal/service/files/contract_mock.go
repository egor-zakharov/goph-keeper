// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package files is a generated GoMock package.
package files

import (
	context "context"
	reflect "reflect"

	models "github.com/egor-zakharov/goph-keeper/internal/models"
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
func (m *MockService) Add(ctx context.Context, stream gophkeeper.GophKeeper_UploadFileServer) (*models.FileData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, stream)
	ret0, _ := ret[0].(*models.FileData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockServiceMockRecorder) Add(ctx, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockService)(nil).Add), ctx, stream)
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

// Download mocks base method.
func (m *MockService) Download(in *gophkeeper.DownloadFileRequest, stream gophkeeper.GophKeeper_DownloadFileServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", in, stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download.
func (mr *MockServiceMockRecorder) Download(in, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockService)(nil).Download), in, stream)
}

// Read mocks base method.
func (m *MockService) Read(ctx context.Context, userID string) (*[]models.FileData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, userID)
	ret0, _ := ret[0].(*[]models.FileData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockServiceMockRecorder) Read(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockService)(nil).Read), ctx, userID)
}

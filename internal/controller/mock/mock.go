// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/controller/controller.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "net-api.com/internal/domain/entities"
)

// MockBookServiceInterface is a mock of BookServiceInterface interface.
type MockBookServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBookServiceInterfaceMockRecorder
}

// MockBookServiceInterfaceMockRecorder is the mock recorder for MockBookServiceInterface.
type MockBookServiceInterfaceMockRecorder struct {
	mock *MockBookServiceInterface
}

// NewMockBookServiceInterface creates a new mock instance.
func NewMockBookServiceInterface(ctrl *gomock.Controller) *MockBookServiceInterface {
	mock := &MockBookServiceInterface{ctrl: ctrl}
	mock.recorder = &MockBookServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookServiceInterface) EXPECT() *MockBookServiceInterfaceMockRecorder {
	return m.recorder
}

// GetBookByID mocks base method.
func (m *MockBookServiceInterface) GetBookByID(hash, bookIdParam string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", hash, bookIdParam)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockBookServiceInterfaceMockRecorder) GetBookByID(hash, bookIdParam interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockBookServiceInterface)(nil).GetBookByID), hash, bookIdParam)
}

// GetBookRandom mocks base method.
func (m *MockBookServiceInterface) GetBookRandom(hash string, w http.ResponseWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookRandom", hash, w)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetBookRandom indicates an expected call of GetBookRandom.
func (mr *MockBookServiceInterfaceMockRecorder) GetBookRandom(hash, w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookRandom", reflect.TypeOf((*MockBookServiceInterface)(nil).GetBookRandom), hash, w)
}

// GetHash mocks base method.
func (m *MockBookServiceInterface) GetHash(user *entities.User) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHash", user)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHash indicates an expected call of GetHash.
func (mr *MockBookServiceInterfaceMockRecorder) GetHash(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHash", reflect.TypeOf((*MockBookServiceInterface)(nil).GetHash), user)
}

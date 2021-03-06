// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/IfuryI/WEB_BACK/internal/ratings (interfaces: UseCase)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/IfuryI/WEB_BACK/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CreateRating mocks base method.
func (m *MockUseCase) CreateRating(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRating", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRating indicates an expected call of CreateRating.
func (mr *MockUseCaseMockRecorder) CreateRating(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRating", reflect.TypeOf((*MockUseCase)(nil).CreateRating), arg0, arg1, arg2)
}

// DeleteRating mocks base method.
func (m *MockUseCase) DeleteRating(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRating", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRating indicates an expected call of DeleteRating.
func (mr *MockUseCaseMockRecorder) DeleteRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRating", reflect.TypeOf((*MockUseCase)(nil).DeleteRating), arg0, arg1)
}

// GetRating mocks base method.
func (m *MockUseCase) GetRating(arg0, arg1 string) (models.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRating", arg0, arg1)
	ret0, _ := ret[0].(models.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRating indicates an expected call of GetRating.
func (mr *MockUseCaseMockRecorder) GetRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRating", reflect.TypeOf((*MockUseCase)(nil).GetRating), arg0, arg1)
}

// UpdateRating mocks base method.
func (m *MockUseCase) UpdateRating(arg0, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRating", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRating indicates an expected call of UpdateRating.
func (mr *MockUseCaseMockRecorder) UpdateRating(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRating", reflect.TypeOf((*MockUseCase)(nil).UpdateRating), arg0, arg1, arg2)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/IfuryI/WEB_BACK/internal/actors (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/IfuryI/WEB_BACK/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetActorByID mocks base method.
func (m *MockRepository) GetActorByID(arg0, arg1 string) (models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorByID", arg0, arg1)
	ret0, _ := ret[0].(models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorByID indicates an expected call of GetActorByID.
func (mr *MockRepositoryMockRecorder) GetActorByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorByID", reflect.TypeOf((*MockRepository)(nil).GetActorByID), arg0, arg1)
}

// GetFavoriteActors mocks base method.
func (m *MockRepository) GetFavoriteActors(arg0 string) ([]models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoriteActors", arg0)
	ret0, _ := ret[0].([]models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoriteActors indicates an expected call of GetFavoriteActors.
func (mr *MockRepositoryMockRecorder) GetFavoriteActors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoriteActors", reflect.TypeOf((*MockRepository)(nil).GetFavoriteActors), arg0)
}

// LikeActor mocks base method.
func (m *MockRepository) LikeActor(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeActor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LikeActor indicates an expected call of LikeActor.
func (mr *MockRepositoryMockRecorder) LikeActor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeActor", reflect.TypeOf((*MockRepository)(nil).LikeActor), arg0, arg1)
}

// SearchActors mocks base method.
func (m *MockRepository) SearchActors(arg0 string) ([]models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchActors", arg0)
	ret0, _ := ret[0].([]models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchActors indicates an expected call of SearchActors.
func (mr *MockRepositoryMockRecorder) SearchActors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchActors", reflect.TypeOf((*MockRepository)(nil).SearchActors), arg0)
}

// UnlikeActor mocks base method.
func (m *MockRepository) UnlikeActor(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlikeActor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlikeActor indicates an expected call of UnlikeActor.
func (mr *MockRepositoryMockRecorder) UnlikeActor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlikeActor", reflect.TypeOf((*MockRepository)(nil).UnlikeActor), arg0, arg1)
}

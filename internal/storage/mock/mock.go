// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/rrbit/GolandProjects/simple-microservice/internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	etherium "gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	gomock "github.com/golang/mock/gomock"
)

// MockEtheriumRepository is a mock of EtheriumRepository interface.
type MockEtheriumRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEtheriumRepositoryMockRecorder
}

// MockEtheriumRepositoryMockRecorder is the mock recorder for MockEtheriumRepository.
type MockEtheriumRepositoryMockRecorder struct {
	mock *MockEtheriumRepository
}

// NewMockEtheriumRepository creates a new mock instance.
func NewMockEtheriumRepository(ctrl *gomock.Controller) *MockEtheriumRepository {
	mock := &MockEtheriumRepository{ctrl: ctrl}
	mock.recorder = &MockEtheriumRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEtheriumRepository) EXPECT() *MockEtheriumRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEtheriumRepository) Create(eth etherium.Etherium) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", eth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockEtheriumRepositoryMockRecorder) Create(eth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEtheriumRepository)(nil).Create), eth)
}

// Delete mocks base method.
func (m *MockEtheriumRepository) Delete(eth etherium.Etherium) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", eth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEtheriumRepositoryMockRecorder) Delete(eth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEtheriumRepository)(nil).Delete), eth)
}

// Read mocks base method.
func (m *MockEtheriumRepository) Read(eth etherium.Etherium) (etherium.Etherium, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", eth)
	ret0, _ := ret[0].(etherium.Etherium)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockEtheriumRepositoryMockRecorder) Read(eth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockEtheriumRepository)(nil).Read), eth)
}

// Update mocks base method.
func (m *MockEtheriumRepository) Update(eth etherium.Etherium) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", eth)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEtheriumRepositoryMockRecorder) Update(eth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEtheriumRepository)(nil).Update), eth)
}

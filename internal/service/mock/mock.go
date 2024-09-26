// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/rrbit/GolandProjects/simple-microservice/internal/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	etherium "gihub.com/robbitancor/simple-microservice/internal/domain/crypto/etherium"
	storage "gihub.com/robbitancor/simple-microservice/internal/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockEtheriumService is a mock of EtheriumService interface.
type MockEtheriumService struct {
	ctrl     *gomock.Controller
	recorder *MockEtheriumServiceMockRecorder
}

// MockEtheriumServiceMockRecorder is the mock recorder for MockEtheriumService.
type MockEtheriumServiceMockRecorder struct {
	mock *MockEtheriumService
}

// NewMockEtheriumService creates a new mock instance.
func NewMockEtheriumService(ctrl *gomock.Controller) *MockEtheriumService {
	mock := &MockEtheriumService{ctrl: ctrl}
	mock.recorder = &MockEtheriumServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEtheriumService) EXPECT() *MockEtheriumServiceMockRecorder {
	return m.recorder
}

// GetBalance mocks base method.
func (m *MockEtheriumService) GetBalance(uri, address, module, action, apiKey string) etherium.Balance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", uri, address, module, action, apiKey)
	ret0, _ := ret[0].(etherium.Balance)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockEtheriumServiceMockRecorder) GetBalance(uri, address, module, action, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockEtheriumService)(nil).GetBalance), uri, address, module, action, apiKey)
}

// GetBlockNumber mocks base method.
func (m *MockEtheriumService) GetBlockNumber(uri, module, action, apiKey string) etherium.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockNumber", uri, module, action, apiKey)
	ret0, _ := ret[0].(etherium.Block)
	return ret0
}

// GetBlockNumber indicates an expected call of GetBlockNumber.
func (mr *MockEtheriumServiceMockRecorder) GetBlockNumber(uri, module, action, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockNumber", reflect.TypeOf((*MockEtheriumService)(nil).GetBlockNumber), uri, module, action, apiKey)
}

// GetGasPrice mocks base method.
func (m *MockEtheriumService) GetGasPrice(uri, module, action, apiKey string) etherium.Gas {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGasPrice", uri, module, action, apiKey)
	ret0, _ := ret[0].(etherium.Gas)
	return ret0
}

// GetGasPrice indicates an expected call of GetGasPrice.
func (mr *MockEtheriumServiceMockRecorder) GetGasPrice(uri, module, action, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGasPrice", reflect.TypeOf((*MockEtheriumService)(nil).GetGasPrice), uri, module, action, apiKey)
}

// SaveEtherium mocks base method.
func (m *MockEtheriumService) SaveEtherium(eth etherium.Etherium) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEtherium", eth)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveEtherium indicates an expected call of SaveEtherium.
func (mr *MockEtheriumServiceMockRecorder) SaveEtherium(eth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEtherium", reflect.TypeOf((*MockEtheriumService)(nil).SaveEtherium), eth)
}

// SetRepo mocks base method.
func (m *MockEtheriumService) SetRepo(repo storage.EtheriumRepository) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRepo", repo)
}

// SetRepo indicates an expected call of SetRepo.
func (mr *MockEtheriumServiceMockRecorder) SetRepo(repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRepo", reflect.TypeOf((*MockEtheriumService)(nil).SetRepo), repo)
}
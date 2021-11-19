// Code generated by MockGen. DO NOT EDIT.
// Source: internal/record/repository/recorder_repository.go

// Package mockrecord is a generated GoMock package.
package mockrecord

import (
	reflect "reflect"

	entity "github.com/enchainte/enchainte-sdk-go/internal/record/entity"
	dto "github.com/enchainte/enchainte-sdk-go/internal/record/entity/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockRecorderRepository is a mock of RecorderRepository interface.
type MockRecorderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRecorderRepositoryMockRecorder
}

// MockRecorderRepositoryMockRecorder is the mock recorder for MockRecorderRepository.
type MockRecorderRepositoryMockRecorder struct {
	mock *MockRecorderRepository
}

// NewMockRecorderRepository creates a new mock instance.
func NewMockRecorderRepository(ctrl *gomock.Controller) *MockRecorderRepository {
	mock := &MockRecorderRepository{ctrl: ctrl}
	mock.recorder = &MockRecorderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecorderRepository) EXPECT() *MockRecorderRepositoryMockRecorder {
	return m.recorder
}

// FetchRecords mocks base method.
func (m *MockRecorderRepository) FetchRecords(records []entity.RecordEntity) ([]dto.RecordRetrieveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRecords", records)
	ret0, _ := ret[0].([]dto.RecordRetrieveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRecords indicates an expected call of FetchRecords.
func (mr *MockRecorderRepositoryMockRecorder) FetchRecords(records interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRecords", reflect.TypeOf((*MockRecorderRepository)(nil).FetchRecords), records)
}

// SendRecords mocks base method.
func (m *MockRecorderRepository) SendRecords(records []entity.RecordEntity) (dto.RecordWriteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRecords", records)
	ret0, _ := ret[0].(dto.RecordWriteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRecords indicates an expected call of SendRecords.
func (mr *MockRecorderRepositoryMockRecorder) SendRecords(records interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRecords", reflect.TypeOf((*MockRecorderRepository)(nil).SendRecords), records)
}
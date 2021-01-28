// Code generated by MockGen. DO NOT EDIT.
// Source: internal/message/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	message "github.com/enchainte/enchainte-sdk-go/internal/message"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MessageService is a mock of Service interface
type MessageService struct {
	ctrl     *gomock.Controller
	recorder *MessageServiceMockRecorder
}

// MessageServiceMockRecorder is the mock recorder for MessageService
type MessageServiceMockRecorder struct {
	mock *MessageService
}

// NewMessageService creates a new mock instance
func NewMessageService(ctrl *gomock.Controller) *MessageService {
	mock := &MessageService{ctrl: ctrl}
	mock.recorder = &MessageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MessageService) EXPECT() *MessageServiceMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MessageService) Send(hash []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MessageServiceMockRecorder) Send(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MessageService)(nil).Send), hash)
}

// Search mocks base method
func (m *MessageService) Search(hash [][]byte) (*message.Receipts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", hash)
	ret0, _ := ret[0].(*message.Receipts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MessageServiceMockRecorder) Search(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MessageService)(nil).Search), hash)
}

// Verify mocks base method
func (m *MessageService) Verify(hashes [][]byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", hashes)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MessageServiceMockRecorder) Verify(hashes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MessageService)(nil).Verify), hashes)
}

// Wait mocks base method
func (m *MessageService) Wait(hashes [][]byte) (*message.Receipts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", hashes)
	ret0, _ := ret[0].(*message.Receipts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wait indicates an expected call of Wait
func (mr *MessageServiceMockRecorder) Wait(hashes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MessageService)(nil).Wait), hashes)
}

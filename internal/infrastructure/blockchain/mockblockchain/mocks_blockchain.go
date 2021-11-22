// Code generated by MockGen. DO NOT EDIT.
// Source: internal/infrastructure/blockchain_client.go

// Package mockblockchain is a generated GoMock package.
package mockblockchain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBlockchainClient is a mock of BlockchainClient interface.
type MockBlockchainClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainClientMockRecorder
}

// MockBlockchainClientMockRecorder is the mock recorder for MockBlockchainClient.
type MockBlockchainClientMockRecorder struct {
	mock *MockBlockchainClient
}

// NewMockBlockchainClient creates a new mock instance.
func NewMockBlockchainClient(ctrl *gomock.Controller) *MockBlockchainClient {
	mock := &MockBlockchainClient{ctrl: ctrl}
	mock.recorder = &MockBlockchainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockchainClient) EXPECT() *MockBlockchainClientMockRecorder {
	return m.recorder
}

// ValidateRoot mocks base method.
func (m *MockBlockchainClient) ValidateRoot(network, root string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRoot", network, root)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateRoot indicates an expected call of ValidateRoot.
func (mr *MockBlockchainClientMockRecorder) ValidateRoot(network, root interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRoot", reflect.TypeOf((*MockBlockchainClient)(nil).ValidateRoot), network, root)
}

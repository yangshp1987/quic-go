// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lucas-clemente/quic-go (interfaces: SealingManager)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	handshake "github.com/lucas-clemente/quic-go/internal/handshake"
)

// MockSealingManager is a mock of SealingManager interface
type MockSealingManager struct {
	ctrl     *gomock.Controller
	recorder *MockSealingManagerMockRecorder
}

// MockSealingManagerMockRecorder is the mock recorder for MockSealingManager
type MockSealingManagerMockRecorder struct {
	mock *MockSealingManager
}

// NewMockSealingManager creates a new mock instance
func NewMockSealingManager(ctrl *gomock.Controller) *MockSealingManager {
	mock := &MockSealingManager{ctrl: ctrl}
	mock.recorder = &MockSealingManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSealingManager) EXPECT() *MockSealingManagerMockRecorder {
	return m.recorder
}

// Get0RTTSealer mocks base method
func (m *MockSealingManager) Get0RTTSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTSealer indicates an expected call of Get0RTTSealer
func (mr *MockSealingManagerMockRecorder) Get0RTTSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTSealer", reflect.TypeOf((*MockSealingManager)(nil).Get0RTTSealer))
}

// Get1RTTSealer mocks base method
func (m *MockSealingManager) Get1RTTSealer() (handshake.ShortHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTSealer")
	ret0, _ := ret[0].(handshake.ShortHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTSealer indicates an expected call of Get1RTTSealer
func (mr *MockSealingManagerMockRecorder) Get1RTTSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTSealer", reflect.TypeOf((*MockSealingManager)(nil).Get1RTTSealer))
}

// GetHandshakeSealer mocks base method
func (m *MockSealingManager) GetHandshakeSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeSealer indicates an expected call of GetHandshakeSealer
func (mr *MockSealingManagerMockRecorder) GetHandshakeSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeSealer", reflect.TypeOf((*MockSealingManager)(nil).GetHandshakeSealer))
}

// GetInitialSealer mocks base method
func (m *MockSealingManager) GetInitialSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialSealer indicates an expected call of GetInitialSealer
func (mr *MockSealingManagerMockRecorder) GetInitialSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialSealer", reflect.TypeOf((*MockSealingManager)(nil).GetInitialSealer))
}

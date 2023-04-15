// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	v5 "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gomock "github.com/golang/mock/gomock"
	bruteforce "github.com/shlima/fortune/internal/pkg/bruteforce"
	key "github.com/shlima/fortune/internal/pkg/key"
)

// MockTelegramCli is a mock of ICli interface.
type MockTelegramCli struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramCliMockRecorder
}

// MockTelegramCliMockRecorder is the mock recorder for MockTelegramCli.
type MockTelegramCliMockRecorder struct {
	mock *MockTelegramCli
}

// NewMockTelegramCli creates a new mock instance.
func NewMockTelegramCli(ctrl *gomock.Controller) *MockTelegramCli {
	mock := &MockTelegramCli{ctrl: ctrl}
	mock.recorder = &MockTelegramCliMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramCli) EXPECT() *MockTelegramCliMockRecorder {
	return m.recorder
}

// HeartBeat mocks base method.
func (m *MockTelegramCli) HeartBeat(heartbit *bruteforce.HeartBit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeartBeat", heartbit)
	ret0, _ := ret[0].(error)
	return ret0
}

// HeartBeat indicates an expected call of HeartBeat.
func (mr *MockTelegramCliMockRecorder) HeartBeat(heartbit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeartBeat", reflect.TypeOf((*MockTelegramCli)(nil).HeartBeat), heartbit)
}

// KeyFound mocks base method.
func (m *MockTelegramCli) KeyFound(chain key.Chain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyFound", chain)
	ret0, _ := ret[0].(error)
	return ret0
}

// KeyFound indicates an expected call of KeyFound.
func (mr *MockTelegramCliMockRecorder) KeyFound(chain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyFound", reflect.TypeOf((*MockTelegramCli)(nil).KeyFound), chain)
}

// MockTelegramApi is a mock of IApi interface.
type MockTelegramApi struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramApiMockRecorder
}

// MockTelegramApiMockRecorder is the mock recorder for MockTelegramApi.
type MockTelegramApiMockRecorder struct {
	mock *MockTelegramApi
}

// NewMockTelegramApi creates a new mock instance.
func NewMockTelegramApi(ctrl *gomock.Controller) *MockTelegramApi {
	mock := &MockTelegramApi{ctrl: ctrl}
	mock.recorder = &MockTelegramApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramApi) EXPECT() *MockTelegramApiMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockTelegramApi) Send(c v5.Chattable) (v5.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", c)
	ret0, _ := ret[0].(v5.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockTelegramApiMockRecorder) Send(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTelegramApi)(nil).Send), c)
}
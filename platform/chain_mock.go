// Code generated by MockGen. DO NOT EDIT.
// Source: ./platform/chain.go

// Package platform is a generated GoMock package.
package platform

import (
	reflect "reflect"

	domain "github.com/ahstn/defair/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockChain is a mock of Chain interface.
type MockChain struct {
	ctrl     *gomock.Controller
	recorder *MockChainMockRecorder
}

// MockChainMockRecorder is the mock recorder for MockChain.
type MockChainMockRecorder struct {
	mock *MockChain
}

// NewMockChain creates a new mock instance.
func NewMockChain(ctrl *gomock.Controller) *MockChain {
	mock := &MockChain{ctrl: ctrl}
	mock.recorder = &MockChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChain) EXPECT() *MockChainMockRecorder {
	return m.recorder
}

// Balances mocks base method.
func (m *MockChain) Balances(arg0, arg1 string, arg2 []domain.Token) ([]domain.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balances", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domain.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balances indicates an expected call of Balances.
func (mr *MockChainMockRecorder) Balances(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balances", reflect.TypeOf((*MockChain)(nil).Balances), arg0, arg1, arg2)
}

// LiquidityPools mocks base method.
func (m *MockChain) LiquidityPools(arg0 string, arg1 domain.Network) ([]domain.LiquidityPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LiquidityPools", arg0, arg1)
	ret0, _ := ret[0].([]domain.LiquidityPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LiquidityPools indicates an expected call of LiquidityPools.
func (mr *MockChainMockRecorder) LiquidityPools(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LiquidityPools", reflect.TypeOf((*MockChain)(nil).LiquidityPools), arg0, arg1)
}

// TokenInfo mocks base method.
func (m *MockChain) TokenInfo(arg0, arg1, arg2 string) (domain.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenInfo indicates an expected call of TokenInfo.
func (mr *MockChainMockRecorder) TokenInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenInfo", reflect.TypeOf((*MockChain)(nil).TokenInfo), arg0, arg1, arg2)
}

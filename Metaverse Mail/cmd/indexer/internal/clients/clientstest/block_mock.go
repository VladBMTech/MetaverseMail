// Code generated by MockGen. DO NOT EDIT.
// Source: block.go

// Package clientstest is a generated GoMock package.
package clientstest

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBlock is a mock of Block interface.
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock.
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance.
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// BlockByNumber mocks base method.
func (m *MockBlock) BlockByNumber(ctx context.Context, blockNo uint64) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, blockNo)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockBlockMockRecorder) BlockByNumber(ctx, blockNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockBlock)(nil).BlockByNumber), ctx, blockNo)
}

// LatestBlockNumber mocks base method.
func (m *MockBlock) LatestBlockNumber(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBlockNumber", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBlockNumber indicates an expected call of LatestBlockNumber.
func (mr *MockBlockMockRecorder) LatestBlockNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBlockNumber", reflect.TypeOf((*MockBlock)(nil).LatestBlockNumber), ctx)
}
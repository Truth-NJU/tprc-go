// Code generated by MockGen. DO NOT EDIT.
// Source: stub/github.com/Truth-NJU/tprc-go/trpcprotocol/online/online.trpc.go

// Package online is a generated GoMock package.
package online

import (
	context "context"
	reflect "reflect"

	client "git.code.oa.com/trpc-go/trpc-go/client"
	gomock "github.com/golang/mock/gomock"
)

// MockMessageService is a mock of MessageService interface.
type MockMessageService struct {
	ctrl     *gomock.Controller
	recorder *MockMessageServiceMockRecorder
}

// MockMessageServiceMockRecorder is the mock recorder for MockMessageService.
type MockMessageServiceMockRecorder struct {
	mock *MockMessageService
}

// NewMockMessageService creates a new mock instance.
func NewMockMessageService(ctrl *gomock.Controller) *MockMessageService {
	mock := &MockMessageService{ctrl: ctrl}
	mock.recorder = &MockMessageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageService) EXPECT() *MockMessageServiceMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockMessageService) SendMessage(ctx context.Context, req *MessageRequest) (*MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", ctx, req)
	ret0, _ := ret[0].(*MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockMessageServiceMockRecorder) SendMessage(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockMessageService)(nil).SendMessage), ctx, req)
}

// MockMessageClientProxy is a mock of MessageClientProxy interface.
type MockMessageClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockMessageClientProxyMockRecorder
}

// MockMessageClientProxyMockRecorder is the mock recorder for MockMessageClientProxy.
type MockMessageClientProxyMockRecorder struct {
	mock *MockMessageClientProxy
}

// NewMockMessageClientProxy creates a new mock instance.
func NewMockMessageClientProxy(ctrl *gomock.Controller) *MockMessageClientProxy {
	mock := &MockMessageClientProxy{ctrl: ctrl}
	mock.recorder = &MockMessageClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageClientProxy) EXPECT() *MockMessageClientProxyMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockMessageClientProxy) SendMessage(ctx context.Context, req *MessageRequest, opts ...client.Option) (*MessageResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessage", varargs...)
	ret0, _ := ret[0].(*MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockMessageClientProxyMockRecorder) SendMessage(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockMessageClientProxy)(nil).SendMessage), varargs...)
}

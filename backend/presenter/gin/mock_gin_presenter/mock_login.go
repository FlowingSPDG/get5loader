// Code generated by MockGen. DO NOT EDIT.
// Source: ./login.go

// Package mock_gin_presenter is a generated GoMock package.
package mock_gin_presenter

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockJWTPresenter is a mock of JWTPresenter interface.
type MockJWTPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockJWTPresenterMockRecorder
}

// MockJWTPresenterMockRecorder is the mock recorder for MockJWTPresenter.
type MockJWTPresenterMockRecorder struct {
	mock *MockJWTPresenter
}

// NewMockJWTPresenter creates a new mock instance.
func NewMockJWTPresenter(ctrl *gomock.Controller) *MockJWTPresenter {
	mock := &MockJWTPresenter{ctrl: ctrl}
	mock.recorder = &MockJWTPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTPresenter) EXPECT() *MockJWTPresenterMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockJWTPresenter) Handle(c *gin.Context, token string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", c, token)
}

// Handle indicates an expected call of Handle.
func (mr *MockJWTPresenterMockRecorder) Handle(c, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockJWTPresenter)(nil).Handle), c, token)
}
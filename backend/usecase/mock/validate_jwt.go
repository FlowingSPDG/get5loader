// Code generated by MockGen. DO NOT EDIT.
// Source: validate_jwt.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	entity "github.com/FlowingSPDG/get5-web-go/backend/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockValidateJWT is a mock of ValidateJWT interface.
type MockValidateJWT struct {
	ctrl     *gomock.Controller
	recorder *MockValidateJWTMockRecorder
}

// MockValidateJWTMockRecorder is the mock recorder for MockValidateJWT.
type MockValidateJWTMockRecorder struct {
	mock *MockValidateJWT
}

// NewMockValidateJWT creates a new mock instance.
func NewMockValidateJWT(ctrl *gomock.Controller) *MockValidateJWT {
	mock := &MockValidateJWT{ctrl: ctrl}
	mock.recorder = &MockValidateJWTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidateJWT) EXPECT() *MockValidateJWTMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidateJWT) Validate(token string) (*entity.TokenUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", token)
	ret0, _ := ret[0].(*entity.TokenUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockValidateJWTMockRecorder) Validate(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidateJWT)(nil).Validate), token)
}

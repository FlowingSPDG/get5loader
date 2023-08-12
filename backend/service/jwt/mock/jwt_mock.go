// Code generated by MockGen. DO NOT EDIT.
// Source: jwt.go

// Package mock_jwt is a generated GoMock package.
package mock_jwt

import (
	reflect "reflect"

	entity "github.com/FlowingSPDG/get5loader/backend/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockJWTService is a mock of JWTService interface.
type MockJWTService struct {
	ctrl     *gomock.Controller
	recorder *MockJWTServiceMockRecorder
}

// MockJWTServiceMockRecorder is the mock recorder for MockJWTService.
type MockJWTServiceMockRecorder struct {
	mock *MockJWTService
}

// NewMockJWTService creates a new mock instance.
func NewMockJWTService(ctrl *gomock.Controller) *MockJWTService {
	mock := &MockJWTService{ctrl: ctrl}
	mock.recorder = &MockJWTServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTService) EXPECT() *MockJWTServiceMockRecorder {
	return m.recorder
}

// IssueJWT mocks base method.
func (m *MockJWTService) IssueJWT(userID entity.UserID, steamID entity.SteamID, admin bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueJWT", userID, steamID, admin)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueJWT indicates an expected call of IssueJWT.
func (mr *MockJWTServiceMockRecorder) IssueJWT(userID, steamID, admin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueJWT", reflect.TypeOf((*MockJWTService)(nil).IssueJWT), userID, steamID, admin)
}

// ValidateJWT mocks base method.
func (m *MockJWTService) ValidateJWT(token string) (*entity.TokenUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateJWT", token)
	ret0, _ := ret[0].(*entity.TokenUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateJWT indicates an expected call of ValidateJWT.
func (mr *MockJWTServiceMockRecorder) ValidateJWT(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateJWT", reflect.TypeOf((*MockJWTService)(nil).ValidateJWT), token)
}

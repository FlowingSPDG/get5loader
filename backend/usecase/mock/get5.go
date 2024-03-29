// Code generated by MockGen. DO NOT EDIT.
// Source: get5.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/FlowingSPDG/get5loader/backend/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockGet5 is a mock of Get5 interface.
type MockGet5 struct {
	ctrl     *gomock.Controller
	recorder *MockGet5MockRecorder
}

// MockGet5MockRecorder is the mock recorder for MockGet5.
type MockGet5MockRecorder struct {
	mock *MockGet5
}

// NewMockGet5 creates a new mock instance.
func NewMockGet5(ctrl *gomock.Controller) *MockGet5 {
	mock := &MockGet5{ctrl: ctrl}
	mock.recorder = &MockGet5MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGet5) EXPECT() *MockGet5MockRecorder {
	return m.recorder
}

// GetMatch mocks base method.
func (m *MockGet5) GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Get5Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatch", ctx, matchID)
	ret0, _ := ret[0].(*entity.Get5Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatch indicates an expected call of GetMatch.
func (mr *MockGet5MockRecorder) GetMatch(ctx, matchID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatch", reflect.TypeOf((*MockGet5)(nil).GetMatch), ctx, matchID)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: get_match.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/FlowingSPDG/get5-web-go/backend/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockGetMatch is a mock of GetMatch interface.
type MockGetMatch struct {
	ctrl     *gomock.Controller
	recorder *MockGetMatchMockRecorder
}

// MockGetMatchMockRecorder is the mock recorder for MockGetMatch.
type MockGetMatchMockRecorder struct {
	mock *MockGetMatch
}

// NewMockGetMatch creates a new mock instance.
func NewMockGetMatch(ctrl *gomock.Controller) *MockGetMatch {
	mock := &MockGetMatch{ctrl: ctrl}
	mock.recorder = &MockGetMatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetMatch) EXPECT() *MockGetMatchMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetMatch) Get(ctx context.Context, matchID entity.MatchID) (*entity.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, matchID)
	ret0, _ := ret[0].(*entity.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetMatchMockRecorder) Get(ctx, matchID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetMatch)(nil).Get), ctx, matchID)
}

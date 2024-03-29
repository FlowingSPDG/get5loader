// Code generated by MockGen. DO NOT EDIT.
// Source: match.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/FlowingSPDG/get5loader/backend/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockMatch is a mock of Match interface.
type MockMatch struct {
	ctrl     *gomock.Controller
	recorder *MockMatchMockRecorder
}

// MockMatchMockRecorder is the mock recorder for MockMatch.
type MockMatchMockRecorder struct {
	mock *MockMatch
}

// NewMockMatch creates a new mock instance.
func NewMockMatch(ctrl *gomock.Controller) *MockMatch {
	mock := &MockMatch{ctrl: ctrl}
	mock.recorder = &MockMatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatch) EXPECT() *MockMatchMockRecorder {
	return m.recorder
}

// BatchGetMatchesByUser mocks base method.
func (m *MockMatch) BatchGetMatchesByUser(ctx context.Context, userIDs []entity.UserID) (map[entity.UserID][]*entity.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetMatchesByUser", ctx, userIDs)
	ret0, _ := ret[0].(map[entity.UserID][]*entity.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetMatchesByUser indicates an expected call of BatchGetMatchesByUser.
func (mr *MockMatchMockRecorder) BatchGetMatchesByUser(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetMatchesByUser", reflect.TypeOf((*MockMatch)(nil).BatchGetMatchesByUser), ctx, userIDs)
}

// CreateMatch mocks base method.
func (m *MockMatch) CreateMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID, team2ID entity.TeamID, maxMaps int, title string) (*entity.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMatch", ctx, userID, serverID, team1ID, team2ID, maxMaps, title)
	ret0, _ := ret[0].(*entity.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMatch indicates an expected call of CreateMatch.
func (mr *MockMatchMockRecorder) CreateMatch(ctx, userID, serverID, team1ID, team2ID, maxMaps, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMatch", reflect.TypeOf((*MockMatch)(nil).CreateMatch), ctx, userID, serverID, team1ID, team2ID, maxMaps, title)
}

// GetMatch mocks base method.
func (m *MockMatch) GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatch", ctx, matchID)
	ret0, _ := ret[0].(*entity.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatch indicates an expected call of GetMatch.
func (mr *MockMatchMockRecorder) GetMatch(ctx, matchID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatch", reflect.TypeOf((*MockMatch)(nil).GetMatch), ctx, matchID)
}

// GetMatchesByUser mocks base method.
func (m *MockMatch) GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*entity.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchesByUser", ctx, userID)
	ret0, _ := ret[0].([]*entity.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchesByUser indicates an expected call of GetMatchesByUser.
func (mr *MockMatchMockRecorder) GetMatchesByUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchesByUser", reflect.TypeOf((*MockMatch)(nil).GetMatchesByUser), ctx, userID)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: team.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	entity "github.com/FlowingSPDG/get5loader/backend/entity"
	usecase "github.com/FlowingSPDG/get5loader/backend/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockTeam is a mock of Team interface.
type MockTeam struct {
	ctrl     *gomock.Controller
	recorder *MockTeamMockRecorder
}

// MockTeamMockRecorder is the mock recorder for MockTeam.
type MockTeamMockRecorder struct {
	mock *MockTeam
}

// NewMockTeam creates a new mock instance.
func NewMockTeam(ctrl *gomock.Controller) *MockTeam {
	mock := &MockTeam{ctrl: ctrl}
	mock.recorder = &MockTeamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeam) EXPECT() *MockTeamMockRecorder {
	return m.recorder
}

// BatchGetTeams mocks base method.
func (m *MockTeam) BatchGetTeams(ctx context.Context, teamIDs []entity.TeamID) (map[entity.TeamID]*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetTeams", ctx, teamIDs)
	ret0, _ := ret[0].(map[entity.TeamID]*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetTeams indicates an expected call of BatchGetTeams.
func (mr *MockTeamMockRecorder) BatchGetTeams(ctx, teamIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetTeams", reflect.TypeOf((*MockTeam)(nil).BatchGetTeams), ctx, teamIDs)
}

// BatchGetTeamsByUsers mocks base method.
func (m *MockTeam) BatchGetTeamsByUsers(ctx context.Context, matchIDs []entity.UserID) (map[entity.UserID][]*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetTeamsByUsers", ctx, matchIDs)
	ret0, _ := ret[0].(map[entity.UserID][]*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetTeamsByUsers indicates an expected call of BatchGetTeamsByUsers.
func (mr *MockTeamMockRecorder) BatchGetTeamsByUsers(ctx, matchIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetTeamsByUsers", reflect.TypeOf((*MockTeam)(nil).BatchGetTeamsByUsers), ctx, matchIDs)
}

// GetTeam mocks base method.
func (m *MockTeam) GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeam", ctx, id)
	ret0, _ := ret[0].(*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeam indicates an expected call of GetTeam.
func (mr *MockTeamMockRecorder) GetTeam(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeam", reflect.TypeOf((*MockTeam)(nil).GetTeam), ctx, id)
}

// GetTeamsByMatch mocks base method.
func (m *MockTeam) GetTeamsByMatch(ctx context.Context, matchID entity.MatchID) (*entity.Team, *entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamsByMatch", ctx, matchID)
	ret0, _ := ret[0].(*entity.Team)
	ret1, _ := ret[1].(*entity.Team)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTeamsByMatch indicates an expected call of GetTeamsByMatch.
func (mr *MockTeamMockRecorder) GetTeamsByMatch(ctx, matchID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamsByMatch", reflect.TypeOf((*MockTeam)(nil).GetTeamsByMatch), ctx, matchID)
}

// GetTeamsByUser mocks base method.
func (m *MockTeam) GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamsByUser", ctx, userID)
	ret0, _ := ret[0].([]*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamsByUser indicates an expected call of GetTeamsByUser.
func (mr *MockTeamMockRecorder) GetTeamsByUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamsByUser", reflect.TypeOf((*MockTeam)(nil).GetTeamsByUser), ctx, userID)
}

// RegisterTeam mocks base method.
func (m *MockTeam) RegisterTeam(ctx context.Context, input usecase.RegisterTeamInput) (*entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTeam", ctx, input)
	ret0, _ := ret[0].(*entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTeam indicates an expected call of RegisterTeam.
func (mr *MockTeamMockRecorder) RegisterTeam(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTeam", reflect.TypeOf((*MockTeam)(nil).RegisterTeam), ctx, input)
}

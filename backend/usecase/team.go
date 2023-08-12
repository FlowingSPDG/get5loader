package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Team interface {
	RegisterTeam(ctx context.Context, userID entity.UserID, name string, flag string, tag string, logo string, publicTeam bool) (*entity.Team, error)
	GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error)
	GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error)
}

type team struct {
	repositoryConnector database.RepositoryConnector
}

func NewTeam(repositoryConnector database.RepositoryConnector) Team {
	return &team{
		repositoryConnector: repositoryConnector,
	}
}

func (t *team) RegisterTeam(ctx context.Context, userID entity.UserID, name string, flag string, tag string, logo string, publicTeam bool) (*entity.Team, error) {
	if err := t.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer t.repositoryConnector.Close()

	repository := t.repositoryConnector.GetTeamsRepository()

	teamID, err := repository.AddTeam(ctx, userID, name, tag, flag, logo, publicTeam)
	if err != nil {
		return nil, err
	}
	team, err := repository.GetTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return team, nil
}

// GetTeam implements Team.
func (t *team) GetTeam(ctx context.Context, id entity.TeamID) (*entity.Team, error) {
	if err := t.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer t.repositoryConnector.Close()

	repository := t.repositoryConnector.GetTeamsRepository()

	team, err := repository.GetTeam(ctx, id)
	if err != nil {
		return nil, err
	}
	return team, nil
}

// GetTeamsByUser implements Team.
func (t *team) GetTeamsByUser(ctx context.Context, userID entity.UserID) ([]*entity.Team, error) {
	if err := t.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer t.repositoryConnector.Close()

	repository := t.repositoryConnector.GetTeamsRepository()

	teams, err := repository.GetTeamsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

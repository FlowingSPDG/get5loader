package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Team interface {
	RegisterTeam(ctx context.Context, userID entity.UserID, name string, flag string, tag string, logo string, publicTeam bool) (entity.TeamID, error)
}

type team struct {
	repositoryConnector database.RepositoryConnector
}

func NewTeam(repositoryConnector database.RepositoryConnector) Team {
	return &team{
		repositoryConnector: repositoryConnector,
	}
}

func (t *team) RegisterTeam(ctx context.Context, userID entity.UserID, name string, flag string, tag string, logo string, publicTeam bool) (entity.TeamID, error) {
	if err := t.repositoryConnector.Open(); err != nil {
		return "", err
	}
	defer t.repositoryConnector.Close()

	repository := t.repositoryConnector.GetTeamsRepository()

	teamID, err := repository.AddTeam(ctx, userID, name, tag, flag, logo, publicTeam)
	if err != nil {
		return "", err
	}

	return teamID, nil
}

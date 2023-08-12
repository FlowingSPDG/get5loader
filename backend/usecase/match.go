package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Match interface {
	GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Match, error)
}

type match struct {
	repositoryConnector database.RepositoryConnector
}

func NewMatch(
	repositoryConnector database.RepositoryConnector,
) Match {
	return &match{
		repositoryConnector: repositoryConnector,
	}
}

func (gm *match) GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Match, error) {
	if err := gm.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer gm.repositoryConnector.Close()

	repository := gm.repositoryConnector.GetMatchesRepository()

	match, err := repository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	return match, nil
}

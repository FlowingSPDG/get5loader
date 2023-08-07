package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type GetMatch interface {
	Handle(ctx context.Context, matchID entity.MatchID) (*entity.Match, error)
}

type getMatch struct {
	repositoryConnector database.RepositoryConnector
}

func NewGetMatch(
	repositoryConnector database.RepositoryConnector,
) GetMatch {
	return &getMatch{
		repositoryConnector: repositoryConnector,
	}
}

// Handle implements GetMatchInfo.
func (gm *getMatch) Handle(ctx context.Context, matchID entity.MatchID) (*entity.Match, error) {
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

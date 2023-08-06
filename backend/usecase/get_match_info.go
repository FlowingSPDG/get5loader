package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type GetMatch interface {
	Handle(ctx context.Context, matchID int64) (*entity.Match, error)
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
func (gm *getMatch) Handle(ctx context.Context, matchID int64) (*entity.Match, error) {
	repository, err := gm.repositoryConnector.OpenMatchesRepository()
	if err != nil {
		return nil, err
	}
	defer gm.repositoryConnector.Close()

	match, err := repository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	return match, nil
}

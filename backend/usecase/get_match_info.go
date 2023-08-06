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
	matchesRepositoryConnector database.RepositoryConnector[database.MatchesRepository]
}

func NewGetMatch(
	matchesRepositoryConnector database.RepositoryConnector[database.MatchesRepository],
) GetMatch {
	return &getMatch{
		matchesRepositoryConnector: matchesRepositoryConnector,
	}
}

// Handle implements GetMatchInfo.
func (gm *getMatch) Handle(ctx context.Context, matchID int64) (*entity.Match, error) {
	repository, err := gm.matchesRepositoryConnector.Open()
	if err != nil {
		return nil, err
	}
	defer gm.matchesRepositoryConnector.Close()

	match, err := repository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	return match, nil
}

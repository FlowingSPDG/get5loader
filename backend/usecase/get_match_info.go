package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/matches"
)

type GetMatch interface {
	Handle(ctx context.Context, matchID int64) (*entity.Match, error)
}

type getMatch struct {
	mysqlConnector database.DBConnector
}

func NewGetMatch(
	mysqlConnector database.DBConnector,
) GetMatch {
	return &getMatch{
		mysqlConnector: mysqlConnector,
	}
}

// Handle implements GetMatchInfo.
func (gm *getMatch) Handle(ctx context.Context, matchID int64) (*entity.Match, error) {
	db, err := gm.mysqlConnector.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repository := matches.NewMatchRepository(db)
	match, err := repository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	return match, nil
}

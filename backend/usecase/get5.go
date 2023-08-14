package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Get5 interface {
	GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Get5Match, error)
}

type get5 struct {
}

func NewGet5() Get5 {
	return &get5{}
}

// GetMatch implements Get5.
func (g *get5) GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Get5Match, error) {
	repositoryConnector := database.GetConnection(ctx)

	matchRepository := repositoryConnector.GetMatchesRepository()
	teamRepository := repositoryConnector.GetTeamsRepository()
	playerRepository := repositoryConnector.GetPlayersRepository()

	match, err := matchRepository.GetMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	team1, err := teamRepository.GetTeam(ctx, match.Team1ID)
	if err != nil {
		return nil, err
	}
	team1players, err := playerRepository.GetPlayersByTeam(ctx, match.Team1ID)
	if err != nil {
		return nil, err
	}

	team2, err := teamRepository.GetTeam(ctx, match.Team2ID)
	if err != nil {
		return nil, err
	}
	team2players, err := playerRepository.GetPlayersByTeam(ctx, match.Team2ID)
	if err != nil {
		return nil, err
	}

	return convertGet5Match(match, team1, team2, team1players, team2players), nil
}

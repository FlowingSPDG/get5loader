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
	repositoryConnector database.RepositoryConnector
}

func NewGet5(
	repositoryConnector database.RepositoryConnector,
) Get5 {
	return &get5{
		repositoryConnector: repositoryConnector,
	}
}

// GetMatch implements Get5.
func (g *get5) GetMatch(ctx context.Context, matchID entity.MatchID) (*entity.Get5Match, error) {
	if err := g.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer g.repositoryConnector.Close()

	matchRepository := g.repositoryConnector.GetMatchesRepository()
	teamRepository := g.repositoryConnector.GetTeamsRepository()
	playerRepository := g.repositoryConnector.GetPlayersRepository()

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

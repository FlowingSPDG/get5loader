package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Player interface {
	GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Player, error)
}

type player struct {
	repositoryConnector database.RepositoryConnector
}

func NewPlayer(repositoryConnector database.RepositoryConnector) Player {
	return &player{
		repositoryConnector: repositoryConnector,
	}
}

// GetPlayersByTeam implements Player.
func (p *player) GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Player, error) {
	if err := p.repositoryConnector.Open(); err != nil {
		return nil, err
	}
	defer p.repositoryConnector.Close()

	playerRepository := p.repositoryConnector.GetPlayersRepository()

	players, err := playerRepository.GetPlayersByTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return convertPlayers(players), nil
}

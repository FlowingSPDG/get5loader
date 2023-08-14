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
}

func NewPlayer() Player {
	return &player{}
}

// GetPlayersByTeam implements Player.
func (p *player) GetPlayersByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Player, error) {
	repositoryConnector := database.GetConnection(ctx)

	playerRepository := repositoryConnector.GetPlayersRepository()

	players, err := playerRepository.GetPlayersByTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	return convertPlayers(players), nil
}

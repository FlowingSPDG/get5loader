package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type PlayerStat interface {
	BatchGetPlayerStatsByMapstat(ctx context.Context, mapstatIDs []entity.MapStatsID) (map[entity.MapStatsID][]*entity.PlayerStat, error)
}

type playerStat struct {
}

func NewPlayerStat() PlayerStat {
	return &playerStat{}
}

// BatchGetPlayerStatsByMapstat implements PlayerStat.
func (ps *playerStat) BatchGetPlayerStatsByMapstat(ctx context.Context, mapstatIDs []entity.MapStatsID) (map[entity.MapStatsID][]*entity.PlayerStat, error) {
	repositoryConnector := database.GetConnection(ctx)

	playerStatRepository := repositoryConnector.GetPlayerStatRepository()

	playerStats, err := playerStatRepository.GetPlayerStatsByMapstats(ctx, mapstatIDs)
	if err != nil {
		return nil, err
	}

	ret := make(map[entity.MapStatsID][]*entity.PlayerStat, len(mapstatIDs))
	for teamID, players := range playerStats {
		ret[teamID] = convertPlayerStats(players)
	}

	return ret, nil
}

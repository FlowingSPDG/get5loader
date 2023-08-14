package usecase

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type Mapstat interface {
	GetMapStat(ctx context.Context, id entity.MapStatsID) (*entity.MapStat, error)
	GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStat, error)
	// GetMapStatsByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.MapStat, error)

	// BATCH
	BatchGetMapstatsByMatch(ctx context.Context, matchIDs []entity.MatchID) (map[entity.MatchID][]*entity.MapStat, error)
}

type mapstat struct {
}

func NewMapStats() Mapstat {
	return &mapstat{}
}

// GetMapStats implements Mapstats.
func (m *mapstat) GetMapStat(ctx context.Context, id entity.MapStatsID) (*entity.MapStat, error) {
	repositoryConnector := database.GetConnection(ctx)

	MapStatRepository := repositoryConnector.GetMapStatRepository()

	mapstats, err := MapStatRepository.GetMapStat(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertMapstat(mapstats), nil
}

// GetMapStatsByMatch implements Mapstats.
func (m *mapstat) GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*entity.MapStat, error) {
	repositoryConnector := database.GetConnection(ctx)

	MapStatRepository := repositoryConnector.GetMapStatRepository()

	mapstats, err := MapStatRepository.GetMapStatsByMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.MapStat, 0, len(mapstats))
	for _, mapstat := range mapstats {
		ret = append(ret, convertMapstat(mapstat))
	}

	return ret, nil
}

// BatchGetMapstatsByMatch implements Mapstat.
func (m *mapstat) BatchGetMapstatsByMatch(ctx context.Context, matchIDs []entity.MatchID) (map[entity.MatchID][]*entity.MapStat, error) {
	repositoryConnector := database.GetConnection(ctx)

	mapStatRepository := repositoryConnector.GetMapStatRepository()

	mapstats, err := mapStatRepository.GetMapStatsByMatches(ctx, matchIDs)
	if err != nil {
		return nil, err
	}

	ret := make(map[entity.MatchID][]*entity.MapStat, len(matchIDs))
	// nilが渡されるのを防ぐため、空のスライスを生成する
	for matchID, mapstats := range mapstats {
		ret[matchID] = convertMapstats(mapstats)
	}
	return ret, nil
}

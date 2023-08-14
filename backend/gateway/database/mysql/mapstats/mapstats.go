package mapstats

import (
	"context"
	"database/sql"
	"errors"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	mapstats_gen "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/mapstats/generated"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type MapStatRepository struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *mapstats_gen.Queries
}

func NewMapStatRepository(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.MapStatRepository {
	queries := mapstats_gen.New(db)
	return &MapStatRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewMapStatRepositoryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.MapStatRepository {
	queries := mapstats_gen.New(db).WithTx(tx)
	return &MapStatRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// GetMapStats implements database.MapStatRepository.
func (msr *MapStatRepository) GetMapStat(ctx context.Context, id entity.MapStatsID) (*database.MapStat, error) {
	res, err := msr.queries.GetMapStat(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	winner := entity.TeamID(res.Winner.String)

	return &database.MapStat{
		ID:         entity.MapStatsID(res.ID),
		MatchID:    entity.MatchID(res.MatchID),
		MapNumber:  res.MapNumber,
		MapName:    res.MapName,
		StartTime:  &res.StartTime.Time,
		EndTime:    &res.EndTime.Time,
		Winner:     &winner,
		Team1Score: res.Team1Score,
		Team2Score: res.Team2Score,
	}, nil
}

// GetMapStatsByMatch implements database.MapStatRepository.
func (msr *MapStatRepository) GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*database.MapStat, error) {
	res, err := msr.queries.GetMapStatsByMatch(ctx, string(matchID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*database.MapStat{}, nil
		}
		return nil, database.NewInternalError(err)
	}

	mapStats := make([]*database.MapStat, 0, len(res))
	for _, m := range res {
		winner := entity.TeamID(m.Winner.String)
		mapStats = append(mapStats, &database.MapStat{
			ID:         entity.MapStatsID(m.ID),
			MatchID:    entity.MatchID(m.MatchID),
			MapNumber:  m.MapNumber,
			MapName:    m.MapName,
			StartTime:  &m.StartTime.Time,
			EndTime:    &m.EndTime.Time,
			Winner:     &winner,
			Team1Score: m.Team1Score,
			Team2Score: m.Team2Score,
		})
	}

	return mapStats, nil
}

// GetMapStatsByMatches implements database.MapStatRepository.
func (msr *MapStatRepository) GetMapStatsByMatches(ctx context.Context, matchIDs []entity.MatchID) (map[entity.MatchID][]*database.MapStat, error) {
	ids := database.IDsToString(matchIDs)
	mapstats, err := msr.queries.GetMapStatsByMatches(ctx, ids)
	if err != nil {
		return nil, database.NewInternalError(err)
	}

	ret := make(map[entity.MatchID][]*database.MapStat, len(matchIDs))
	// nilが渡されるのを防ぐため、空のスライスを生成する
	for _, matchID := range matchIDs {
		ret[matchID] = make([]*database.MapStat, 0)
	}

	for _, m := range mapstats {
		winner := entity.TeamID(m.Winner.String)
		ret[entity.MatchID(m.MatchID)] = append(ret[entity.MatchID(m.MatchID)], &database.MapStat{
			ID:         entity.MapStatsID(m.ID),
			MatchID:    entity.MatchID(m.MatchID),
			MapNumber:  m.MapNumber,
			MapName:    m.MapName,
			StartTime:  &m.StartTime.Time,
			EndTime:    &m.EndTime.Time,
			Winner:     &winner,
			Team1Score: m.Team1Score,
			Team2Score: m.Team2Score,
		})
	}
	return ret, nil
}

// GetMapStatsByMatchAndMap implements database.MapStatRepository.
func (msr *MapStatRepository) GetMapStatsByMatchAndMap(ctx context.Context, matchID entity.MatchID, mapNumber uint32) (*database.MapStat, error) {
	res, err := msr.queries.GetMapStatsByMatchAndMap(ctx, mapstats_gen.GetMapStatsByMatchAndMapParams{
		MatchID:   string(matchID),
		MapNumber: mapNumber,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	winner := entity.TeamID(res.Winner.String)

	return &database.MapStat{
		ID:         entity.MapStatsID(res.ID),
		MatchID:    entity.MatchID(res.MatchID),
		MapNumber:  res.MapNumber,
		MapName:    res.MapName,
		StartTime:  &res.StartTime.Time,
		EndTime:    &res.EndTime.Time,
		Winner:     &winner,
		Team1Score: res.Team1Score,
		Team2Score: res.Team2Score,
	}, nil
}

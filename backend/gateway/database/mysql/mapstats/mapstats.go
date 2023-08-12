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

type mapStatsRepository struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *mapstats_gen.Queries
}

func NewMapStatsRepository(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.MapStatsRepository {
	queries := mapstats_gen.New(db)
	return &mapStatsRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewMapStatsRepositoryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.MapStatsRepository {
	queries := mapstats_gen.New(db).WithTx(tx)
	return &mapStatsRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// GetMapStats implements database.MapStatsRepository.
func (msr *mapStatsRepository) GetMapStats(ctx context.Context, id entity.MapStatsID) (*database.MapStats, error) {
	res, err := msr.queries.GetMapStats(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	winner := entity.TeamID(res.Winner.String)

	return &database.MapStats{
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

// GetMapStatsByMatch implements database.MapStatsRepository.
func (msr *mapStatsRepository) GetMapStatsByMatch(ctx context.Context, matchID entity.MatchID) ([]*database.MapStats, error) {
	res, err := msr.queries.GetMapStatsByMatch(ctx, string(matchID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	mapStats := make([]*database.MapStats, 0, len(res))
	for _, m := range res {
		winner := entity.TeamID(m.Winner.String)
		mapStats = append(mapStats, &database.MapStats{
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

// GetMapStatsByMatchAndMap implements database.MapStatsRepository.
func (msr *mapStatsRepository) GetMapStatsByMatchAndMap(ctx context.Context, matchID entity.MatchID, mapNumber uint32) (*database.MapStats, error) {
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

	return &database.MapStats{
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

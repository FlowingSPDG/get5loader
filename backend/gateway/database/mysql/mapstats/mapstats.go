package mapstats

import (
	"context"
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	mapstats_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/mapstats/generated"
)

type mapStatsRepository struct {
	dsn string
}

func NewMapStatsRepository(dsn string) database.MapStatsRepository {
	return &mapStatsRepository{
		dsn: dsn,
	}
}

func (mr *mapStatsRepository) open() (*sql.DB, error) {
	return sql.Open("mysql", mr.dsn)
}

// GetMapStats implements database.MapStatsRepository.
func (msr *mapStatsRepository) GetMapStats(ctx context.Context, id int64) (*entity.MapStats, error) {
	db, err := msr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := mapstats_gen.New(db)

	res, err := queries.GetMapStats(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.MapStats{
		ID:         res.ID,
		MatchID:    res.MatchID,
		MapNumber:  res.MapNumber,
		MapName:    res.MapName,
		StartTime:  &res.StartTime.Time,
		EndTime:    &res.EndTime.Time,
		Winner:     &res.Winner.Int64,
		Team1Score: res.Team1Score,
		Team2Score: res.Team2Score,
	}, nil
}

// GetMapStatsByMatch implements database.MapStatsRepository.
func (msr *mapStatsRepository) GetMapStatsByMatch(ctx context.Context, matchID int64) ([]*entity.MapStats, error) {
	db, err := msr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := mapstats_gen.New(db)

	res, err := queries.GetMapStatsByMatch(ctx, matchID)
	if err != nil {
		return nil, err
	}

	mapStats := make([]*entity.MapStats, 0, len(res))
	for _, m := range res {
		mapStats = append(mapStats, &entity.MapStats{
			ID:         m.ID,
			MatchID:    m.MatchID,
			MapNumber:  m.MapNumber,
			MapName:    m.MapName,
			StartTime:  &m.StartTime.Time,
			EndTime:    &m.EndTime.Time,
			Winner:     &m.Winner.Int64,
			Team1Score: m.Team1Score,
			Team2Score: m.Team2Score,
		})
	}

	return mapStats, nil
}

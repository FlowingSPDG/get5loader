package matches

import (
	"context"
	"database/sql"
	"time"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	matches_gen "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/matches/generated"
)

type matchRepository struct {
	dsn string
}

func NewMatchRepository(dsn string) database.MatchesRepository {
	return &matchRepository{
		dsn: dsn,
	}
}

func (mr *matchRepository) open() (*sql.DB, error) {
	return sql.Open("mysql", mr.dsn)
}

// AddMatch implements database.MatchRepository.
func (mr *matchRepository) AddMatch(ctx context.Context, userID int64, serverID int64, team1ID int64, team2ID int64, startTime time.Time, endTime time.Time, maxMaps int32, title string, skipVeto bool, apiKey string) (*entity.Match, error) {
	db, err := mr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	res, err := queries.AddMatch(ctx, matches_gen.AddMatchParams{
		UserID:    userID,
		ServerID:  serverID,
		Team1ID:   team1ID,
		Team2ID:   team2ID,
		StartTime: sql.NullTime{Valid: false},
		EndTime:   sql.NullTime{Valid: false},
		MaxMaps:   maxMaps,
		Title:     title,
		SkipVeto:  skipVeto,
		ApiKey:    apiKey,
	})

	if err != nil {
		return nil, err
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	match, err := queries.GetMatch(ctx, insertedID)
	if err != nil {
		return nil, err
	}

	return &entity.Match{
		ID:        match.ID,
		UserID:    match.UserID,
		ServerID:  match.ServerID,
		Team1ID:   match.Team1ID,
		Team2ID:   match.Team2ID,
		StartTime: &match.StartTime.Time,
		EndTime:   &match.EndTime.Time,
		MaxMaps:   match.MaxMaps,
		Title:     match.Title,
		SkipVeto:  match.SkipVeto,
		APIKey:    match.ApiKey,
	}, nil
}

// GetMatch implements database.MatchRepository.
func (mr *matchRepository) GetMatch(ctx context.Context, id int64) (*entity.Match, error) {
	db, err := mr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	match, err := queries.GetMatch(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Match{
		ID:         match.ID,
		UserID:     match.UserID,
		ServerID:   match.ServerID,
		Team1ID:    match.Team1ID,
		Team2ID:    match.Team2ID,
		Winner:     &match.Winner.Int64,
		Cancelled:  match.Cancelled,
		StartTime:  &match.StartTime.Time,
		EndTime:    &match.EndTime.Time,
		MaxMaps:    match.MaxMaps,
		Title:      match.Title,
		SkipVeto:   match.SkipVeto,
		APIKey:     match.ApiKey,
		Team1Score: match.Team1Score,
		Team2Score: match.Team2Score,
		Forfeit:    &match.Forfeit.Bool,
	}, nil
}

// GetMatchesByUser implements database.MatchRepository.
func (mr *matchRepository) GetMatchesByUser(ctx context.Context, userID int64) ([]*entity.Match, error) {
	db, err := mr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	matches, err := queries.GetMatchesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var ret []*entity.Match
	for _, match := range matches {
		ret = append(ret, &entity.Match{
			ID:         match.ID,
			UserID:     match.UserID,
			ServerID:   match.ServerID,
			Team1ID:    match.Team1ID,
			Team2ID:    match.Team2ID,
			Winner:     &match.Winner.Int64,
			Cancelled:  match.Cancelled,
			StartTime:  &match.StartTime.Time,
			EndTime:    &match.EndTime.Time,
			MaxMaps:    match.MaxMaps,
			Title:      match.Title,
			SkipVeto:   match.SkipVeto,
			APIKey:     match.ApiKey,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Forfeit:    &match.Forfeit.Bool,
		})
	}

	return ret, nil
}

// CancelMatch implements database.MatchRepository.
func (mr *matchRepository) CancelMatch(ctx context.Context, matchID int64) error {
	db, err := mr.open()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	if _, err = queries.CancelMatch(ctx, matchID); err != nil {
		return err
	}

	return nil
}

// GetMatchesByTeam implements database.MatchRepository.
func (mr *matchRepository) GetMatchesByTeam(ctx context.Context, teamID int64) ([]*entity.Match, error) {
	db, err := mr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	matches, err := queries.GetMatchesByTeam(ctx, matches_gen.GetMatchesByTeamParams{
		Team1ID: teamID,
		Team2ID: teamID,
	})
	if err != nil {
		return nil, err
	}

	var ret []*entity.Match
	for _, match := range matches {
		ret = append(ret, &entity.Match{
			ID:         match.ID,
			UserID:     match.UserID,
			ServerID:   match.ServerID,
			Team1ID:    match.Team1ID,
			Team2ID:    match.Team2ID,
			Winner:     &match.Winner.Int64,
			Cancelled:  match.Cancelled,
			StartTime:  &match.StartTime.Time,
			EndTime:    &match.EndTime.Time,
			MaxMaps:    match.MaxMaps,
			Title:      match.Title,
			SkipVeto:   match.SkipVeto,
			APIKey:     match.ApiKey,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Forfeit:    &match.Forfeit.Bool,
		})
	}

	return ret, nil
}

// GetMatchesByWinner implements database.MatchRepository.
func (mr *matchRepository) GetMatchesByWinner(ctx context.Context, teamID int64) ([]*entity.Match, error) {
	db, err := mr.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	matches, err := queries.GetMatchesByWinner(ctx, sql.NullInt64{Int64: teamID, Valid: true})
	if err != nil {
		return nil, err
	}

	var ret []*entity.Match
	for _, match := range matches {
		ret = append(ret, &entity.Match{
			ID:         match.ID,
			UserID:     match.UserID,
			ServerID:   match.ServerID,
			Team1ID:    match.Team1ID,
			Team2ID:    match.Team2ID,
			Winner:     &match.Winner.Int64,
			Cancelled:  match.Cancelled,
			StartTime:  &match.StartTime.Time,
			EndTime:    &match.EndTime.Time,
			MaxMaps:    match.MaxMaps,
			Title:      match.Title,
			SkipVeto:   match.SkipVeto,
			APIKey:     match.ApiKey,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Forfeit:    &match.Forfeit.Bool,
		})
	}

	return ret, nil
}

// StartMatch implements database.MatchRepository.
func (mr *matchRepository) StartMatch(ctx context.Context, matchID int64) error {
	db, err := mr.open()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	if _, err = queries.StartMatch(ctx, matches_gen.StartMatchParams{
		ID:        matchID,
		StartTime: sql.NullTime{Valid: true, Time: time.Now()},
	}); err != nil {
		return err
	}

	return nil
}

// UpdateMatchWinner implements database.MatchRepository.
func (mr *matchRepository) UpdateMatchWinner(ctx context.Context, matchID int64, winnerID int64) error {
	db, err := mr.open()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	if _, err = queries.UpdateMatchWinner(ctx, matches_gen.UpdateMatchWinnerParams{
		ID:     matchID,
		Winner: sql.NullInt64{Int64: winnerID, Valid: true},
	}); err != nil {
		return err
	}

	return nil
}

// UpdateTeam1Score implements database.MatchRepository.
func (mr *matchRepository) UpdateTeam1Score(ctx context.Context, matchID int64, score int32) error {
	db, err := mr.open()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	if _, err = queries.UpdateTeam1Score(ctx, matches_gen.UpdateTeam1ScoreParams{
		ID:         matchID,
		Team1Score: score,
	}); err != nil {
		return err
	}

	return nil
}

// UpdateTeam2Score implements database.MatchRepository.
func (mr *matchRepository) UpdateTeam2Score(ctx context.Context, matchID int64, score int32) error {
	db, err := mr.open()
	if err != nil {
		return err
	}
	defer db.Close()

	queries := matches_gen.New(db)

	if _, err = queries.UpdateTeam2Score(ctx, matches_gen.UpdateTeam2ScoreParams{
		ID:         matchID,
		Team2Score: score,
	}); err != nil {
		return err
	}

	return nil
}

package matches

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	matches_gen "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/matches/generated"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type matchRepository struct {
	uuidGenerator uuid.UUIDGenerator
	queries       *matches_gen.Queries
}

func NewMatchRepository(uuidGenerator uuid.UUIDGenerator, db *sql.DB) database.MatchesRepository {
	queries := matches_gen.New(db)
	return &matchRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

func NewMatchRepositoryWithTx(uuidGenerator uuid.UUIDGenerator, db *sql.DB, tx *sql.Tx) database.MatchesRepository {
	queries := matches_gen.New(db).WithTx(tx)
	return &matchRepository{
		uuidGenerator: uuidGenerator,
		queries:       queries,
	}
}

// AddMatch implements database.MatchRepository.
func (mr *matchRepository) AddMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID entity.TeamID, team2ID entity.TeamID, startTime time.Time, endTime time.Time, maxMaps int32, title string, skipVeto bool, apiKey string) error {
	if _, err := mr.queries.AddMatch(ctx, matches_gen.AddMatchParams{
		ID:        mr.uuidGenerator.Generate(),
		UserID:    string(userID),
		ServerID:  string(serverID),
		Team1ID:   string(team1ID),
		Team2ID:   string(team2ID),
		StartTime: sql.NullTime{Valid: false},
		EndTime:   sql.NullTime{Valid: false},
		MaxMaps:   maxMaps,
		Title:     title,
		SkipVeto:  skipVeto,
		ApiKey:    apiKey,
		Status:    int32(entity.MATCH_STATUS_PENDING),
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// GetMatch implements database.MatchRepository.
func (mr *matchRepository) GetMatch(ctx context.Context, id entity.MatchID) (*entity.Match, error) {
	match, err := mr.queries.GetMatch(ctx, string(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	winner := entity.TeamID(match.Winner.String)
	forfeit := match.Forfeit.Bool

	return &entity.Match{
		ID:         entity.MatchID(match.ID),
		UserID:     entity.UserID(match.UserID),
		ServerID:   entity.GameServerID(match.ServerID),
		Team1ID:    entity.TeamID(match.Team1ID),
		Team2ID:    entity.TeamID(match.Team2ID),
		Winner:     &winner,
		StartTime:  &match.StartTime.Time,
		EndTime:    &match.EndTime.Time,
		MaxMaps:    match.MaxMaps,
		Title:      match.Title,
		SkipVeto:   match.SkipVeto,
		APIKey:     match.ApiKey,
		Team1Score: match.Team1Score,
		Team2Score: match.Team2Score,
		Forfeit:    &forfeit,
		Status:     entity.MATCH_STATUS(match.Status),
	}, nil
}

// GetMatchesByUser implements database.MatchRepository.
func (mr *matchRepository) GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*entity.Match, error) {
	matches, err := mr.queries.GetMatchesByUser(ctx, string(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*entity.Match, 0, len(matches))
	for _, match := range matches {
		winner := entity.TeamID(match.Winner.String)
		ret = append(ret, &entity.Match{
			ID:         entity.MatchID(match.ID),
			UserID:     entity.UserID(match.UserID),
			ServerID:   entity.GameServerID(match.ServerID),
			Team1ID:    entity.TeamID(match.Team1ID),
			Team2ID:    entity.TeamID(match.Team2ID),
			Winner:     &winner,
			StartTime:  &match.StartTime.Time,
			EndTime:    &match.EndTime.Time,
			MaxMaps:    match.MaxMaps,
			Title:      match.Title,
			SkipVeto:   match.SkipVeto,
			APIKey:     match.ApiKey,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Forfeit:    &match.Forfeit.Bool,
			Status:     entity.MATCH_STATUS(match.Status),
		})
	}

	return ret, nil
}

// CancelMatch implements database.MatchRepository.
func (mr *matchRepository) CancelMatch(ctx context.Context, matchID entity.MatchID) error {
	if _, err := mr.queries.CancelMatch(ctx, string(matchID)); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// GetMatchesByTeam implements database.MatchRepository.
func (mr *matchRepository) GetMatchesByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Match, error) {
	matches, err := mr.queries.GetMatchesByTeam(ctx, matches_gen.GetMatchesByTeamParams{
		Team1ID: string(teamID),
		Team2ID: string(teamID),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*entity.Match, 0, len(matches))
	for _, match := range matches {
		winner := entity.TeamID(match.Winner.String)
		ret = append(ret, &entity.Match{
			ID:         entity.MatchID(match.ID),
			UserID:     entity.UserID(match.UserID),
			ServerID:   entity.GameServerID(match.ServerID),
			Team1ID:    entity.TeamID(match.Team1ID),
			Team2ID:    entity.TeamID(match.Team2ID),
			Winner:     &winner,
			StartTime:  &match.StartTime.Time,
			EndTime:    &match.EndTime.Time,
			MaxMaps:    match.MaxMaps,
			Title:      match.Title,
			SkipVeto:   match.SkipVeto,
			APIKey:     match.ApiKey,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Forfeit:    &match.Forfeit.Bool,
			Status:     entity.MATCH_STATUS(match.Status),
		})
	}

	return ret, nil
}

// GetMatchesByWinner implements database.MatchRepository.
func (mr *matchRepository) GetMatchesByWinner(ctx context.Context, teamID entity.TeamID) ([]*entity.Match, error) {
	matches, err := mr.queries.GetMatchesByWinner(ctx, sql.NullString{String: string(teamID), Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.NewNotFoundError(err)
		}
		return nil, database.NewInternalError(err)
	}

	ret := make([]*entity.Match, 0, len(matches))
	for _, match := range matches {
		winner := entity.TeamID(match.Winner.String)
		ret = append(ret, &entity.Match{
			ID:         entity.MatchID(match.ID),
			UserID:     entity.UserID(match.UserID),
			ServerID:   entity.GameServerID(match.ServerID),
			Team1ID:    entity.TeamID(match.Team1ID),
			Team2ID:    entity.TeamID(match.Team2ID),
			Winner:     &winner,
			StartTime:  &match.StartTime.Time,
			EndTime:    &match.EndTime.Time,
			MaxMaps:    match.MaxMaps,
			Title:      match.Title,
			SkipVeto:   match.SkipVeto,
			APIKey:     match.ApiKey,
			Team1Score: match.Team1Score,
			Team2Score: match.Team2Score,
			Forfeit:    &match.Forfeit.Bool,
			Status:     entity.MATCH_STATUS(match.Status),
		})
	}

	return ret, nil
}

// StartMatch implements database.MatchRepository.
func (mr *matchRepository) StartMatch(ctx context.Context, matchID entity.MatchID) error {
	if _, err := mr.queries.StartMatch(ctx, matches_gen.StartMatchParams{
		ID:        string(matchID),
		StartTime: sql.NullTime{Valid: true, Time: time.Now()},
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// UpdateMatchWinner implements database.MatchRepository.
func (mr *matchRepository) UpdateMatchWinner(ctx context.Context, matchID entity.MatchID, winnerID entity.TeamID) error {
	if _, err := mr.queries.UpdateMatchWinner(ctx, matches_gen.UpdateMatchWinnerParams{
		ID:     string(matchID),
		Winner: sql.NullString{String: string(winnerID), Valid: true},
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// UpdateTeam1Score implements database.MatchRepository.
func (mr *matchRepository) UpdateTeam1Score(ctx context.Context, matchID entity.MatchID, score uint32) error {
	if _, err := mr.queries.UpdateTeam1Score(ctx, matches_gen.UpdateTeam1ScoreParams{
		ID:         string(matchID),
		Team1Score: score,
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

// UpdateTeam2Score implements database.MatchRepository.
func (mr *matchRepository) UpdateTeam2Score(ctx context.Context, matchID entity.MatchID, score uint32) error {
	if _, err := mr.queries.UpdateTeam2Score(ctx, matches_gen.UpdateTeam2ScoreParams{
		ID:         string(matchID),
		Team2Score: score,
	}); err != nil {
		return database.NewInternalError(err)
	}

	return nil
}

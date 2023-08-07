package matches

import (
	"context"
	"time"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type mockMatchesRepository struct {
	matchesToReturnID     map[entity.MatchID]*entity.Match
	matchesToReturnUserID map[entity.UserID][]*entity.Match
	matchesToReturnTeamID map[entity.TeamID][]*entity.Match
	matchesToReturnWinner map[entity.TeamID][]*entity.Match
}

func NewMockMatchesRepository(
	matchesToReturnID map[entity.MatchID]*entity.Match,
	matchesToReturnUserID map[entity.UserID][]*entity.Match,
	matchesToReturnTeamID map[entity.TeamID][]*entity.Match,
	matchesToReturnWinner map[entity.TeamID][]*entity.Match,
) database.MatchesRepository {
	return &mockMatchesRepository{
		matchesToReturnID:     matchesToReturnID,
		matchesToReturnUserID: matchesToReturnUserID,
		matchesToReturnTeamID: matchesToReturnTeamID,
		matchesToReturnWinner: matchesToReturnWinner,
	}
}

// AddMatch implements database.MatchesRepository.
func (mmr *mockMatchesRepository) AddMatch(ctx context.Context, userID entity.UserID, serverID entity.GameServerID, team1ID entity.TeamID, team2ID entity.TeamID, startTime time.Time, endTime time.Time, maxMaps int32, title string, skipVeto bool, apiKey string) error {
	return nil
}

// CancelMatch implements database.MatchesRepository.
func (mmr *mockMatchesRepository) CancelMatch(ctx context.Context, matchID entity.MatchID) error {
	return nil
}

// GetMatch implements database.MatchesRepository.
func (mmr *mockMatchesRepository) GetMatch(ctx context.Context, id entity.MatchID) (*entity.Match, error) {
	v, ok := mmr.matchesToReturnID[id]
	if !ok {
		return nil, database.ErrNotFound
	}
	return v, nil
}

// GetMatchesByTeam implements database.MatchesRepository.
func (mmr *mockMatchesRepository) GetMatchesByTeam(ctx context.Context, teamID entity.TeamID) ([]*entity.Match, error) {
	v, ok := mmr.matchesToReturnTeamID[teamID]
	if !ok {
		return nil, database.ErrNotFound
	}
	return v, nil
}

// GetMatchesByUser implements database.MatchesRepository.
func (mmr *mockMatchesRepository) GetMatchesByUser(ctx context.Context, userID entity.UserID) ([]*entity.Match, error) {
	v, ok := mmr.matchesToReturnUserID[userID]
	if !ok {
		return nil, database.ErrNotFound
	}
	return v, nil
}

// GetMatchesByWinner implements database.MatchesRepository.
func (mmr *mockMatchesRepository) GetMatchesByWinner(ctx context.Context, teamID entity.TeamID) ([]*entity.Match, error) {
	v, ok := mmr.matchesToReturnWinner[teamID]
	if !ok {
		return nil, database.ErrNotFound
	}
	return v, nil
}

// StartMatch implements database.MatchesRepository.
func (mmr *mockMatchesRepository) StartMatch(ctx context.Context, matchID entity.MatchID) error {
	return nil
}

// UpdateMatchWinner implements database.MatchesRepository.
func (mmr *mockMatchesRepository) UpdateMatchWinner(ctx context.Context, matchID entity.MatchID, winnerID entity.TeamID) error {
	return nil
}

// UpdateTeam1Score implements database.MatchesRepository.
func (mmr *mockMatchesRepository) UpdateTeam1Score(ctx context.Context, matchID entity.MatchID, score uint32) error {
	return nil
}

// UpdateTeam2Score implements database.MatchesRepository.
func (mmr *mockMatchesRepository) UpdateTeam2Score(ctx context.Context, matchID entity.MatchID, score uint32) error {
	return nil
}

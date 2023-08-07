package connector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type mockRepositoryConnector struct {
	mockUserRepository        database.UsersRepositry
	mockGaneServersRepository database.GameServersRepository
	mockMapStatsRepository    database.MapStatsRepository
	mockMatchesRepository     database.MatchesRepository
	mockPlayerStatsRepository database.PlayerStatsRepository
	mockTeamsRepository       database.TeamsRepository
	mockPlayersRepository     database.PlayersRepository
}

func NewMockRepositoryConnector(
	mockUserRepository database.UsersRepositry,
	mockGaneServersRepository database.GameServersRepository,
	mockMapStatsRepository database.MapStatsRepository,
	mockMatchesRepository database.MatchesRepository,
	mockPlayerStatsRepository database.PlayerStatsRepository,
	mockTeamsRepository database.TeamsRepository,
	mockPlayersRepository database.PlayersRepository,
) database.RepositoryConnector {
	return &mockRepositoryConnector{
		mockUserRepository:        mockUserRepository,
		mockGaneServersRepository: mockGaneServersRepository,
		mockMapStatsRepository:    mockMapStatsRepository,
		mockMatchesRepository:     mockMatchesRepository,
		mockPlayerStatsRepository: mockPlayerStatsRepository,
		mockTeamsRepository:       mockTeamsRepository,
		mockPlayersRepository:     mockPlayersRepository,
	}
}

// Close implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) Close() error {
	return nil
}

// GetGameServersRepository implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) GetGameServersRepository() database.GameServersRepository {
	return mrc.mockGaneServersRepository
}

// GetMapStatsRepository implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) GetMapStatsRepository() database.MapStatsRepository {
	return mrc.mockMapStatsRepository
}

// GetMatchesRepository implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) GetMatchesRepository() database.MatchesRepository {
	return mrc.mockMatchesRepository
}

// GetPlayerStatsRepository implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) GetPlayerStatsRepository() database.PlayerStatsRepository {
	return mrc.mockPlayerStatsRepository
}

// GetPlayersRepository implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) GetPlayersRepository() database.PlayersRepository {
	return mrc.mockPlayersRepository
}

// GetTeamsRepository implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) GetTeamsRepository() database.TeamsRepository {
	return mrc.mockTeamsRepository
}

// GetUserRepository implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) GetUserRepository() database.UsersRepositry {
	return mrc.mockUserRepository
}

// Open implements database.RepositoryConnector.
func (mrc *mockRepositoryConnector) Open() error {
	return nil
}

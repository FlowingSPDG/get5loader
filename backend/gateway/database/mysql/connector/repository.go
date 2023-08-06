package mysqlconnector

import (
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/gameservers"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/mapstats"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/matches"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/players"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/playerstats"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/teams"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users"
)

type mysqlRepositoryConnector struct {
	connector database.DBConnector
	db        *sql.DB
}

func NewMySQLRepositoryConnector(connector database.DBConnector) database.RepositoryConnector {
	return &mysqlRepositoryConnector{connector: connector}
}

// Close implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) Close() error {
	return mrc.connector.Close()
}

func (mrc *mysqlRepositoryConnector) Open() error {
	if err := mrc.connector.Open(); err != nil {
		return err
	}

	mrc.db = mrc.connector.GetConnection()
	return nil
}

// OpenGameServersRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetGameServersRepository() (database.GameServersRepository, error) {
	repository := gameservers.NewGameServerRepository(mrc.db)
	return repository, nil
}

// OpenMapStatsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetMapStatsRepository() (database.MapStatsRepository, error) {
	repository := mapstats.NewMapStatsRepository(mrc.db)
	return repository, nil
}

// OpenMatchesRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetMatchesRepository() (database.MatchesRepository, error) {
	repository := matches.NewMatchRepository(mrc.db)
	return repository, nil
}

// OpenPlayerStatsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetPlayerStatsRepository() (database.PlayerStatsRepository, error) {
	repository := playerstats.NewPlayerStatsRepository(mrc.db)
	return repository, nil
}

// OpenPlayersRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetPlayersRepository() (database.PlayersRepository, error) {
	repository := players.NewPlayersRepository(mrc.db)
	return repository, nil
}

// OpenTeamsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetTeamsRepository() (database.TeamsRepository, error) {
	repository := teams.NewTeamssRepository(mrc.db)
	return repository, nil
}

// OpenUserRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetUserRepository() (database.UsersRepositry, error) {
	repository := users.NewUsersRepositry(mrc.db)
	return repository, nil
}

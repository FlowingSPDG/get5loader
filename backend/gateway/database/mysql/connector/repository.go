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

func (mrc *mysqlRepositoryConnector) open() error {
	if err := mrc.connector.Open(); err != nil {
		return err
	}

	mrc.db = mrc.connector.GetConnection()
	return nil
}

// OpenGameServersRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) OpenGameServersRepository() (database.GameServersRepository, error) {
	if err := mrc.open(); err != nil {
		return nil, err
	}

	repository := gameservers.NewGameServerRepository(mrc.db)
	return repository, nil
}

// OpenMapStatsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) OpenMapStatsRepository() (database.MapStatsRepository, error) {
	if err := mrc.open(); err != nil {
		return nil, err
	}

	repository := mapstats.NewMapStatsRepository(mrc.db)
	return repository, nil
}

// OpenMatchesRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) OpenMatchesRepository() (database.MatchesRepository, error) {
	if err := mrc.open(); err != nil {
		return nil, err
	}
	repository := matches.NewMatchRepository(mrc.db)
	return repository, nil
}

// OpenPlayerStatsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) OpenPlayerStatsRepository() (database.PlayerStatsRepository, error) {
	if err := mrc.open(); err != nil {
		return nil, err
	}
	repository := playerstats.NewPlayerStatsRepository(mrc.db)
	return repository, nil
}

// OpenPlayersRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) OpenPlayersRepository() (database.PlayersRepository, error) {
	if err := mrc.open(); err != nil {
		return nil, err
	}
	repository := players.NewPlayersRepository(mrc.db)
	return repository, nil
}

// OpenTeamsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) OpenTeamsRepository() (database.TeamsRepository, error) {
	if err := mrc.open(); err != nil {
		return nil, err
	}
	repository := teams.NewTeamssRepository(mrc.db)
	return repository, nil
}

// OpenUserRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) OpenUserRepository() (database.UsersRepositry, error) {
	if err := mrc.open(); err != nil {
		return nil, err
	}
	repository := users.NewUsersRepositry(mrc.db)
	return repository, nil
}

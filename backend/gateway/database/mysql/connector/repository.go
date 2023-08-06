package mysqlconnector

import (
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
	return nil
}

// OpenGameServersRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetGameServersRepository() database.GameServersRepository {
	conn := mrc.connector.GetConnection()
	return gameservers.NewGameServerRepository(conn)
}

// OpenMapStatsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetMapStatsRepository() database.MapStatsRepository {
	conn := mrc.connector.GetConnection()
	return mapstats.NewMapStatsRepository(conn)
}

// OpenMatchesRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetMatchesRepository() database.MatchesRepository {
	conn := mrc.connector.GetConnection()
	return matches.NewMatchRepository(conn)
}

// OpenPlayerStatsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetPlayerStatsRepository() database.PlayerStatsRepository {
	conn := mrc.connector.GetConnection()
	return playerstats.NewPlayerStatsRepository(conn)
}

// OpenPlayersRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetPlayersRepository() database.PlayersRepository {
	conn := mrc.connector.GetConnection()
	return players.NewPlayersRepository(conn)
}

// OpenTeamsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetTeamsRepository() database.TeamsRepository {
	conn := mrc.connector.GetConnection()
	return teams.NewTeamsRepository(conn)
}

// OpenUserRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetUserRepository() database.UsersRepositry {
	conn := mrc.connector.GetConnection()
	return users.NewUsersRepositry(conn)
}

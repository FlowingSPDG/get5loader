package mysqlconnector

import (
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/gameservers"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/mapstats"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/matches"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/players"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/playerstats"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/teams"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/users"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type mysqlRepositoryConnector struct {
	uuidGenerator uuid.UUIDGenerator
	connector     database.DBConnector
}

func NewMySQLRepositoryConnector(uuidGenerator uuid.UUIDGenerator, connector database.DBConnector) database.RepositoryConnector {
	return &mysqlRepositoryConnector{
		uuidGenerator: uuidGenerator,
		connector:     connector,
	}
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
	return gameservers.NewGameServerRepository(mrc.uuidGenerator, conn)
}

// OpenMapStatRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetMapStatRepository() database.MapStatRepository {
	conn := mrc.connector.GetConnection()
	return mapstats.NewMapStatRepository(mrc.uuidGenerator, conn)
}

// OpenMatchesRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetMatchesRepository() database.MatchesRepository {
	conn := mrc.connector.GetConnection()
	return matches.NewMatchRepository(mrc.uuidGenerator, conn)
}

// OpenPlayerStatRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetPlayerStatRepository() database.PlayerStatRepository {
	conn := mrc.connector.GetConnection()
	return playerstats.NewPlayerStatRepository(mrc.uuidGenerator, conn)
}

// OpenPlayersRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetPlayersRepository() database.PlayersRepository {
	conn := mrc.connector.GetConnection()
	return players.NewPlayersRepository(mrc.uuidGenerator, conn)
}

// OpenTeamsRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetTeamsRepository() database.TeamsRepository {
	conn := mrc.connector.GetConnection()
	return teams.NewTeamsRepository(mrc.uuidGenerator, conn)
}

// OpenUserRepository implements database.RepositoryConnector.
func (mrc *mysqlRepositoryConnector) GetUserRepository() database.UsersRepositry {
	conn := mrc.connector.GetConnection()
	return users.NewUsersRepositry(mrc.uuidGenerator, conn)
}

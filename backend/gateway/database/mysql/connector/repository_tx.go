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

type mysqlRepositoryConnectorWithTx struct {
	uuidGenerator uuid.UUIDGenerator
	connector     database.DBConnectorWithTx
}

func NewMySQLRepositoryConnectorWithTx(uuidGenerator uuid.UUIDGenerator, connector database.DBConnectorWithTx) database.RepositoryConnectorWithTx {
	return &mysqlRepositoryConnectorWithTx{
		uuidGenerator: uuidGenerator,
		connector:     connector,
	}
}

func (mrctx *mysqlRepositoryConnectorWithTx) Open() error {
	if err := mrctx.connector.Open(); err != nil {
		return err
	}

	if err := mrctx.connector.BeginTx(); err != nil {
		return err
	}

	return nil
}

// Close implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) Close() error {
	return mrctx.connector.Close()
}

// Commit implements database.RepositoryConnectorWithTx.
func (mrctx *mysqlRepositoryConnectorWithTx) Commit() error {
	tx := mrctx.connector.GetTx()
	return tx.Commit()
}

// Rollback implements database.RepositoryConnectorWithTx.
func (mrctx *mysqlRepositoryConnectorWithTx) Rollback() error {
	tx := mrctx.connector.GetTx()
	return tx.Rollback()
}

// OpenGameServersRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetGameServersRepository() database.GameServersRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return gameservers.NewGameServerRepositoryWithTx(mrctx.uuidGenerator, conn, tx)
}

// OpenMapStatRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetMapStatRepository() database.MapStatRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return mapstats.NewMapStatRepositoryWithTx(mrctx.uuidGenerator, conn, tx)
}

// OpenMatchesRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetMatchesRepository() database.MatchesRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return matches.NewMatchRepositoryWithTx(mrctx.uuidGenerator, conn, tx)
}

// OpenPlayerStatRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetPlayerStatRepository() database.PlayerStatRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return playerstats.NewPlayerStatRepositoryWithTx(mrctx.uuidGenerator, conn, tx)
}

// OpenPlayersRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetPlayersRepository() database.PlayersRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return players.NewPlayersRepositoryWithTx(mrctx.uuidGenerator, conn, tx)
}

// OpenTeamsRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetTeamsRepository() database.TeamsRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return teams.NewTeamsRepositoryWithTx(mrctx.uuidGenerator, conn, tx)
}

// OpenUserRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetUserRepository() database.UsersRepositry {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return users.NewUsersRepositryWithTx(mrctx.uuidGenerator, conn, tx)
}

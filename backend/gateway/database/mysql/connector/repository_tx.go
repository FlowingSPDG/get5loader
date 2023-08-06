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

type mysqlRepositoryConnectorWithTx struct {
	connector database.DBConnectorWithTx
}

func NewMySQLRepositoryConnectorWithTx(connector database.DBConnectorWithTx) database.RepositoryConnectorWithTx {
	return &mysqlRepositoryConnectorWithTx{connector: connector}
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
	return gameservers.NewGameServerRepositoryWithTx(conn, tx)
}

// OpenMapStatsRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetMapStatsRepository() database.MapStatsRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return mapstats.NewMapStatsRepositoryWithTx(conn, tx)
}

// OpenMatchesRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetMatchesRepository() database.MatchesRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return matches.NewMatchRepositoryWithTx(conn, tx)
}

// OpenPlayerStatsRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetPlayerStatsRepository() database.PlayerStatsRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return playerstats.NewPlayerStatsRepositoryWithTx(conn, tx)
}

// OpenPlayersRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetPlayersRepository() database.PlayersRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return players.NewPlayersRepositoryWithTx(conn, tx)
}

// OpenTeamsRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetTeamsRepository() database.TeamsRepository {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return teams.NewTeamsRepositoryWithTx(conn, tx)
}

// OpenUserRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetUserRepository() database.UsersRepositry {
	conn := mrctx.connector.GetConnection()
	tx := mrctx.connector.GetTx()
	return users.NewUsersRepositryWithTx(conn, tx)
}

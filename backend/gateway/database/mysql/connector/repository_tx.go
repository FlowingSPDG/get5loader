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

type mysqlRepositoryConnectorWithTx struct {
	connector database.DBConnectorWithTx
	db        *sql.DB
	tx        *sql.Tx
}

func NewMySQLRepositoryConnectorWithTx(connector database.DBConnectorWithTx) database.RepositoryConnectorWithTx {
	return &mysqlRepositoryConnectorWithTx{connector: connector}
}

func (mrctx *mysqlRepositoryConnectorWithTx) Open() error {
	if err := mrctx.connector.Open(); err != nil {
		return err
	}

	mrctx.db = mrctx.connector.GetConnection()
	if err := mrctx.connector.BeginTx(); err != nil {
		return err
	}
	mrctx.tx = mrctx.connector.GetTx()
	return nil
}

// Close implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) Close() error {
	return mrctx.connector.Close()
}

// Commit implements database.RepositoryConnectorWithTx.
func (mrctx *mysqlRepositoryConnectorWithTx) Commit() error {
	return mrctx.tx.Commit()
}

// Rollback implements database.RepositoryConnectorWithTx.
func (mrctx *mysqlRepositoryConnectorWithTx) Rollback() error {
	return mrctx.tx.Rollback()
}

// OpenGameServersRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetGameServersRepository() (database.GameServersRepository, error) {
	repository := gameservers.NewGameServerRepositoryWithTx(mrctx.db, mrctx.tx)
	return repository, nil
}

// OpenMapStatsRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetMapStatsRepository() (database.MapStatsRepository, error) {
	repository := mapstats.NewMapStatsRepositoryWithTx(mrctx.db, mrctx.tx)
	return repository, nil
}

// OpenMatchesRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetMatchesRepository() (database.MatchesRepository, error) {
	repository := matches.NewMatchRepositoryWithTx(mrctx.db, mrctx.tx)
	return repository, nil
}

// OpenPlayerStatsRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetPlayerStatsRepository() (database.PlayerStatsRepository, error) {
	repository := playerstats.NewPlayerStatsRepositoryWithTx(mrctx.db, mrctx.tx)
	return repository, nil
}

// OpenPlayersRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetPlayersRepository() (database.PlayersRepository, error) {
	repository := players.NewPlayersRepositoryWithTx(mrctx.db, mrctx.tx)
	return repository, nil
}

// OpenTeamsRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetTeamsRepository() (database.TeamsRepository, error) {
	repository := teams.NewTeamsRepositoryWithTx(mrctx.db, mrctx.tx)
	return repository, nil
}

// OpenUserRepository implements database.RepositoryConnector.
func (mrctx *mysqlRepositoryConnectorWithTx) GetUserRepository() (database.UsersRepositry, error) {
	repository := users.NewUsersRepositryWithTx(mrctx.db, mrctx.tx)
	return repository, nil
}

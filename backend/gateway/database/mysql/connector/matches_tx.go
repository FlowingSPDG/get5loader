package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/matches"
)

type mysqlMatchesRepositoryConnectorWithTx struct {
	connector database.DBConnectorWithTx
}

func NewMySQLMatchesRepositoryConnectorWithTx(connector database.DBConnectorWithTx) database.RepositoryConnectorWithTx[database.MatchesRepository] {
	return &mysqlMatchesRepositoryConnectorWithTx{connector: connector}
}

// Open MySQL database connection with transaction. transaction starts immediately after opening.
func (mmrctc *mysqlMatchesRepositoryConnectorWithTx) Open() (database.MatchesRepository, error) {
	if err := mmrctc.connector.Open(); err != nil {
		return nil, err
	}

	db := mmrctc.connector.GetConnection()
	if err := mmrctc.connector.BeginTx(); err != nil {
		return nil, err
	}
	tx := mmrctc.connector.GetTx()

	repository := matches.NewMatchRepositoryWithTx(db, tx)

	return repository, nil
}

func (mmrctc *mysqlMatchesRepositoryConnectorWithTx) Close() error {
	return mmrctc.connector.Close()
}

func (mmrctc *mysqlMatchesRepositoryConnectorWithTx) Commit() error {
	return mmrctc.connector.GetTx().Commit()
}

func (mmrctc *mysqlMatchesRepositoryConnectorWithTx) Rollback() error {
	return mmrctc.connector.GetTx().Rollback()
}

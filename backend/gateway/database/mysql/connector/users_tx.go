package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users"
)

type mysqlUsersRepositoryConnectorWithTx struct {
	connector database.DBConnectorWithTx
}

// NewMySQLUsersRepositoryConnectorWithTx returns a new mysqlUsersRepositoryConnectorWithTx.
func NewMySQLUsersRepositoryConnectorWithTx(connector database.DBConnectorWithTx) database.RepositoryConnectorWithTx[database.UsersRepositry] {
	return &mysqlUsersRepositoryConnectorWithTx{connector: connector}
}

// Open MySQL database connection with transaction. transaction starts immediately after opening.
func (murctx *mysqlUsersRepositoryConnectorWithTx) Open() (database.UsersRepositry, error) {
	if err := murctx.connector.Open(); err != nil {
		return nil, err
	}

	db := murctx.connector.GetConnection()
	if err := murctx.connector.BeginTx(); err != nil {
		return nil, err
	}
	tx := murctx.connector.GetTx()

	repository := users.NewUsersRepositryWithTx(db, tx)

	return repository, nil
}

func (murctx *mysqlUsersRepositoryConnectorWithTx) Close() error {
	return murctx.connector.Close()
}

func (murctx *mysqlUsersRepositoryConnectorWithTx) Commit() error {
	return murctx.connector.GetTx().Commit()
}

func (murctx *mysqlUsersRepositoryConnectorWithTx) Rollback() error {
	return murctx.connector.GetTx().Rollback()
}

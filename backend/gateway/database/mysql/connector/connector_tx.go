package mysqlconnector

import (
	"database/sql"

	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
)

type mysqlConnectorWithTx struct {
	mysqlConnector
	tx *sql.Tx // since mysqlConnectorWithTx only stores one transaction, only up to one transaction can be used at a time.
}

func NewMysqlConnectorWithTx(dsn string) database.DBConnectorWithTx {
	return &mysqlConnectorWithTx{mysqlConnector: mysqlConnector{dsn: dsn}}
}

func (mctx *mysqlConnectorWithTx) Open() error {
	return mctx.mysqlConnector.Open()
}

func (mctx *mysqlConnectorWithTx) GetConnection() *sql.DB {
	return mctx.mysqlConnector.GetConnection()
}

// BeginTx implements database.DBConnectorWithTx.
func (mctx *mysqlConnectorWithTx) BeginTx() error {
	tx, err := mctx.db.Begin()
	if err != nil {
		return err
	}
	mctx.tx = tx
	return nil
}

// GetTx implements database.DBConnectorWithTx.
func (mctx *mysqlConnectorWithTx) GetTx() *sql.Tx {
	return mctx.tx
}

func (mctx *mysqlConnectorWithTx) Close() error {
	return mctx.mysqlConnector.Close()
}

// Commit implements database.DBConnectorWithTx.
func (mctx *mysqlConnectorWithTx) Commit() error {
	return mctx.tx.Commit()
}

// Rollback implements database.DBConnectorWithTx.
func (mctx *mysqlConnectorWithTx) Rollback() error {
	return mctx.tx.Rollback()
}

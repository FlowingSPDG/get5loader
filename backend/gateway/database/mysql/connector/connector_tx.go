package mysqlconnector

import (
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
)

type mysqlConnectorWithTx struct {
	dsn string
	db  *sql.DB
	tx  *sql.Tx // since mysqlConnectorWithTx only stores one transaction, only up to one transaction can be used at a time.
}

func NewMysqlConnectorWithTx(dsn string) database.DBConnectorWithTx {
	return &mysqlConnectorWithTx{dsn: dsn}
}

func (mctx *mysqlConnectorWithTx) Open() error {
	// すでに接続済みの場合は何もしない
	if mctx.db != nil {
		return nil
	}

	db, err := sql.Open("mysql", mctx.dsn)
	if err != nil {
		return err
	}
	mctx.db = db

	return nil
}

func (mctx *mysqlConnectorWithTx) GetConnection() *sql.DB {
	return mctx.db
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
	if err := mctx.db.Close(); err != nil {
		return err
	}
	mctx.db = nil
	return nil
}

// Commit implements database.DBConnectorWithTx.
func (mctx *mysqlConnectorWithTx) Commit() error {
	return mctx.tx.Commit()
}

// Rollback implements database.DBConnectorWithTx.
func (mctx *mysqlConnectorWithTx) Rollback() error {
	return mctx.tx.Rollback()
}

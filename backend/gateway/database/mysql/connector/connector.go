package mysqlconnector

import (
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

type mysqlConnector struct {
	dsn string
	db  *sql.DB
}

func NewMysqlConnector(dsn string) database.DBConnector {
	return &mysqlConnector{dsn: dsn}
}

func (mc *mysqlConnector) Open() error {
	// すでに接続済みの場合は何もしない
	if mc.db != nil {
		return nil
	}

	db, err := sql.Open("mysql", mc.dsn)
	if err != nil {
		return err
	}
	mc.db = db

	return nil
}

func (mc *mysqlConnector) GetConnection() *sql.DB {
	return mc.db
}

func (mc *mysqlConnector) Close() error {
	if err := mc.db.Close(); err != nil {
		return err
	}
	mc.db = nil
	return nil
}

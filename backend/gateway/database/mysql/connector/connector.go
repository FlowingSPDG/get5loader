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

// BeginTx implements database.DBConnector.
func (mc *mysqlConnector) BeginTx() (*sql.Tx, error) {
	return mc.db.Begin()
}

func (mc *mysqlConnector) Close() error {
	return mc.db.Close()
}

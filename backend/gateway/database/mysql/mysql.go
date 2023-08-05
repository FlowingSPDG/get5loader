package mysql

import (
	"database/sql"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// Connectorを作成する

type mysqlConnector struct {
	dsn string
}

func NewMysqlConnector(dsn string) database.DBConnector {
	return &mysqlConnector{dsn: dsn}
}

// Connect implements database.DBConnector.
func (mc *mysqlConnector) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", mc.dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

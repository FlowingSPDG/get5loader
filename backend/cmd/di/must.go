package di

import (
	"fmt"

	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	mysqlconnector "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/connector"
)

func mustGetWriteConnector(cfg config.Config) database.DBConnector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4", cfg.DBWriteUser, cfg.DBWritePass, cfg.DBWriteHost, cfg.DBWritePort, cfg.DBWriteName)
	return mysqlconnector.NewMysqlConnector(dsn)
}

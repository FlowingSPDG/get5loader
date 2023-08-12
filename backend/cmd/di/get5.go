package di

import (
	got5 "github.com/FlowingSPDG/Got5"
	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	"github.com/FlowingSPDG/get5loader/backend/controller/get5"
	mysqlconnector "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/connector"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

func InitializeGet5EventController() got5.EventHandler {
	return get5.NewGot5EventController()
}

func InitializeGet5AuthController() got5.Auth {
	return get5.NewGot5AuthController()
}

func InitializeGet5matchLoaderController(cfg config.Config) got5.MatchLoader {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	return get5.NewGot5MatchLoader(mysqlUsersRepositoryConnector)
}

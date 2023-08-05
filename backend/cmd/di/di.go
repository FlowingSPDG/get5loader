package di

import (
	"fmt"

	config "github.com/FlowingSPDG/get5-web-go/backend/cfg"
	"github.com/FlowingSPDG/get5-web-go/backend/controller/gin/api"
	api_controller "github.com/FlowingSPDG/get5-web-go/backend/controller/gin/api"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql"
	api_presenter "github.com/FlowingSPDG/get5-web-go/backend/presenter/gin/api"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

var (
	mapPool = []string{
		"de_inferno",
		"de_mirage",
		"de_nuke",
		"de_overpass",
		"de_vertigo",
		"de_ancient",
		"de_anubis",
	}
)

func mustGetWriteConnector(cfg config.Config) database.DBConnector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4", cfg.DBWriteUser, cfg.DBWritePass, cfg.DBWriteHost, cfg.DBWritePort, cfg.DBWriteName)
	return mysql.NewMysqlConnector(dsn)
}

func InitializeGetMatchController(cfg config.Config) api.GetMatchController {
	mysqlConnector := mustGetWriteConnector(cfg)
	uc := usecase.NewGetMatch(mysqlConnector)
	presenter := api_presenter.NewMatchPresenter()
	return api_controller.NewGetMatchController(uc, presenter)
}

func InitializeGetVersionController() api.GetVersionController {
	return api_controller.NewGetVersionController()
}

func InitializeGetMaplistController() api.GetMaplistController {
	return api_controller.NewGetMaplistController(mapPool, []string{})
}

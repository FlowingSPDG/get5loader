package di

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	config "github.com/FlowingSPDG/get5-web-go/backend/cfg"
	"github.com/FlowingSPDG/get5-web-go/backend/controller/gin/api"
	api_controller "github.com/FlowingSPDG/get5-web-go/backend/controller/gin/api"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	mysqlconnector "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/connector"
	api_presenter "github.com/FlowingSPDG/get5-web-go/backend/presenter/gin/api"
	"github.com/FlowingSPDG/get5-web-go/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5-web-go/backend/service/password_hash"
	"github.com/FlowingSPDG/get5-web-go/backend/service/uuid"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

// TODO: wire

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
	return mysqlconnector.NewMysqlConnector(dsn)
}

func InitializeGetMatchController(cfg config.Config) api.GetMatchController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	uc := usecase.NewGetMatch(mysqlUsersRepositoryConnector)
	presenter := api_presenter.NewMatchPresenter()
	return api_controller.NewGetMatchController(uc, presenter)
}

func InitializeGetVersionController() api.GetVersionController {
	return api_controller.NewGetVersionController()
}

func InitializeGetMaplistController() api.GetMaplistController {
	return api_controller.NewGetMaplistController(mapPool, []string{})
}

func InitializeUserLoginController(cfg config.Config) api.UserLoginController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretMey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUserLogin(jwtService, passwordHasher, mysqlUsersRepositoryConnector)
	presenter := api_presenter.NewJWTPresenter()
	return api_controller.NewUserLoginController(uc, presenter)
}

func InitializeUserRegisterController(cfg config.Config) api.UserRegisterController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretMey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUserRegister(jwtService, passwordHasher, mysqlUsersRepositoryConnector)
	presenter := api_presenter.NewJWTPresenter()
	return api_controller.NewUserRegisterController(uc, presenter)
}

func InitializeJWTAuthController(cfg config.Config) api.JWTAuthController {
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretMey))
	uc := usecase.NewValidateJWT(jwtService)
	return api_controller.NewJWTAuthController(uc)
}

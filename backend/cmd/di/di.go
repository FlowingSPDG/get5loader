package di

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	config "github.com/FlowingSPDG/get5-web-go/backend/cfg"
	gin_controller "github.com/FlowingSPDG/get5-web-go/backend/controller/gin"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	mysqlconnector "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/connector"
	gin_presenter "github.com/FlowingSPDG/get5-web-go/backend/presenter/gin"
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

func InitializeGetMatchController(cfg config.Config) gin_controller.GetMatchController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	uc := usecase.NewGetMatch(mysqlUsersRepositoryConnector)
	presenter := gin_presenter.NewMatchPresenter()
	return gin_controller.NewGetMatchController(uc, presenter)
}

func InitializeGetVersionController() gin_controller.GetVersionController {
	return gin_controller.NewGetVersionController()
}

func InitializeGetMaplistController() gin_controller.GetMaplistController {
	return gin_controller.NewGetMaplistController(mapPool, []string{})
}

func InitializeUserLoginController(cfg config.Config) gin_controller.UserLoginController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretMey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUserLogin(jwtService, passwordHasher, mysqlUsersRepositoryConnector)
	presenter := gin_presenter.NewJWTPresenter()
	return gin_controller.NewUserLoginController(uc, presenter)
}

func InitializeUserRegisterController(cfg config.Config) gin_controller.UserRegisterController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretMey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUserRegister(jwtService, passwordHasher, mysqlUsersRepositoryConnector)
	presenter := gin_presenter.NewJWTPresenter()
	return gin_controller.NewUserRegisterController(uc, presenter)
}

func InitializeJWTAuthController(cfg config.Config) gin_controller.JWTAuthController {
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretMey))
	uc := usecase.NewValidateJWT(jwtService)
	return gin_controller.NewJWTAuthController(uc)
}

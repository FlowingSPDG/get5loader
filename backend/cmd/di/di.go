package di

import (
	"golang.org/x/crypto/bcrypt"

	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	gin_controller "github.com/FlowingSPDG/get5loader/backend/controller/gin"
	mysqlconnector "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/connector"
	gin_presenter "github.com/FlowingSPDG/get5loader/backend/presenter/gin"
	"github.com/FlowingSPDG/get5loader/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5loader/backend/service/password_hash"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

func InitializeUserLoginController(cfg config.Config) gin_controller.UserLoginController {
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretKey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUser(jwtService, passwordHasher)
	presenter := gin_presenter.NewJWTPresenter()
	return gin_controller.NewUserLoginController(uc, presenter)
}

func InitializeUserRegisterController(cfg config.Config) gin_controller.UserRegisterController {
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretKey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUser(jwtService, passwordHasher)
	presenter := gin_presenter.NewJWTPresenter()
	return gin_controller.NewUserRegisterController(uc, presenter)
}

func InitializeDatabaseConnectionMiddleware(cfg config.Config) gin_controller.DatabaseConnectionMiddleware {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	return gin_controller.NewDatabaseConnectionMiddleware(uuidGenerator, mysqlUsersRepositoryConnector)
}

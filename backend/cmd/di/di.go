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

func InitializeJWTAuthController(cfg config.Config) gin_controller.JWTAuthController {
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretKey))
	uc := usecase.NewValidateJWT(jwtService)
	return gin_controller.NewJWTAuthController(uc)
}

func InitializeUserLoginController(cfg config.Config) gin_controller.UserLoginController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretKey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUser(jwtService, passwordHasher, mysqlUsersRepositoryConnector)
	presenter := gin_presenter.NewJWTPresenter()
	return gin_controller.NewUserLoginController(uc, presenter)
}

func InitializeUserRegisterController(cfg config.Config) gin_controller.UserRegisterController {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretKey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	uc := usecase.NewUser(jwtService, passwordHasher, mysqlUsersRepositoryConnector)
	presenter := gin_presenter.NewJWTPresenter()
	return gin_controller.NewUserRegisterController(uc, presenter)
}

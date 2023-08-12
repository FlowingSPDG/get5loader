package di

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	"github.com/FlowingSPDG/get5loader/backend/g5ctx"
	mysqlconnector "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/connector"
	"github.com/FlowingSPDG/get5loader/backend/graph"
	"github.com/FlowingSPDG/get5loader/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5loader/backend/service/password_hash"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

func InitializeGraphQLHandler(cfg config.Config) gin.HandlerFunc {
	uuidGenerator := uuid.NewUUIDGenerator()
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretKey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	gameServerUc := usecase.NewGameServer(mysqlUsersRepositoryConnector)
	matchUc := usecase.NewMatch(mysqlUsersRepositoryConnector)
	userUc := usecase.NewUser(jwtService, passwordHasher, mysqlUsersRepositoryConnector)

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		GameServerUsecase: gameServerUc,
		MatchUsecase:      matchUc,
		UserUsecase:       userUc,
	}}))

	return func(c *gin.Context) {
		g5ctx.SetOperationGinContext(c, g5ctx.OperationTypeUser)
		h.ServeHTTP(c.Writer, c.Request)
	}
}

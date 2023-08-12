package di

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	mysqlconnector "github.com/FlowingSPDG/get5loader/backend/gateway/database/mysql/connector"
	"github.com/FlowingSPDG/get5loader/backend/graph"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

func InitializeGraphQLHandler(cfg config.Config) gin.HandlerFunc {
	uuidGenerator := uuid.NewUUIDGenerator()
	mysqlConnector := mustGetWriteConnector(cfg)
	mysqlUsersRepositoryConnector := mysqlconnector.NewMySQLRepositoryConnector(uuidGenerator, mysqlConnector)
	uc := usecase.NewGameServer(mysqlUsersRepositoryConnector)
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		GameServerUsecase: uc,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

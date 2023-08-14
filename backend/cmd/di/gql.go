package di

import (
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	"github.com/FlowingSPDG/get5loader/backend/g5ctx"
	"github.com/FlowingSPDG/get5loader/backend/graph"
	"github.com/FlowingSPDG/get5loader/backend/service/jwt"
	hash "github.com/FlowingSPDG/get5loader/backend/service/password_hash"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

func InitializeGraphQLHandler(cfg config.Config) gin.HandlerFunc {
	// dependencies
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretKey))
	passwordHasher := hash.NewPasswordHasher(bcrypt.DefaultCost)

	// usecases
	gameServerUc := usecase.NewGameServer()
	userUc := usecase.NewUser(jwtService, passwordHasher)
	matchUc := usecase.NewMatch()
	mapStatUc := usecase.NewMapStats()
	teamUc := usecase.NewTeam()
	playerUc := usecase.NewPlayer()

	// handler
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		GameServerUsecase: gameServerUc,
		UserUsecase:       userUc,
		MatchUsecase:      matchUc,
		MapstatUsecase:    mapStatUc,
		TeamUsecase:       teamUc,
		PlayerUsecase:     playerUc,
	}}))

	// middleware
	validateJWTusecase := usecase.NewValidateJWT(jwtService)
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		authorization, _ = strings.CutPrefix(authorization, "Bearer ")
		token, err := validateJWTusecase.Validate(authorization)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		// contextにuserIDを入れる
		g5ctx.SetUserTokenGinContext(c, token)
		g5ctx.SetOperationGinContext(c, g5ctx.OperationTypeUser)
		c.Request = c.Request.WithContext(c)
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func InitializePlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

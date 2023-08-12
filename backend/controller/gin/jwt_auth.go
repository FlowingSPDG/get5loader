package gin_controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5loader/backend/g5ctx"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

type JWTAuthController interface {
	Handle(c *gin.Context)
}

type jwtAuthController struct {
	uc usecase.ValidateJWT
}

// Handle implements JWTAuthController.
func (jac *jwtAuthController) Handle(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	authorization, _ = strings.CutPrefix(authorization, "Bearer ")
	token, err := jac.uc.Validate(authorization)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	// contextにuserIDを入れる
	g5ctx.SetUserToken(c, token)
	g5ctx.SetOperation(c, g5ctx.OperationTypeUser)
	c.Next()
}

func NewJWTAuthController(
	uc usecase.ValidateJWT,
) JWTAuthController {
	return &jwtAuthController{
		uc: uc,
	}
}

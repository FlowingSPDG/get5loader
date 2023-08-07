package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

type JWTAuthController interface {
	Handle(c *gin.Context)
}

type jwtAuthController struct {
	adminOnly bool
	uc        usecase.ValidateJWT
}

// Handle implements JWTAuthController.
func (jac *jwtAuthController) Handle(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	authorization, _ = strings.CutPrefix(authorization, "Bearer ")
	admin, err := jac.uc.Validate(authorization)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	if jac.adminOnly && !admin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "you are not admin",
		})
		return
	}

	// contextにadmin,userIDを入れる
	g5ctx.
		c.Next()
}

func NewJWTAuthController(
	adminOnly bool,
	uc usecase.ValidateJWT,
) JWTAuthController {
	return &jwtAuthController{
		adminOnly: adminOnly,
		uc:        uc,
	}
}

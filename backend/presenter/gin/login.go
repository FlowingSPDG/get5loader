package gin_presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JWTPresenter interface {
	Handle(c *gin.Context, token string)
}

type jwtPresenter struct{}

func NewJWTPresenter() JWTPresenter {
	return &jwtPresenter{}
}

type jwtResponse struct {
	Token string `json:"token"`
}

// Handle implements JWTPresenter.
func (jp *jwtPresenter) Handle(c *gin.Context, token string) {
	c.JSON(http.StatusOK, &jwtResponse{
		Token: token,
	})
}

package gin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	gin_presenter "github.com/FlowingSPDG/get5-web-go/backend/presenter/gin"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

type UserLoginController interface {
	Handle(c *gin.Context)
}

type userLoginController struct {
	uc        usecase.UserLogin
	presenter gin_presenter.JWTPresenter
}

func NewUserLoginController(
	uc usecase.UserLogin,
	presenter gin_presenter.JWTPresenter,
) UserLoginController {
	return &userLoginController{
		uc:        uc,
		presenter: presenter,
	}
}

type userLoginRequest struct {
	SteamID  entity.SteamID `json:"steamid"`
	Password string         `json:"password"`
}

// Handle implements UserLoginController.
func (ulc *userLoginController) Handle(c *gin.Context) {
	var req userLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	jwt, err := ulc.uc.IssueJWTBySteamID(c, req.SteamID, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	ulc.presenter.Handle(c, jwt)
}
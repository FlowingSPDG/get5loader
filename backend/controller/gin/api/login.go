package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5-web-go/backend/presenter/gin/api"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

type UserLoginController interface {
	Handle(c *gin.Context)
}

type userLoginController struct {
	uc        usecase.UserLogin
	presenter api.JWTPresenter
}

func NewUserLoginController(
	uc usecase.UserLogin,
	presenter api.JWTPresenter,
) UserLoginController {
	return &userLoginController{
		uc:        uc,
		presenter: presenter,
	}
}

type userLoginRequest struct {
	SteamID  string `json:"steamid"`
	Password string `json:"password"`
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

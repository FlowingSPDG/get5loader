package gin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	gin_presenter "github.com/FlowingSPDG/get5-web-go/backend/presenter/gin"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

type UserRegisterController interface {
	Handle(c *gin.Context)
}

type userRegisterController struct {
	uc        usecase.UserRegister
	presenter gin_presenter.JWTPresenter
}

func NewUserRegisterController(
	uc usecase.UserRegister,
	presenter gin_presenter.JWTPresenter,
) UserRegisterController {
	return &userRegisterController{
		uc:        uc,
		presenter: presenter,
	}
}

type userRegisterRequest struct {
	Name     string         `json:"name"`
	SteamID  entity.SteamID `json:"steamid"`
	Password string         `json:"password"`
}

// Handle implements UserLoginController.
func (urc *userRegisterController) Handle(c *gin.Context) {
	var req userRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	jwt, err := urc.uc.RegisterUser(c, req.SteamID, req.Name, false, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	urc.presenter.Handle(c, jwt)
}

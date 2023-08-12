package di

import (
	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	gin_controller "github.com/FlowingSPDG/get5loader/backend/controller/gin"
	"github.com/FlowingSPDG/get5loader/backend/service/jwt"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

func InitializeJWTAuthController(cfg config.Config) gin_controller.JWTAuthController {
	jwtService := jwt.NewJWTGateway([]byte(cfg.SecretMey))
	uc := usecase.NewValidateJWT(jwtService)
	return gin_controller.NewJWTAuthController(uc)
}

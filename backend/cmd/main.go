package main

import (
	"fmt"
	"log"

	route "github.com/FlowingSPDG/Got5/route/gin"
	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	"github.com/FlowingSPDG/get5loader/backend/cmd/di"
	"github.com/FlowingSPDG/get5loader/backend/controller/got5"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.GetConfig()

	v1 := r.Group("/api/v1")
	v1.POST("/login", di.InitializeUserLoginController(cfg).Handle)
	v1.POST("/register", di.InitializeUserRegisterController(cfg).Handle)

	v1auth := v1.Group("/")
	v1auth.Use(di.InitializeJWTAuthController(cfg).Handle)
	v1auth.POST("/query", di.InitializeGraphQLHandler(cfg))

	g5 := v1.Group("/get5_event")
	evh := got5.NewGot5EventController()
	ah := got5.NewGot5AuthController()
	route.SetupEventHandlers(evh, ah, g5)

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Panicf("Failed to listen port %v", r.Run(addr))
}

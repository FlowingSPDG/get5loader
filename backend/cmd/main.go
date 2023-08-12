package main

import (
	"fmt"
	"log"

	route "github.com/FlowingSPDG/Got5/route/gin"
	config "github.com/FlowingSPDG/get5loader/backend/cfg"
	"github.com/FlowingSPDG/get5loader/backend/cmd/di"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.GetConfig()

	v1 := r.Group("/api/v1")
	v1.POST("/login", di.InitializeUserLoginController(cfg).Handle)
	v1.POST("/register", di.InitializeUserRegisterController(cfg).Handle)

	v1.POST("/query", di.InitializeGraphQLHandler(cfg))
	v1.GET("/playground", di.InitializePlaygroundHandler())

	g5 := v1.Group("/get5")
	ec := di.InitializeGet5EventController()
	ac := di.InitializeGet5AuthController()
	ml := di.InitializeGet5matchLoaderController(cfg)
	route.SetupEventHandlers(ec, ac, g5)
	route.SetupMatchLoadHandler(ml, ac, g5)

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Panicf("Failed to listen port %v", r.Run(addr))
}

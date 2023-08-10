package gin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMaplistController interface {
	Handle(c *gin.Context)
}

type getMaplistController struct {
	activeMapPool  []string
	reserveMapPool []string
}

func NewGetMaplistController(activeMapPool, reserveMapPool []string) GetMaplistController {
	return &getMaplistController{
		activeMapPool:  activeMapPool,
		reserveMapPool: reserveMapPool,
	}
}

func (gmlc *getMaplistController) Handle(c *gin.Context) {
	// TODO: usecaseから取得するか検討する
	c.JSON(http.StatusOK, gin.H{
		"active":  gmlc.activeMapPool,
		"reserve": gmlc.reserveMapPool,
	})
}

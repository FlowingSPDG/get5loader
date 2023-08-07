package gin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// VERSION is the current version of the get5-web-go.
	VERSION = "0.2.0"
)

type GetVersionController interface {
	Handle(c *gin.Context)
}

type getVersionController struct {
}

func (gvc *getVersionController) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": VERSION,
	})
}

func NewGetVersionController() GetVersionController {
	return &getVersionController{}
}

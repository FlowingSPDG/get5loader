package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/presenter/gin/api"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

type GetMatchController interface {
	Handle(c *gin.Context)
}

type getMatchController struct {
	uc        usecase.GetMatch
	presenter api.MatchPresenter
}

func NewGetMatchController(
	uc usecase.GetMatch,
	presenter api.MatchPresenter,
) GetMatchController {
	return &getMatchController{
		uc:        uc,
		presenter: presenter,
	}
}

// Handle implements GetMatchController.
func (gmc *getMatchController) Handle(c *gin.Context) {
	matchidStr := c.Params.ByName("matchID")
	matchid, err := strconv.Atoi(matchidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "matchID is not int",
		})
		return
	}

	match, err := gmc.uc.Handle(c, int64(matchid))
	if err != nil {
		if database.IsNotFound(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "match not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	b, err := gmc.presenter.Handle(match)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, b)
}

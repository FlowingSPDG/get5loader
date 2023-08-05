package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

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
		c.JSON(400, gin.H{
			"error": "matchID is not int",
		})
		return
	}

	match, err := gmc.uc.Handle(c, int64(matchid))
	if err != nil {
		// TODO: エラーハンドリング
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := gmc.presenter.Handle(match)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

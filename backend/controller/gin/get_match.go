package gin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	gin_presenter "github.com/FlowingSPDG/get5-web-go/backend/presenter/gin"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

type GetMatchController interface {
	Handle(c *gin.Context)
}

type getMatchController struct {
	uc        usecase.GetMatch
	presenter gin_presenter.MatchPresenter
}

func NewGetMatchController(
	uc usecase.GetMatch,
	presenter gin_presenter.MatchPresenter,
) GetMatchController {
	return &getMatchController{
		uc:        uc,
		presenter: presenter,
	}
}

// Handle implements GetMatchController.
func (gmc *getMatchController) Handle(c *gin.Context) {
	matchid := c.Params.ByName("matchID")

	match, err := gmc.uc.Get(c, entity.MatchID(matchid))
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

	gmc.presenter.Handle(c, match)
}

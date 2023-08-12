package gin_presenter

import (
	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5loader/backend/entity"
)

// 受け取ったentityを特定の形式に変換して返す

type MatchPresenter interface {
	Handle(c *gin.Context, m *entity.Match)
}
type matchPresenter struct {
}

func NewMatchPresenter() MatchPresenter {
	return &matchPresenter{}
}

// Handle implements MatchPresenter.
func (mp *matchPresenter) Handle(c *gin.Context, m *entity.Match) {
	data := match{
		ID:         m.ID,
		UserID:     m.UserID,
		Team1ID:    m.Team1ID,
		Team2ID:    m.Team2ID,
		Winner:     m.Winner,
		Cancelled:  m.Status == entity.MATCH_STATUS_CANCELLED,
		StartTime:  m.StartTime,
		EndTime:    m.EndTime,
		MaxMaps:    int(m.MaxMaps),
		Title:      m.Title,
		SkipVeto:   m.SkipVeto,
		Team1Score: int(m.Team1Score),
		Team2Score: int(m.Team2Score),
		Forfeit:    *m.Forfeit,
		ServerID:   m.ServerID,
		Status:     m.Status.String(),
	}
	c.JSON(200, data)
}

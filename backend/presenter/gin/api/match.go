package api

import (
	"encoding/json"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
)

// 受け取ったentityを特定の形式に変換して返す

type MatchPresenter interface {
	Handle(m *entity.Match) ([]byte, error)
}
type matchPresenter struct {
}

func NewMatchPresenter() MatchPresenter {
	return &matchPresenter{}
}

// Handle implements MatchPresenter.
func (mp *matchPresenter) Handle(m *entity.Match) ([]byte, error) {
	data := match{
		ID:         int(m.ID),
		UserID:     int(m.UserID),
		Team1ID:    int(m.Team1ID),
		Team2ID:    int(m.Team2ID),
		Winner:     int(*m.Winner),
		Cancelled:  m.Status == entity.MATCH_STATUS_CANCELLED,
		StartTime:  m.StartTime,
		EndTime:    m.EndTime,
		MaxMaps:    int(m.MaxMaps),
		Title:      m.Title,
		SkipVeto:   m.SkipVeto,
		Team1Score: int(m.Team1Score),
		Team2Score: int(m.Team2Score),
		Forfeit:    *m.Forfeit,
		ServerID:   int(m.ServerID),
		Status:     m.Status.String(),
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

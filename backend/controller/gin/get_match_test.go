package gin_controller_test

// controllerをpresenterのテストを両方行っている

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	gin_controller "github.com/FlowingSPDG/get5-web-go/backend/controller/gin"
	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	gin_presenter "github.com/FlowingSPDG/get5-web-go/backend/presenter/gin"
	mock_usecase "github.com/FlowingSPDG/get5-web-go/backend/usecase/mock"
)

func TestGetMatch(t *testing.T) {
	tt := []struct {
		name     string
		input    entity.MatchID
		expected struct {
			code  int
			body  string
			err   error
			match *entity.Match
		}
	}{
		{
			name:  "success",
			input: "1",
			expected: struct {
				code  int
				body  string
				err   error
				match *entity.Match
			}{
				code: http.StatusOK,
				body: `{"id":"1","user_id":"","team1":"","team2":"","winner":null,"server_id":"","cancelled":false,"start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z","max_maps":0,"title":"","skip_veto":false,"team1_score":0,"team2_score":0,"forfeit":false,"map_stats":null,"status":"MATCH_STATUS_UNKNOWN"}`,
				err:  nil,
				match: &entity.Match{
					ID:         "1",
					UserID:     "",
					ServerID:   "",
					Team1ID:    "",
					Team2ID:    "",
					Winner:     nil,
					StartTime:  &time.Time{},
					EndTime:    &time.Time{},
					MaxMaps:    0,
					Title:      "",
					SkipVeto:   false,
					APIKey:     "",
					Team1Score: 0,
					Team2Score: 0,
					Forfeit:    new(bool),
					Status:     0,
				},
			},
		},
	}

	for _, tc := range tt {
		// mock controllerの作成
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// mock usecaseの作成
		mockGetMatch := mock_usecase.NewMockGetMatch(ctrl)
		mockGetMatch.EXPECT().Get(gomock.Any(), tc.input).Return(tc.expected.match, tc.expected.err)

		// presenterの作成
		presenter := gin_presenter.NewMatchPresenter()

		controller := gin_controller.NewGetMatchController(mockGetMatch, presenter)
		t.Run(tc.name, func(t *testing.T) {
			router := setupGET("/api/match/:matchID", controller.Handle)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/match/%s", tc.input), nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expected.code, w.Code)
			assert.JSONEq(t, tc.expected.body, w.Body.String())
		})
	}
}

package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/g5ctx"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mock/connector"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mock/matches"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

func TestGetMatch(t *testing.T) {
	tt := []struct {
		name     string
		input    entity.MatchID
		expected *entity.Match
		err      error
	}{
		{
			name:  "success",
			input: "1",
			expected: &entity.Match{
				ID:         "",
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
			err: nil,
		},
		{
			name:     "not found",
			input:    "2",
			expected: nil,
			err:      database.ErrNotFound,
		},
	}

	// mock connectorの作成
	mockConnector := connector.NewMockRepositoryConnector(
		nil,
		nil,
		nil,
		matches.NewMockMatchesRepository(
			map[entity.MatchID]*entity.Match{
				"1": {
					ID:         "",
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
			map[entity.UserID][]*entity.Match{},
			map[entity.TeamID][]*entity.Match{},
			map[entity.TeamID][]*entity.Match{},
		),
		nil,
		nil,
		nil,
	)
	uc := usecase.NewGetMatch(mockConnector)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctx = g5ctx.SetOperation(ctx, g5ctx.OperationTypeUser)

			actual, err := uc.Handle(ctx, tc.input)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}

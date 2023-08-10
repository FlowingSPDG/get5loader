package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/g5ctx"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	mock_database "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mock"
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
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctx = g5ctx.SetOperation(ctx, g5ctx.OperationTypeUser)

			// mock controllerの作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// mock connectorの作成
			mockConnector := mock_database.NewMockRepositoryConnector(ctrl)
			mockConnector.EXPECT().Open().Return(nil)
			mockConnector.EXPECT().Close().Return(nil)

			// mock MatchesRepositoryの作成
			mockMatchesRepository := mock_database.NewMockMatchesRepository(ctrl)
			mockMatchesRepository.EXPECT().GetMatch(ctx, tc.input).Return(tc.expected, tc.err)
			mockConnector.EXPECT().GetMatchesRepository().Return(mockMatchesRepository)

			// テストの実行とassert
			uc := usecase.NewGetMatch(mockConnector)
			actual, err := uc.Get(ctx, tc.input)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}

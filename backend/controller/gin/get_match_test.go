package gin_controller_test

import (
	"testing"

	gin_controller "github.com/FlowingSPDG/get5-web-go/backend/controller/gin"
	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	mock_usecase "github.com/FlowingSPDG/get5-web-go/backend/usecase/mock"
	"go.uber.org/mock/gomock"
)

func TestGetMatch(t *testing.T) {
	tt := []struct {
		name     string
		input    entity.MatchID
		expected *entity.Match
		err      error
	}{
		{
			name:     "success",
			input:    "1",
			expected: &entity.Match{},
			err:      nil,
		},
	}

	for _, tc := range tt {
		// mock controllerの作成
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGetMatch := mock_usecase.NewMockGetMatch(ctrl)
		controller := gin_controller.NewGetMatchController(mockGetMatch)
		t.Run(tc.name, func(t *testing.T) {
			setupGET("/api/match/:matchID", controller.Handle)
		})
	}
}

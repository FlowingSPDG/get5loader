package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/g5ctx"
	mock_jwt "github.com/FlowingSPDG/get5loader/backend/service/jwt/mock"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestValidateJWT(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected *entity.TokenUser
		err      error
	}{
		{
			name:  "success",
			input: "test",
			expected: &entity.TokenUser{
				UserID:  "1",
				SteamID: 76561198072054549,
				Admin:   true,
			},
			err: nil,
		},
		{
			name:     "invalid token",
			input:    "invalid",
			expected: nil,
			err:      errors.New("invalid token"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctx = g5ctx.SetOperation(ctx, g5ctx.OperationTypeUser)

			// mock connectorの作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// mock JWTServiceの作成
			mockJwtService := mock_jwt.NewMockJWTService(ctrl)
			mockJwtService.EXPECT().ValidateJWT(tc.input).Return(tc.expected, tc.err)

			uc := usecase.NewValidateJWT(mockJwtService)
			actual, err := uc.Validate(tc.input)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}

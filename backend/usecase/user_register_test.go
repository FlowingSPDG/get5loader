package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/g5ctx"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	mock_database "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mock"
	mock_jwt "github.com/FlowingSPDG/get5-web-go/backend/service/jwt/mock"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
)

func TestRegisterUser(t *testing.T) {
	tt := []struct {
		name  string
		input struct {
			steamid  entity.SteamID
			name     string
			admin    bool
			password string
		}
		expected struct {
			jwt  string
			user *entity.User
		}
		err error
	}{
		{
			name: "success",
			input: struct {
				steamid  entity.SteamID
				name     string
				admin    bool
				password string
			}{
				steamid:  76561198072054549,
				name:     "test",
				admin:    true,
				password: "test",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctx = g5ctx.SetOperation(ctx, g5ctx.OperationTypeUser)

			// mock connectorの作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// mockの作成
			mockConnector := mock_database.NewMockRepositoryConnector(ctrl)
			mockConnector.EXPECT().Open().Return(nil)
			mockConnector.EXPECT().Close().Return(nil)
			mockUsersRepository := mock_database.NewMockUsersRepositry(ctrl)
			mockUsersRepository.EXPECT().GetUserBySteamID(ctx, tc.input.steamid).Return(nil, database.ErrNotFound).Times(1)
			mockUsersRepository.EXPECT().CreateUser(ctx, tc.input.steamid, tc.input.name, tc.input.admin, gomock.Any()).Return(nil)
			mockUsersRepository.EXPECT().GetUserBySteamID(ctx, tc.input.steamid).Return(tc.expected.user, nil).Times(1)
			mockConnector.EXPECT().GetUserRepository().Return(mockUsersRepository)
			mockJwtService := mock_jwt.NewMockJWTService(ctrl)
			mockJwtService.EXPECT().IssueJWT(tc.expected.user).Return(tc.expected.jwt, nil)

			uc := usecase.NewUserRegister(mockJwtService, mockConnector)
			jwt, err := uc.RegisterUser(ctx, tc.input.steamid, tc.input.name, tc.input.admin, tc.input.password)
			assert.Equal(t, tc.expected.jwt, jwt)
			assert.Equal(t, tc.err, err)
		})
	}
}

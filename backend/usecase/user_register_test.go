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
	mock_hash "github.com/FlowingSPDG/get5-web-go/backend/service/password_hash/mock"
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
			hash []byte
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
			expected: struct {
				jwt  string
				hash []byte
				user *entity.User
			}{
				jwt:  "test",
				hash: []byte{},
				user: &entity.User{
					ID:      "test",
					Name:    "test",
					SteamID: 76561198072054549,
					Admin:   true,
					Hash:    nil,
				},
			},
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

			// mock UsersRepositoryの作成
			mockUsersRepository := mock_database.NewMockUsersRepositry(ctrl)
			mockUsersRepository.EXPECT().GetUserBySteamID(gomock.Any(), tc.input.steamid).Return(nil, database.ErrNotFound).Times(1)
			mockUsersRepository.EXPECT().CreateUser(gomock.Any(), tc.input.steamid, tc.input.name, tc.input.admin, gomock.Any()).Return(nil)
			mockUsersRepository.EXPECT().GetUserBySteamID(gomock.Any(), tc.input.steamid).Return(tc.expected.user, nil).Times(1)
			mockConnector.EXPECT().GetUserRepository().Return(mockUsersRepository)

			// mock JWTServiceの作成
			mockJwtService := mock_jwt.NewMockJWTService(ctrl)
			mockJwtService.EXPECT().IssueJWT(tc.expected.user).Return(tc.expected.jwt, nil)

			// mock PasswordHasherの作成
			mockPasswordHasher := mock_hash.NewMockPasswordHasher(ctrl)
			mockPasswordHasher.EXPECT().Hash(tc.input.password).Return(tc.expected.hash, nil)

			// テストの実行とassert
			uc := usecase.NewUserRegister(mockJwtService, mockPasswordHasher, mockConnector)
			jwt, err := uc.RegisterUser(ctx, tc.input.steamid, tc.input.name, tc.input.admin, tc.input.password)
			assert.Equal(t, tc.expected.jwt, jwt)
			assert.Equal(t, tc.err, err)
		})
	}
}
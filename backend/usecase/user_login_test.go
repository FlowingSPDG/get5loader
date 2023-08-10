package usecase_test

import (
	"context"
	"testing"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	mock_database "github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mock"
	mock_jwt "github.com/FlowingSPDG/get5-web-go/backend/service/jwt/mock"
	mock_hash "github.com/FlowingSPDG/get5-web-go/backend/service/password_hash/mock"
	"github.com/FlowingSPDG/get5-web-go/backend/usecase"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIssueJWTBySteamID(t *testing.T) {
	tt := []struct {
		name  string
		input struct {
			steamid  entity.SteamID
			password string
		}
		expected struct {
			jwt   string
			err   error
			user  *entity.User
			token *entity.TokenUser
		}
	}{
		{
			name: "success",
			input: struct {
				steamid  entity.SteamID
				password string
			}{
				steamid:  76561198072054549,
				password: "test",
			},
			expected: struct {
				jwt   string
				err   error
				user  *entity.User
				token *entity.TokenUser
			}{
				jwt: "test",
				err: nil,
				user: &entity.User{
					ID:      "test",
					Name:    "test",
					SteamID: 76561198072054549,
					Admin:   true,
					Hash:    nil,
				},
				token: &entity.TokenUser{
					UserID:  "test",
					SteamID: 76561198072054549,
					Admin:   true,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			// mock controllerの作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// mock connectorの作成
			mockConnector := mock_database.NewMockRepositoryConnector(ctrl)
			mockConnector.EXPECT().Open().Return(nil)
			mockConnector.EXPECT().Close().Return(nil)

			// mock UsersRepositoryの作成
			mockUsersRepository := mock_database.NewMockUsersRepositry(ctrl)
			mockUsersRepository.EXPECT().GetUserBySteamID(gomock.Any(), tc.input.steamid).Return(tc.expected.user, nil)
			mockConnector.EXPECT().GetUserRepository().Return(mockUsersRepository)

			// mock JWTServiceの作成
			mockJwtService := mock_jwt.NewMockJWTService(ctrl)
			mockJwtService.EXPECT().IssueJWT(tc.expected.user).Return(tc.expected.jwt, nil)

			// mock PasswordHasherの作成
			mockPasswordHasher := mock_hash.NewMockPasswordHasher(ctrl)
			mockPasswordHasher.EXPECT().Compare(tc.expected.user.Hash, tc.input.password).Return(nil)

			uc := usecase.NewUserLogin(mockJwtService, mockPasswordHasher, mockConnector)
			actual, err := uc.IssueJWTBySteamID(ctx, tc.input.steamid, tc.input.password)
			assert.Equal(t, tc.expected.jwt, actual)
			assert.Equal(t, tc.expected.err, err)
		})
	}

}

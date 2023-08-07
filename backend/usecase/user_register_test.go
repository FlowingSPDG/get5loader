package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FlowingSPDG/get5-web-go/backend/entity"
	"github.com/FlowingSPDG/get5-web-go/backend/g5ctx"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mock/connector"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mock/users"
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
		jwt string
		err error
	}{
		{},
	}

	// mock connectorの作成
	mockConnector := connector.NewMockRepositoryConnector(
		users.NewMockUsersRepositry(
			map[entity.UserID]*entity.User{},
			map[entity.SteamID]*entity.User{},
		),
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	uc := usecase.NewUserRegister(nil, mockConnector)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			ctx = g5ctx.SetOperation(ctx, g5ctx.OperationTypeUser)

			jwt, err := uc.RegisterUser(ctx, tc.input.steamid, tc.input.name, tc.input.admin, tc.input.password)
			assert.Equal(t, tc.jwt, jwt)
			assert.Equal(t, tc.err, err)
		})
	}
}

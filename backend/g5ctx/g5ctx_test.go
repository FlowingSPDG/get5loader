package g5ctx_test

import (
	"context"
	"testing"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/g5ctx"
	"github.com/stretchr/testify/assert"
)

func TestSetOperation(t *testing.T) {
	tc := []struct {
		name      string
		set       bool
		operation g5ctx.OperationType
		err       error
	}{
		{
			name:      "system",
			set:       true,
			operation: g5ctx.OperationTypeSystem,
			err:       nil,
		},
		{
			name:      "user",
			set:       true,
			operation: g5ctx.OperationTypeUser,
			err:       nil,
		},
		{
			name:      "not set",
			set:       false,
			operation: g5ctx.OperationTypeUnknown,
			err:       g5ctx.ErrContextValueNotFound,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			if c.set {
				ctx = g5ctx.SetOperation(ctx, c.operation)
			}
			op, err := g5ctx.GetOperation(ctx)
			assert.Equal(t, c.operation, op)
			assert.Equal(t, c.err, err)
		})
	}
}

func TestSetUser(t *testing.T) {
	tc := []struct {
		name      string
		set       bool
		tokenUser *entity.TokenUser
		err       error
	}{
		{
			name: "set tokenUser 1",
			set:  true,
			tokenUser: &entity.TokenUser{
				UserID:  "1",
				SteamID: 76561198072054549,
				Admin:   true,
			},
			err: nil,
		},
		{
			name: "set userID 2",
			set:  true,
			tokenUser: &entity.TokenUser{
				UserID:  "2",
				SteamID: 76561198072054549,
				Admin:   false,
			},
			err: nil,
		},
		{
			name:      "not set",
			set:       false,
			tokenUser: nil,
			err:       g5ctx.ErrContextValueNotFound,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			if c.set {
				ctx = g5ctx.SetUserToken(ctx, c.tokenUser)
			}
			user, err := g5ctx.GetUserToken(ctx)
			assert.Equal(t, c.tokenUser, user)
			assert.Equal(t, c.err, err)
		})
	}
}

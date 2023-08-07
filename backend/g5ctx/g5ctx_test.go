package g5ctx_test

import (
	"context"
	"testing"

	"github.com/FlowingSPDG/get5-web-go/backend/g5ctx"
	"github.com/stretchr/testify/assert"
)

func TestSetAdmin(t *testing.T) {
	tc := []struct {
		name  string
		set   bool
		admin bool
	}{
		{
			name:  "admin",
			set:   true,
			admin: true,
		},
		{
			name:  "not admin",
			set:   true,
			admin: false,
		},
		{
			name:  "not set",
			set:   false,
			admin: false,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			if c.set {
				ctx = g5ctx.SetAdmin(ctx, c.admin)
			}
			isAdmin := g5ctx.GetAdmin(ctx)
			assert.Equal(t, c.admin, isAdmin)
		})
	}
}

func TestSetUserID(t *testing.T) {
	tc := []struct {
		name   string
		set    bool
		userID int
	}{
		{
			name:   "set userID 1",
			set:    true,
			userID: 1,
		},
		{
			name:   "set userID 2",
			set:    true,
			userID: 2,
		},
		{
			name:   "not set",
			set:    false,
			userID: 0,
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			if c.set {
				ctx = g5ctx.SetUserID(ctx, c.userID)
			}
			userID := g5ctx.GetUserID(ctx)
			assert.Equal(t, c.userID, userID)
		})
	}
}

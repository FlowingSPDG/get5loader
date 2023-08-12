package get5

import (
	"context"

	got5 "github.com/FlowingSPDG/Got5"
)

type authController struct {
}

func NewGot5AuthController() got5.Auth {
	return &authController{}
}

// CheckDemoAuth implements controller.Auth.
func (ac *authController) CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID string, auth string) error {
	return nil
}

// EventAuth implements controller.Auth.
func (ac *authController) EventAuth(ctx context.Context, serverID string, auth string) error {
	// TODO: 認証を実行するusecaseを呼び出す
	return nil
}

// MatchAuth implements controller.Auth.
func (ac *authController) MatchAuth(ctx context.Context, mid string, auth string) error {
	// TODO: 認証を実行するusecaseを呼び出す
	return nil
}

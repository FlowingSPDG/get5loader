package got5

import (
	"context"

	g5controller "github.com/FlowingSPDG/Got5/controller"
)

// interfaceを満たしているかどうか確認する
var _ g5controller.Auth = (*authController)(nil)

type authController struct {
}

func NewGot5AuthController() g5controller.Auth {
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
	return nil
}

package graph

import (
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GameServerUsecase usecase.GameServer
	UserUsecase       usecase.User
	MatchUsecase      usecase.Match
	MapstatUsecase    usecase.Mapstat
	TeamUsecase       usecase.Team
	PlayerUsecase     usecase.Player
}

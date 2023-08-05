package got5

import (
	"context"

	g5controller "github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

// interfaceを満たしているかどうか確認する
var _ g5controller.EventHandler = (*eventController)(nil)

type eventController struct {
}

func NewGot5EventController() g5controller.EventHandler {
	return &eventController{}
}

// Close implements controller.EventHandler.
func (*eventController) Close() error {
	panic("unimplemented")
}

// HandleOnBackupRestore implements controller.EventHandler.
func (*eventController) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	panic("unimplemented")
}

// HandleOnBombDefused implements controller.EventHandler.
func (*eventController) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	panic("unimplemented")
}

// HandleOnBombExploded implements controller.EventHandler.
func (*eventController) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	panic("unimplemented")
}

// HandleOnBombPlanted implements controller.EventHandler.
func (*eventController) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	panic("unimplemented")
}

// HandleOnDecoyStarted implements controller.EventHandler.
func (*eventController) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	panic("unimplemented")
}

// HandleOnDemoFinished implements controller.EventHandler.
func (*eventController) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	panic("unimplemented")
}

// HandleOnDemoUploadEnded implements controller.EventHandler.
func (*eventController) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	panic("unimplemented")
}

// HandleOnFlashbangDetonated implements controller.EventHandler.
func (*eventController) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnGameStateChanged implements controller.EventHandler.
func (*eventController) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	panic("unimplemented")
}

// HandleOnGoingLive implements controller.EventHandler.
func (*eventController) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	panic("unimplemented")
}

// HandleOnGrenadeThrown implements controller.EventHandler.
func (*eventController) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	panic("unimplemented")
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler.
func (*eventController) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnKnifeRoundStarted implements controller.EventHandler.
func (*eventController) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	panic("unimplemented")
}

// HandleOnKnifeRoundWon implements controller.EventHandler.
func (*eventController) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	panic("unimplemented")
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler.
func (*eventController) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	panic("unimplemented")
}

// HandleOnMapPicked implements controller.EventHandler.
func (*eventController) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	panic("unimplemented")
}

// HandleOnMapResult implements controller.EventHandler.
func (*eventController) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	panic("unimplemented")
}

// HandleOnMapVetoed implements controller.EventHandler.
func (*eventController) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	panic("unimplemented")
}

// HandleOnMatchPaused implements controller.EventHandler.
func (*eventController) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	panic("unimplemented")
}

// HandleOnMatchUnpaused implements controller.EventHandler.
func (*eventController) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	panic("unimplemented")
}

// HandleOnMolotovDetonated implements controller.EventHandler.
func (*eventController) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnPauseBegan implements controller.EventHandler.
func (*eventController) HandleOnPauseBegan(ctx context.Context, p models.OnPauseBeganPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerBecameMVP implements controller.EventHandler.
func (*eventController) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerConnected implements controller.EventHandler.
func (*eventController) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerDeath implements controller.EventHandler.
func (*eventController) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerDisconnected implements controller.EventHandler.
func (*eventController) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	panic("unimplemented")
}

// HandleOnPlayerSay implements controller.EventHandler.
func (*eventController) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	panic("unimplemented")
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler.
func (*eventController) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	panic("unimplemented")
}

// HandleOnRoundEnd implements controller.EventHandler.
func (*eventController) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	panic("unimplemented")
}

// HandleOnRoundStart implements controller.EventHandler.
func (*eventController) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	panic("unimplemented")
}

// HandleOnRoundStatsUpdated implements controller.EventHandler.
func (*eventController) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	panic("unimplemented")
}

// HandleOnSeriesInit implements controller.EventHandler.
func (*eventController) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	panic("unimplemented")
}

// HandleOnSeriesResult implements controller.EventHandler.
func (*eventController) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	panic("unimplemented")
}

// HandleOnSidePicked implements controller.EventHandler.
func (*eventController) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	panic("unimplemented")
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler.
func (*eventController) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	panic("unimplemented")
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler.
func (*eventController) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	panic("unimplemented")
}

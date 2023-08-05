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
func (ec *eventController) Close() error {
	return nil
}

// HandleOnBackupRestore implements controller.EventHandler.
func (ec *eventController) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	return nil
}

// HandleOnBombDefused implements controller.EventHandler.
func (ec *eventController) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	return nil
}

// HandleOnBombExploded implements controller.EventHandler.
func (ec *eventController) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	return nil
}

// HandleOnBombPlanted implements controller.EventHandler.
func (ec *eventController) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	return nil
}

// HandleOnDecoyStarted implements controller.EventHandler.
func (ec *eventController) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	return nil
}

// HandleOnDemoFinished implements controller.EventHandler.
func (ec *eventController) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	return nil
}

// HandleOnDemoUploadEnded implements controller.EventHandler.
func (ec *eventController) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	return nil
}

// HandleOnFlashbangDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	return nil
}

// HandleOnGameStateChanged implements controller.EventHandler.
func (ec *eventController) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	return nil
}

// HandleOnGoingLive implements controller.EventHandler.
func (ec *eventController) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	return nil
}

// HandleOnGrenadeThrown implements controller.EventHandler.
func (ec *eventController) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnKnifeRoundStarted implements controller.EventHandler.
func (ec *eventController) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	return nil
}

// HandleOnKnifeRoundWon implements controller.EventHandler.
func (ec *eventController) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler.
func (ec *eventController) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	return nil
}

// HandleOnMapPicked implements controller.EventHandler.
func (ec *eventController) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	return nil
}

// HandleOnMapResult implements controller.EventHandler.
func (ec *eventController) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	return nil
}

// HandleOnMapVetoed implements controller.EventHandler.
func (ec *eventController) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	return nil
}

// HandleOnMatchPaused implements controller.EventHandler.
func (ec *eventController) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	return nil
}

// HandleOnMatchUnpaused implements controller.EventHandler.
func (ec *eventController) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	return nil
}

// HandleOnMolotovDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	return nil
}

// HandleOnPauseBegan implements controller.EventHandler.
func (ec *eventController) HandleOnPauseBegan(ctx context.Context, p models.OnPauseBeganPayload) error {
	return nil
}

// HandleOnPlayerBecameMVP implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	return nil
}

// HandleOnPlayerConnected implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	return nil
}

// HandleOnPlayerDeath implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	return nil
}

// HandleOnPlayerDisconnected implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	return nil
}

// HandleOnPlayerSay implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler.
func (ec *eventController) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	return nil
}

// HandleOnRoundEnd implements controller.EventHandler.
func (ec *eventController) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	return nil
}

// HandleOnRoundStart implements controller.EventHandler.
func (ec *eventController) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	return nil
}

// HandleOnRoundStatsUpdated implements controller.EventHandler.
func (ec *eventController) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	return nil
}

// HandleOnSeriesInit implements controller.EventHandler.
func (ec *eventController) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	return nil
}

// HandleOnSeriesResult implements controller.EventHandler.
func (ec *eventController) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	return nil
}

// HandleOnSidePicked implements controller.EventHandler.
func (ec *eventController) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler.
func (ec *eventController) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	return nil
}

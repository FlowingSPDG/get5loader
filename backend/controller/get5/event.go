package get5

import (
	"context"

	got5 "github.com/FlowingSPDG/Got5"
)

type eventController struct {
}

func NewGot5EventController() got5.EventHandler {
	return &eventController{}
}

// Close implements controller.EventHandler.
func (ec *eventController) Close() error {
	return nil
}

// HandleOnBackupRestore implements controller.EventHandler.
func (ec *eventController) HandleOnBackupRestore(ctx context.Context, p got5.OnBackupRestorePayload) error {
	return nil
}

// HandleOnBombDefused implements controller.EventHandler.
func (ec *eventController) HandleOnBombDefused(ctx context.Context, p got5.OnBombDefusedPayload) error {
	return nil
}

// HandleOnBombExploded implements controller.EventHandler.
func (ec *eventController) HandleOnBombExploded(ctx context.Context, p got5.OnBombExplodedPayload) error {
	return nil
}

// HandleOnBombPlanted implements controller.EventHandler.
func (ec *eventController) HandleOnBombPlanted(ctx context.Context, p got5.OnBombPlantedPayload) error {
	return nil
}

// HandleOnDecoyStarted implements controller.EventHandler.
func (ec *eventController) HandleOnDecoyStarted(ctx context.Context, p got5.OnDecoyStartedPayload) error {
	return nil
}

// HandleOnDemoFinished implements controller.EventHandler.
func (ec *eventController) HandleOnDemoFinished(ctx context.Context, p got5.OnDemoFinishedPayload) error {
	return nil
}

// HandleOnDemoUploadEnded implements controller.EventHandler.
func (ec *eventController) HandleOnDemoUploadEnded(ctx context.Context, p got5.OnDemoUploadEndedPayload) error {
	return nil
}

// HandleOnFlashbangDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnFlashbangDetonated(ctx context.Context, p got5.OnFlashbangDetonatedPayload) error {
	return nil
}

// HandleOnGameStateChanged implements controller.EventHandler.
func (ec *eventController) HandleOnGameStateChanged(ctx context.Context, p got5.OnGameStateChangedPayload) error {
	return nil
}

// HandleOnGoingLive implements controller.EventHandler.
func (ec *eventController) HandleOnGoingLive(ctx context.Context, p got5.OnGoingLivePayload) error {
	return nil
}

// HandleOnGrenadeThrown implements controller.EventHandler.
func (ec *eventController) HandleOnGrenadeThrown(ctx context.Context, p got5.OnGrenadeThrownPayload) error {
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnHEGrenadeDetonated(ctx context.Context, p got5.OnHEGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnKnifeRoundStarted implements controller.EventHandler.
func (ec *eventController) HandleOnKnifeRoundStarted(ctx context.Context, p got5.OnKnifeRoundStartedPayload) error {
	return nil
}

// HandleOnKnifeRoundWon implements controller.EventHandler.
func (ec *eventController) HandleOnKnifeRoundWon(ctx context.Context, p got5.OnKnifeRoundWonPayload) error {
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler.
func (ec *eventController) HandleOnLoadMatchConfigFailed(ctx context.Context, p got5.OnLoadMatchConfigFailedPayload) error {
	return nil
}

// HandleOnMapPicked implements controller.EventHandler.
func (ec *eventController) HandleOnMapPicked(ctx context.Context, p got5.OnMapPickedPayload) error {
	return nil
}

// HandleOnMapResult implements controller.EventHandler.
func (ec *eventController) HandleOnMapResult(ctx context.Context, p got5.OnMapResultPayload) error {
	return nil
}

// HandleOnMapVetoed implements controller.EventHandler.
func (ec *eventController) HandleOnMapVetoed(ctx context.Context, p got5.OnMapVetoedPayload) error {
	return nil
}

// HandleOnMatchPaused implements controller.EventHandler.
func (ec *eventController) HandleOnMatchPaused(ctx context.Context, p got5.OnMatchPausedPayload) error {
	return nil
}

// HandleOnMatchUnpaused implements controller.EventHandler.
func (ec *eventController) HandleOnMatchUnpaused(ctx context.Context, p got5.OnMatchUnpausedPayload) error {
	return nil
}

// HandleOnMolotovDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnMolotovDetonated(ctx context.Context, p got5.OnMolotovDetonatedPayload) error {
	return nil
}

// HandleOnPauseBegan implements controller.EventHandler.
func (ec *eventController) HandleOnPauseBegan(ctx context.Context, p got5.OnPauseBeganPayload) error {
	return nil
}

// HandleOnPlayerBecameMVP implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerBecameMVP(ctx context.Context, p got5.OnPlayerBecameMVPPayload) error {
	return nil
}

// HandleOnPlayerConnected implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerConnected(ctx context.Context, p got5.OnPlayerConnectedPayload) error {
	return nil
}

// HandleOnPlayerDeath implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerDeath(ctx context.Context, p got5.OnPlayerDeathPayload) error {
	return nil
}

// HandleOnPlayerDisconnected implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerDisconnected(ctx context.Context, p got5.OnPlayerDisconnectedPayload) error {
	return nil
}

// HandleOnPlayerSay implements controller.EventHandler.
func (ec *eventController) HandleOnPlayerSay(ctx context.Context, p got5.OnPlayerSayPayload) error {
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler.
func (ec *eventController) HandleOnPreLoadMatchConfig(ctx context.Context, p got5.OnPreLoadMatchConfigPayload) error {
	return nil
}

// HandleOnRoundEnd implements controller.EventHandler.
func (ec *eventController) HandleOnRoundEnd(ctx context.Context, p got5.OnRoundEndPayload) error {
	return nil
}

// HandleOnRoundStart implements controller.EventHandler.
func (ec *eventController) HandleOnRoundStart(ctx context.Context, p got5.OnRoundStartPayload) error {
	return nil
}

// HandleOnRoundStatsUpdated implements controller.EventHandler.
func (ec *eventController) HandleOnRoundStatsUpdated(ctx context.Context, p got5.OnRoundStatsUpdatedPayload) error {
	return nil
}

// HandleOnSeriesInit implements controller.EventHandler.
func (ec *eventController) HandleOnSeriesInit(ctx context.Context, p got5.OnSeriesInitPayload) error {
	return nil
}

// HandleOnSeriesResult implements controller.EventHandler.
func (ec *eventController) HandleOnSeriesResult(ctx context.Context, p got5.OnSeriesResultPayload) error {
	return nil
}

// HandleOnSidePicked implements controller.EventHandler.
func (ec *eventController) HandleOnSidePicked(ctx context.Context, p got5.OnSidePickedPayload) error {
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler.
func (ec *eventController) HandleOnSmokeGrenadeDetonated(ctx context.Context, p got5.OnSmokeGrenadeDetonatedPayload) error {
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler.
func (ec *eventController) HandleOnTeamReadyStatusChanged(ctx context.Context, p got5.OnTeamReadyStatusChangedPayload) error {
	return nil
}

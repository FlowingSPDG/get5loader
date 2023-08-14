package dataloaders

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader/v7"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

type teamPlayersLoader struct {
	player usecase.Player
}

func (t *teamPlayersLoader) BatchGetPlayers(ctx context.Context, IDs []entity.TeamID) []*dataloader.Result[[]*entity.Player] {
	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[[]*entity.Player], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[[]*entity.Player]{
			Error: errors.New("not found"),
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[entity.TeamID]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	teamPlayers, err := t.player.BatchGetPlayersByTeam(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for teamID, players := range teamPlayers {
		var result *dataloader.Result[[]*entity.Player]
		if err != nil {
			result = &dataloader.Result[[]*entity.Player]{
				Error: err,
			}
			continue
		}
		result = &dataloader.Result[[]*entity.Player]{
			Error: nil,
			Data:  players,
		}
		results[indexs[teamID]] = result
	}
	return results
}

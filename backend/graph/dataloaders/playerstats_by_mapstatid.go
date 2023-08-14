package dataloaders

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

type mapstatPlayerStats struct {
	playerstat usecase.PlayerStat
}

func (m *mapstatPlayerStats) BatchGetMapstats(ctx context.Context, IDs []entity.MapStatsID) []*dataloader.Result[[]*entity.PlayerStat] {
	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[[]*entity.PlayerStat], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[[]*entity.PlayerStat]{
			Data: []*entity.PlayerStat{},
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[entity.MapStatsID]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	matchMapstats, err := m.playerstat.BatchGetPlayerStatsByMapstat(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for matchID, mapstats := range matchMapstats {
		var result *dataloader.Result[[]*entity.PlayerStat]
		if err != nil {
			result = &dataloader.Result[[]*entity.PlayerStat]{
				Error: err,
			}
			continue
		}
		result = &dataloader.Result[[]*entity.PlayerStat]{
			Error: nil,
			Data:  mapstats,
		}
		results[indexs[matchID]] = result
	}
	return results
}

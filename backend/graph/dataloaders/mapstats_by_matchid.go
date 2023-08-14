package dataloaders

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader/v7"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
)

type matchMapstat struct {
	mapstat usecase.Mapstat
}

func (m *matchMapstat) BatchGetMapStats(ctx context.Context, IDs []entity.MatchID) []*dataloader.Result[[]*entity.MapStat] {
	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[[]*entity.MapStat], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[[]*entity.MapStat]{
			Error: errors.New("not found"),
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[entity.MatchID]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	matchMapstats, err := m.mapstat.BatchGetMapstatsByMatch(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for matchID, mapstats := range matchMapstats {
		var result *dataloader.Result[[]*entity.MapStat]
		if err != nil {
			result = &dataloader.Result[[]*entity.MapStat]{
				Error: err,
			}
			continue
		}
		result = &dataloader.Result[[]*entity.MapStat]{
			Error: nil,
			Data:  mapstats,
		}
		results[indexs[matchID]] = result
	}
	return results
}

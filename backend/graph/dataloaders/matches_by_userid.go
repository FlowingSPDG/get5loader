package dataloaders

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
	"github.com/graph-gophers/dataloader/v7"
)

type userMatchLoader struct {
	match usecase.Match
}

func (u *userMatchLoader) BatchGetMatches(ctx context.Context, IDs []entity.UserID) []*dataloader.Result[[]*entity.Match] {
	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[[]*entity.Match], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[[]*entity.Match]{
			Data: []*entity.Match{},
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[entity.UserID]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	userMatches, err := u.match.BatchGetMatchesByUser(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for userID, matches := range userMatches {
		var result *dataloader.Result[[]*entity.Match]
		if err != nil {
			result = &dataloader.Result[[]*entity.Match]{
				Error: err,
			}
			continue
		}
		result = &dataloader.Result[[]*entity.Match]{
			Error: nil,
			Data:  matches,
		}
		results[indexs[userID]] = result
	}
	return results
}

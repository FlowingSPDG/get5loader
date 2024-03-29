package dataloaders

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
	"github.com/graph-gophers/dataloader/v7"
)

type userTeamLoader struct {
	team usecase.Team
}

func (u *userTeamLoader) BatchGetTeams(ctx context.Context, IDs []entity.UserID) []*dataloader.Result[[]*entity.Team] {
	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[[]*entity.Team], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[[]*entity.Team]{
			Data: []*entity.Team{},
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[entity.UserID]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	userTeams, err := u.team.BatchGetTeamsByUsers(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for userID, teams := range userTeams {
		var result *dataloader.Result[[]*entity.Team]
		if err != nil {
			result = &dataloader.Result[[]*entity.Team]{
				Error: err,
			}
			continue
		}
		result = &dataloader.Result[[]*entity.Team]{
			Error: nil,
			Data:  teams,
		}
		results[indexs[userID]] = result
	}
	return results
}

package dataloaders

import (
	"context"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/FlowingSPDG/get5loader/backend/usecase"
	"github.com/graph-gophers/dataloader/v7"
)

type userServersLoader struct {
	server usecase.GameServer
}

func (u *userServersLoader) BatchGetServers(ctx context.Context, IDs []entity.UserID) []*dataloader.Result[[]*entity.GameServer] {
	// 引数と戻り値のスライスlenは等しくする
	results := make([]*dataloader.Result[[]*entity.GameServer], len(IDs))
	for i := range results {
		results[i] = &dataloader.Result[[]*entity.GameServer]{
			Data: []*entity.GameServer{},
		}
	}

	// 検索条件であるIDが、引数でもらったIDsスライスの何番目のインデックスに格納されていたのか検索できるようにmap化する
	indexs := make(map[entity.UserID]int, len(IDs))
	for i, ID := range IDs {
		indexs[ID] = i
	}

	userServers, err := u.server.BatchGetGameServersByUser(ctx, IDs)

	// 取得結果を、戻り値resultの中の適切な場所に格納する
	for userID, servers := range userServers {
		var result *dataloader.Result[[]*entity.GameServer]
		if err != nil {
			result = &dataloader.Result[[]*entity.GameServer]{
				Error: err,
			}
			continue
		}
		result = &dataloader.Result[[]*entity.GameServer]{
			Error: nil,
			Data:  servers,
		}
		results[indexs[userID]] = result
	}
	return results
}

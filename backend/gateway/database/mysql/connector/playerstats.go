package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/playerstats"
)

type mysqlPlayerstatsRepositoryConnector struct {
	connector database.DBConnector
}

func NewMySQLPlayerstatsRepositoryConnector(connector database.DBConnector) database.RepositoryConnector[database.PlayerStatsRepository] {
	return &mysqlPlayerstatsRepositoryConnector{connector: connector}
}

// TODO: トランザクション処理を含んだバージョンを作成する

// Open implements database.RepositoryConnector.
func (mmrc *mysqlPlayerstatsRepositoryConnector) Open() (database.PlayerStatsRepository, error) {
	if err := mmrc.connector.Open(); err != nil {
		return nil, err
	}

	db := mmrc.connector.GetConnection()

	repository := playerstats.NewPlayerStatsRepository(db)

	return repository, nil
}

func (mmrc *mysqlPlayerstatsRepositoryConnector) Close() error {
	return mmrc.connector.Close()
}

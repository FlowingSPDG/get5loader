package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/mapstats"
)

type mysqlMapstatsRepositoryConnector struct {
	connector database.DBConnector
}

func NewMySQLMapstatsRepositoryConnector(connector database.DBConnector) database.RepositoryConnector[database.MapStatsRepository] {
	return &mysqlMapstatsRepositoryConnector{connector: connector}
}

// TODO: トランザクション処理を含んだバージョンを作成する

// Open implements database.RepositoryConnector.
func (mmrc *mysqlMapstatsRepositoryConnector) Open() (database.MapStatsRepository, error) {
	if err := mmrc.connector.Open(); err != nil {
		return nil, err
	}

	db := mmrc.connector.GetConnection()

	repository := mapstats.NewMapStatsRepository(db)

	return repository, nil
}

func (mmrc *mysqlMapstatsRepositoryConnector) Close() error {
	return mmrc.connector.Close()
}

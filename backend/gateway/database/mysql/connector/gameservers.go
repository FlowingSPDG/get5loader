package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/gameservers"
)

type mysqlGameserversRepositoryConnector struct {
	connector database.DBConnector
}

func NewMySQLGameservesRepositoryConnector(connector database.DBConnector) database.RepositoryConnector[database.GameServersRepository] {
	return &mysqlGameserversRepositoryConnector{connector: connector}
}

// TODO: トランザクション処理を含んだバージョンを作成する

// Open implements database.RepositoryConnector.
func (mgrc *mysqlGameserversRepositoryConnector) Open() (database.GameServersRepository, error) {
	if err := mgrc.connector.Open(); err != nil {
		return nil, err
	}

	db := mgrc.connector.GetConnection()

	repository := gameservers.NewGameServerRepository(db)

	return repository, nil
}

func (mgrc *mysqlGameserversRepositoryConnector) Close() error {
	return mgrc.connector.Close()
}

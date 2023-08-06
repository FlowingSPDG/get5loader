package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/matches"
)

type mysqlMatchesRepositoryConnector struct {
	connector database.DBConnector
}

func NewMySQLMatchesRepositoryConnector(connector database.DBConnector) database.RepositoryConnector[database.MatchesRepository] {
	return &mysqlMatchesRepositoryConnector{connector: connector}
}

// TODO: トランザクション処理を含んだバージョンを作成する

// Open implements database.RepositoryConnector.
func (mmrc *mysqlMatchesRepositoryConnector) Open() (database.MatchesRepository, error) {
	if err := mmrc.connector.Open(); err != nil {
		return nil, err
	}

	db := mmrc.connector.GetConnection()

	repository := matches.NewMatchRepository(db)

	return repository, nil
}

func (mmrc *mysqlMatchesRepositoryConnector) Close() error {
	return mmrc.connector.Close()
}

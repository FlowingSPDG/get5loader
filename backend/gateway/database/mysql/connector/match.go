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

// Open implements database.UserRepositoryConnector.
func (mrc *mysqlMatchesRepositoryConnector) Open() (database.MatchesRepository, error) {
	if err := mrc.connector.Open(); err != nil {
		return nil, err
	}

	db := mrc.connector.GetConnection()

	repository := matches.NewMatchRepository(db)

	return repository, nil
}

func (mrc *mysqlMatchesRepositoryConnector) Close() error {
	return mrc.connector.Close()
}

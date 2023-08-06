package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/teams"
)

type mysqlTeamsRepositoryConnector struct {
	connector database.DBConnector
}

func NewMySQLTeamsRepositoryConnector(connector database.DBConnector) database.RepositoryConnector[database.TeamsRepository] {
	return &mysqlTeamsRepositoryConnector{connector: connector}
}

// TODO: トランザクション処理を含んだバージョンを作成する

// Open implements database.RepositoryConnector.
func (mtrc *mysqlTeamsRepositoryConnector) Open() (database.TeamsRepository, error) {
	if err := mtrc.connector.Open(); err != nil {
		return nil, err
	}

	db := mtrc.connector.GetConnection()

	repository := teams.NewPlayersRepository(db)

	return repository, nil
}

func (mtrc *mysqlTeamsRepositoryConnector) Close() error {
	return mtrc.connector.Close()
}

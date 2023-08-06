package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/players"
)

type mysqlPlayersRepositoryConnector struct {
	connector database.DBConnector
}

func NewMySQLPlayersRepositoryConnector(connector database.DBConnector) database.RepositoryConnector[database.PlayersRepository] {
	return &mysqlPlayersRepositoryConnector{connector: connector}
}

// TODO: トランザクション処理を含んだバージョンを作成する

// Open implements database.RepositoryConnector.
func (mmrc *mysqlPlayersRepositoryConnector) Open() (database.PlayersRepository, error) {
	if err := mmrc.connector.Open(); err != nil {
		return nil, err
	}

	db := mmrc.connector.GetConnection()

	repository := players.NewPlayersRepository(db)

	return repository, nil
}

func (mmrc *mysqlPlayersRepositoryConnector) Close() error {
	return mmrc.connector.Close()
}

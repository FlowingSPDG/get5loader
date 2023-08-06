package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/matches"
)

type mysqlMatchesRepositoryConnector struct {
	dsn       string
	connector database.DBConnector
}

func NewMySQLMatchesRepositoryConnector(connector database.DBConnector) database.MatchesRepositoryConnector {
	return &mysqlMatchesRepositoryConnector{connector: connector}
}

// Open implements database.UserRepositoryConnector.
func (murc *mysqlMatchesRepositoryConnector) Open() (database.MatchesRepository, error) {
	if err := murc.connector.Open(); err != nil {
		return nil, err
	}

	db := murc.connector.GetConnection()

	repository := matches.NewMatchRepository(db)

	return repository, nil
}

func (murc *mysqlMatchesRepositoryConnector) Close() error {
	return murc.connector.Close()
}

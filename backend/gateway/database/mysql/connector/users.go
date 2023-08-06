package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users"
)

type mysqlUsersRepositoryConnector struct {
	dsn       string
	connector database.DBConnector
}

func NewMySQLUsersRepositoryConnector(connector database.DBConnector) database.UserRepositoryConnector {
	return &mysqlUsersRepositoryConnector{connector: connector}
}

// Open implements database.UserRepositoryConnector.
func (murc *mysqlUsersRepositoryConnector) Open() (database.UsersRepositry, error) {
	if err := murc.connector.Open(); err != nil {
		return nil, err
	}

	db := murc.connector.GetConnection()

	repository := users.NewUsersRepositry(db)

	return repository, nil
}

func (murc *mysqlUsersRepositoryConnector) Close() error {
	return murc.connector.Close()
}

package mysqlconnector

import (
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database"
	"github.com/FlowingSPDG/get5-web-go/backend/gateway/database/mysql/users"
)

type mysqlUsersRepositoryConnector struct {
	connector database.DBConnector
}

func NewMySQLUsersRepositoryConnector(connector database.DBConnector) database.RepositoryConnector[database.UsersRepositry] {
	return &mysqlUsersRepositoryConnector{connector: connector}
}

// TODO: トランザクション処理を含んだバージョンを作成する

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

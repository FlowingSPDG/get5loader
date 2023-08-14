package gin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/get5loader/backend/gateway/database"
	"github.com/FlowingSPDG/get5loader/backend/service/uuid"
)

type DatabaseConnectionMiddleware interface {
	Handle(c *gin.Context)
}

type databaseConnectionMiddleware struct {
	uuidGenerator uuid.UUIDGenerator
	connector     database.RepositoryConnector
}

func NewDatabaseConnectionMiddleware(
	uuidGenerator uuid.UUIDGenerator,
	connector database.RepositoryConnector,
) DatabaseConnectionMiddleware {
	return &databaseConnectionMiddleware{
		uuidGenerator: uuidGenerator,
		connector:     connector,
	}
}

func (d *databaseConnectionMiddleware) Handle(c *gin.Context) {
	if err := d.connector.Open(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	defer d.connector.Close()

	database.SetConnectionForGinContext(c, d.connector)

	c.Next()
}

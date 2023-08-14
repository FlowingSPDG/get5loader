package database

import (
	"context"

	"github.com/gin-gonic/gin"
)

var (
	// RepositoryConnectorKey is a key to store repositoryConnector in gin.Context.
	RepositoryConnectorKey = "repositoryConnector"
)

// GetConnectionFromGinContext gin.ContextからRepositoryConnectorを取得する
func GetConnectionFromGinContext(c *gin.Context) RepositoryConnector {
	return c.MustGet(RepositoryConnectorKey).(RepositoryConnector)
}

// SetConnectionForGinContext gin.ContextにRepositoryConnectorを設定する
func SetConnectionForGinContext(c *gin.Context, repositoryConnector RepositoryConnector) {
	c.Set(RepositoryConnectorKey, repositoryConnector)
}

func SetConnection(ctx context.Context, repositoryConnector RepositoryConnector) context.Context {
	return context.WithValue(ctx, RepositoryConnectorKey, repositoryConnector)
}

func GetConnection(ctx context.Context) RepositoryConnector {
	return ctx.Value(RepositoryConnectorKey).(RepositoryConnector)
}

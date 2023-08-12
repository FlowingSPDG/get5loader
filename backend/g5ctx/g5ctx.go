package g5ctx

import (
	"context"
	"errors"

	"github.com/FlowingSPDG/get5loader/backend/entity"
	"github.com/gin-gonic/gin"
)

var (
	ErrContextValueNotFound = errors.New("context value not found")
)

type ctxKey string

var (
	operatorKey = "operation"
	userKey     = "user"
)

type OperationType int

const (
	OperationTypeUnknown OperationType = iota
	OperationTypeSystem
	OperationTypeUser
)

func SetOperationGinContext(c *gin.Context, op OperationType) {
	c.Set(operatorKey, op)
}

// SetOperation sets operator to context.
func SetOperation(ctx context.Context, op OperationType) context.Context {
	return context.WithValue(ctx, operatorKey, op)
}

// GetOperation gets operator from context.
func GetOperation(ctx context.Context) (OperationType, error) {
	op, ok := ctx.Value(operatorKey).(OperationType)
	if !ok {
		return OperationTypeUnknown, ErrContextValueNotFound
	}
	return op, nil
}

func SetUserTokenGinContext(c *gin.Context, user *entity.TokenUser) {
	c.Set(userKey, user)
}

func SetUserToken(ctx context.Context, user *entity.TokenUser) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// GetUser gets user from context.
func GetUserToken(ctx context.Context) (*entity.TokenUser, error) {
	user, ok := ctx.Value(userKey).(*entity.TokenUser)
	if !ok {
		return nil, ErrContextValueNotFound
	}
	return user, nil
}

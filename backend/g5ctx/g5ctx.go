package g5ctx

import "context"

type ctxKey struct{}

var (
	// AdminKey is a key for admin.
	AdminKey = ctxKey{}
	// UserIDKey is a key for userID.
	UserIDKey = ctxKey{}
)

// SetAdmin sets admin to context.
func SetAdmin(ctx context.Context, admin bool) context.Context {
	return context.WithValue(ctx, AdminKey, admin)
}

// SetUserID sets userID to context.
func SetUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetAdmin gets admin from context.
func GetAdmin(ctx context.Context) bool {
	admin, ok := ctx.Value(AdminKey).(bool)
	if !ok {
		return false
	}
	return admin
}

// GetUserID gets userID from context.
func GetUserID(ctx context.Context) int {
	userID, ok := ctx.Value(UserIDKey).(int)
	if !ok {
		return 0
	}
	return userID
}

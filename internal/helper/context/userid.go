package context

import (
	"context"
)

type userIDContextKey struct{}

var userIDKey userIDContextKey

func GetUser(ctx context.Context) string {
	val := ctx.Value(userIDKey)
	if val == nil {
		return ""
	}
	return val.(string)
}

func SetUser(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

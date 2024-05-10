package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserIDContext(t *testing.T) {
	ctx := context.Background()
	userID := "test-id"

	t.Run("GetUser return empty string on userIDKey not set", func(t *testing.T) {
		shouldEmpty := GetUser(ctx)
		assert.Equal(t, "", shouldEmpty)
	})

	t.Run("GetUser return userID after SetUser", func(t *testing.T) {
		newCtx := SetUser(ctx, userID)
		assert.Equal(t, userID, GetUser(newCtx))
	})
}

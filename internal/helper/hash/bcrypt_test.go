package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBCrypt(t *testing.T) {
	password := "123456"

	t.Run("BCrypt success and match", func(t *testing.T) {
		passwordHash, err := BCrypt(password)
		match := CheckBCrypt(password, passwordHash)

		assert.NoError(t, err)
		assert.True(t, match)
	})

	t.Run("BCrypt success and not match", func(t *testing.T) {
		passwordHash, err := BCrypt(password)
		match := CheckBCrypt("invalid-pass", passwordHash)

		assert.NoError(t, err)
		assert.False(t, match)
	})

	t.Run("BCrypt failed", func(t *testing.T) {
		_, err := BCrypt("1234567890123456789012345678901234567890123456789012345678901234567890123")

		assert.Error(t, err)
	})
}

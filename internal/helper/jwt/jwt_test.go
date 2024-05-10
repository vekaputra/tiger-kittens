package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
)

func TestJwt(t *testing.T) {
	expiredBefore := time.Now().Add(-time.Hour).Unix()
	expiredAfter := time.Now().Add(time.Hour).Unix()
	user := entity.User{
		ID:       "test-id",
		Email:    "test-email@mail.com",
		Password: "test-password",
		Username: "test-username",
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	assert.NoError(t, err)

	t.Run("GenerateAccessToken success and valid", func(t *testing.T) {
		token, errGenerate := GenerateAccessToken(privateKey, expiredAfter, user)
		claims, errDecode := DecodeAccessToken(privateKey, token)

		assert.NoError(t, errGenerate)
		assert.NoError(t, errDecode)
		assert.Equal(t, user.ID, claims["sub"].(string))
		assert.Equal(t, user.Username, claims["username"].(string))
		assert.Equal(t, user.Email, claims["email"].(string))
		assert.Equal(t, float64(expiredAfter), claims["exp"].(float64))
	})

	t.Run("GenerateAccessToken success and expired", func(t *testing.T) {
		token, errGenerate := GenerateAccessToken(privateKey, expiredBefore, user)
		claims, errDecode := DecodeAccessToken(privateKey, token)

		assert.NoError(t, errGenerate)
		assert.Error(t, errDecode)
		assert.Nil(t, claims)
	})
}

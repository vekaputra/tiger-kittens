package jwt

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

func GenerateAccessToken(key *rsa.PrivateKey, user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
	result, err := token.SignedString(key)
	if err != nil {
		return "", pkgerr.ErrWithStackTrace(err)
	}

	return result, nil
}

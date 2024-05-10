package jwt

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/repository/entity"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

func GenerateAccessToken(key *rsa.PrivateKey, expiredAfter int64, user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      expiredAfter,
	})

	result, err := token.SignedString(key)
	if err != nil {
		return "", pkgerr.ErrWithStackTrace(err)
	}

	return result, nil
}

func DecodeAccessToken(jwtPrivateKey *rsa.PrivateKey, accessToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPrivateKey.Public(), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to decode token")
		return nil, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidAccessToken)
	}
	return claims, nil
}

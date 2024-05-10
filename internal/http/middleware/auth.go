package middleware

import (
	"crypto/rsa"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/helper/context"
	"github.com/vekaputra/tiger-kittens/internal/helper/customerror"
	"github.com/vekaputra/tiger-kittens/internal/helper/response"
	pkgerr "github.com/vekaputra/tiger-kittens/pkg/error"
)

func Auth(jwtPrivateKey *rsa.PrivateKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerSchema := "Bearer "
			accessToken := strings.TrimPrefix(c.Request().Header.Get(echo.HeaderAuthorization), bearerSchema)
			claims := jwt.MapClaims{}

			_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtPrivateKey.Public(), nil
			})
			if err != nil {
				log.Error().Err(err).Msg("failed to decode token")
				return response.SendResponseWithNativeError(c, pkgerr.ErrWithStackTrace(customerror.ErrorInvalidAccessToken))
			}

			ctx := c.Request().Context()
			ctx = context.SetUser(ctx, claims["sub"].(string))

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

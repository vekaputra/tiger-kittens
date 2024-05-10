package middleware

import (
	"crypto/rsa"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/vekaputra/tiger-kittens/internal/helper/context"
	"github.com/vekaputra/tiger-kittens/internal/helper/jwt"
	"github.com/vekaputra/tiger-kittens/internal/helper/response"
)

func Auth(jwtPrivateKey *rsa.PrivateKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerSchema := "Bearer "
			accessToken := strings.TrimPrefix(c.Request().Header.Get(echo.HeaderAuthorization), bearerSchema)

			claims, err := jwt.DecodeAccessToken(jwtPrivateKey, accessToken)
			if err != nil {
				return response.SendResponseWithNativeError(c, err)
			}

			ctx := c.Request().Context()
			ctx = context.SetUser(ctx, claims["sub"].(string))

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

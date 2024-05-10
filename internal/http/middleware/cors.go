package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func HandleOptionsRequest(methods ...string) func(echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
		c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-Forwarded-For, X-Real-IP, X-Forwarded-Proto")
		return c.NoContent(http.StatusOK)
	}
}

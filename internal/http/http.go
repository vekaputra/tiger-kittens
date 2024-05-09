package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vekaputra/tiger-kittens/internal/model"
	"github.com/vekaputra/tiger-kittens/internal/service"
)

type AppServer struct {
	UserService service.UserServiceProvider
}

func (s *AppServer) ReadinessCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, model.HealthCheckResponse{
		Status: "OK",
	})
}

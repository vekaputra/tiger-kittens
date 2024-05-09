package app

import (
	nhttp "net/http"

	"github.com/labstack/echo/v4"
	"github.com/vekaputra/tiger-kittens/internal/config"
	"github.com/vekaputra/tiger-kittens/internal/http"
	"github.com/vekaputra/tiger-kittens/internal/http/middleware"
)

type Server struct {
	Connection *Connection
	Server     *EchoServer
}

func NewServer(appConfig *config.Config) *Server {
	conn := NewConnection(appConfig)
	repo := NewRepo(conn)
	service := NewService(repo)

	srv := &http.AppServer{
		UserService: service.UserService,
	}

	e := NewEcho(appConfig)
	route(e.Echo, srv, appConfig)

	return &Server{
		Connection: conn,
		Server:     e,
	}
}

func route(e *echo.Echo, srv *http.AppServer, appConfig *config.Config) {
	e.GET("/healthcheck", srv.ReadinessCheck)

	v1Group := e.Group("/v1")

	userGroup := v1Group.Group("/user")
	userGroup.POST("/register", srv.RegisterUser)

	if appConfig.IsAllowCORS {
		e.OPTIONS("/healthcheck", middleware.HandleOptionsRequest(nhttp.MethodGet))

		userGroup.OPTIONS("/register", middleware.HandleOptionsRequest(nhttp.MethodPost))
	}
}

package app

import (
	nhttp "net/http"

	"github.com/vekaputra/tiger-kittens/internal/helper/mailqueue"

	"github.com/labstack/echo/v4"
	"github.com/vekaputra/tiger-kittens/internal/config"
	"github.com/vekaputra/tiger-kittens/internal/http"
	"github.com/vekaputra/tiger-kittens/internal/http/middleware"
)

type Server struct {
	Connection *Connection
	Server     *EchoServer
	MailQueue  *mailqueue.MailQueue
}

func NewServer(appConfig *config.Config) *Server {
	conn := NewConnection(appConfig)
	repo := NewRepo(conn)
	mailQueue := mailqueue.New(appConfig.EmailConfig, repo.TigerRepo, repo.UserRepo)
	service := NewService(repo, appConfig, mailQueue)

	srv := &http.AppServer{
		TigerService: service.TigerService,
		UserService:  service.UserService,
	}

	e := NewEcho(appConfig)
	route(e.Echo, srv, appConfig)

	return &Server{
		Connection: conn,
		Server:     e,
		MailQueue:  mailQueue,
	}
}

func route(e *echo.Echo, srv *http.AppServer, appConfig *config.Config) {
	e.GET("/healthcheck", srv.ReadinessCheck)

	v1Group := e.Group("/v1")

	userGroup := v1Group.Group("/user")
	userGroup.POST("/register", srv.RegisterUser)
	userGroup.POST("/login", srv.LoginUser)

	tigerGroup := v1Group.Group("/tiger")
	tigerGroup.GET("", srv.ListTiger)
	tigerGroup.GET("/sighting", srv.ListSighting)
	tigerGroup.POST("", srv.CreateTiger, middleware.Auth(appConfig.JWTConfig.PrivateKey))
	tigerGroup.POST("/sighting", srv.CreateSighting, middleware.Auth(appConfig.JWTConfig.PrivateKey))

	if appConfig.IsAllowCORS {
		e.OPTIONS("/healthcheck", middleware.HandleOptionsRequest(nhttp.MethodGet))

		userGroup.OPTIONS("/register", middleware.HandleOptionsRequest(nhttp.MethodPost))
		userGroup.OPTIONS("/login", middleware.HandleOptionsRequest(nhttp.MethodPost))

		tigerGroup.OPTIONS("", middleware.HandleOptionsRequest(nhttp.MethodPost, nhttp.MethodGet))
		tigerGroup.OPTIONS("/sighting", middleware.HandleOptionsRequest(nhttp.MethodPost, nhttp.MethodGet))
	}
}

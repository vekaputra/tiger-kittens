package app

import (
	nhttp "net/http"

	"github.com/go-playground/validator/v10"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/vekaputra/tiger-kittens/graph"
	"github.com/vekaputra/tiger-kittens/internal/config"
	"github.com/vekaputra/tiger-kittens/internal/helper/mailqueue"
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

	gqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					Validate:     validator.New(validator.WithRequiredStructEnabled()),
					TigerService: service.TigerService,
					UserService:  service.UserService,
				},
			},
		),
	)
	playGroundHandler := playground.Handler("Tiger Kittens", "/gql/query")

	e := NewEcho(appConfig)
	route(e.Echo, srv, gqlHandler, playGroundHandler, appConfig)

	return &Server{
		Connection: conn,
		Server:     e,
		MailQueue:  mailQueue,
	}
}

func route(e *echo.Echo, srv *http.AppServer, gql *handler.Server, playground nhttp.HandlerFunc, appConfig *config.Config) {
	e.GET("/healthcheck", srv.ReadinessCheck)

	gqlGroup := e.Group("/gql")
	gqlGroup.POST("/query", func(c echo.Context) error {
		gql.ServeHTTP(c.Response(), c.Request())
		return nil
	}, middleware.Auth(appConfig.JWTConfig.PrivateKey))
	gqlGroup.GET("/playground", func(c echo.Context) error {
		playground.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	v1Group := e.Group("/v1")

	userGroup := v1Group.Group("/user")
	userGroup.POST("/register", srv.RegisterUser)
	userGroup.POST("/login", srv.LoginUser)

	tigerGroup := v1Group.Group("/tiger")
	tigerGroup.GET("", srv.ListTiger)
	tigerGroup.GET("/:tigerID/sighting", srv.ListSighting)
	tigerGroup.POST("", srv.CreateTiger, middleware.Auth(appConfig.JWTConfig.PrivateKey))
	tigerGroup.POST("/:tigerID/sighting", srv.CreateSighting, middleware.Auth(appConfig.JWTConfig.PrivateKey))

	if appConfig.IsAllowCORS {
		e.OPTIONS("/healthcheck", middleware.HandleOptionsRequest(nhttp.MethodGet))

		userGroup.OPTIONS("/register", middleware.HandleOptionsRequest(nhttp.MethodPost))
		userGroup.OPTIONS("/login", middleware.HandleOptionsRequest(nhttp.MethodPost))

		tigerGroup.OPTIONS("", middleware.HandleOptionsRequest(nhttp.MethodPost, nhttp.MethodGet))
		tigerGroup.OPTIONS("/sighting", middleware.HandleOptionsRequest(nhttp.MethodPost, nhttp.MethodGet))
	}
}

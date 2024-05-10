package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/vekaputra/tiger-kittens/internal/config"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type EchoServer struct {
	Echo *echo.Echo
	Port int
}

func NewEcho(config *config.Config) *EchoServer {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Intercept the 404 handler to prevent using the default error handlers
	// that returns a non nil error causing datadog trace spans to be set
	// to an error state.
	e.RouteNotFound("/*", notFoundHandler)

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var headers map[string][]string
			if config.IsEnableDebug {
				headers = v.Headers
			}

			log.Info().
				Str("http.method", v.Method).
				Str("http.url", v.URI).
				Int("http.status_code", v.Status).
				Str("http.url_details.host", c.Request().Host).
				Str("http.url_details.path", c.Path()).
				Str("http.url_details.queryString", c.Request().URL.Query().Encode()).
				Str("http.url_details.scheme", c.Scheme()).
				Dur("duration", v.Latency).
				Interface("request_headers", headers).
				Msg("incoming request")

			return nil
		},
	}))

	return &EchoServer{
		Echo: e,
		Port: config.Port,
	}
}

func (s *EchoServer) Serve() error {
	log.Info().Msg(fmt.Sprintf("starting server at port: %d", s.Port))

	httpPort := fmt.Sprintf(":%d", s.Port)

	srv := http.Server{
		Addr:         httpPort,
		Handler:      s.Echo,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  660 * time.Second,
	}
	srv.SetKeepAlivesEnabled(false)
	return srv.ListenAndServe()
}

func (s *EchoServer) GracefulStop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.Echo.Shutdown(ctx)
}

func notFoundHandler(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "Not Found"})
}

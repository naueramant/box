package rest

import (
	"box/internal/rest/routes"
	"box/internal/rest/validator"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	gzipCompressionLevel = 5
)

type Config struct {
	Port uint32 `required:"true"`
}

type Server struct {
	echo   *echo.Echo
	config Config
}

func NewServer(conf Config, logger *logrus.Logger) (*Server, error) {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Validator = validator.New()

	l := logger.WithFields(logrus.Fields{
		"subsystem": "REST",
	})

	e.Use(
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level:   gzipCompressionLevel,
			Skipper: middleware.DefaultGzipConfig.Skipper,
		}),
		middleware.RequestID(),
	)

	root := e.Group("/api")

	routes.Register(root, l)

	return &Server{
		echo:   e,
		config: conf,
	}, nil
}

func (s *Server) Start() error {
	return errors.Wrap(s.echo.Start(fmt.Sprintf(":%d", s.config.Port)), "Failed to start server")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return errors.Wrap(s.echo.Shutdown(ctx), "Failed to shutdown server")
}

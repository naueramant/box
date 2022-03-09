package routes

import (
	"box/internal/authentication"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Container struct {
	Authentication authentication.Service
}

type HandlerFunction func(ctx echo.Context, l *logrus.Entry) error

func Register(e *echo.Group, l *logrus.Entry) *Container {
	c := &Container{
		Authentication: authentication.New(),
	}

	authGroup := e.Group("/authentication")
	authGroup.POST("/login", wrap(c.PostLogin, l))
	authGroup.POST("/refresh", wrap(c.PostRefresh, l))

	return c
}

func wrap(fn HandlerFunction, logger *logrus.Entry) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		l := logger.WithFields(logrus.Fields{
			"method":     ctx.Request().Method,
			"path":       ctx.Request().URL.Path,
			"request_id": ctx.Response().Header().Get("X-Request-ID"),
		})

		return fn(ctx, l)
	}
}

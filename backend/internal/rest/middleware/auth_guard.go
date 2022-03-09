package middleware

import "github.com/labstack/echo/v4"

func AuthenticatedGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

package middlewares

import (
	"github.com/labstack/echo/v4"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/4.0")

		return next(c)
	}
}

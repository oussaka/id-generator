package routes

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	r := echo.New()
	initRoute(r)
	v1 := r.Group("/v1")
	{
		UsersRoute(v1)
		PingRoute(v1)
	}

	return r
}

func initRoute(r *echo.Echo) {
}

func InitEcho() {
}
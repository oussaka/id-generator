package routes

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func New() *echo.Echo {
	r := echo.New()
	initRoute(r)

	v1 := r.Group("/v1")
	{
		AuthRoute(v1)
		PingRoute(v1)
		UsersRoute(v1)
	}

	return r
}

func initRoute(r *echo.Echo) {
}

func InitEcho() {
}

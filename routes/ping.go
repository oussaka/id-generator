package routes

import (
	"id-generator/controllers"
	"id-generator/middlewares"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func PingRoute(router *echo.Group) {

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	router.Use(echojwt.WithConfig(config))

	ping := router.Group("/ping", middlewares.IsLoggedIn, middlewares.ServerHeader)
	{
		ping.GET(
			"",
			controllers.Ping,
		)
	}
}

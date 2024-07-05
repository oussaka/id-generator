package routes

import (
	"id-generator/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoute(router *echo.Group) {
	auth := router.Group("/login")
	{
		auth.GET(
			"",
			controllers.Login,
		)
	}
}

package routes

import (
	"id-generator/controllers"
	"github.com/labstack/echo/v4"
)
		
func PingRoute(router *echo.Group) {
	ping := router.Group("/ping")
	{
		ping.GET(
			"",
			controllers.Ping,
		)
	}
}
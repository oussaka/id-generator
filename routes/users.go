package routes

import (
	"id-generator/controllers"
	"github.com/labstack/echo/v4"
)

func UsersRoute(router *echo.Group) {
	users := router.Group("/users")
	{
		users.POST(
			"",
			controllers.CreateUser,
		)
	}
}
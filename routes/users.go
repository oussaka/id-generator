package routes

import (
	"id-generator/controllers"

	"github.com/labstack/echo/v4"
)

func UsersRoute(router *echo.Group) {
	users := router.Group("/users")
	{
		users.GET(
			"",
			controllers.GetUsers,
		)

		users.GET(
			"/:uid",
			controllers.GetUser,
		)

		users.POST(
			"",
			controllers.CreateUser,
		)

		users.PUT(
			"/:uid",
			controllers.UpdateUser,
		)

		users.DELETE(
			"/:uid",
			controllers.DeleteUser,
		)
	}
}

package controllers

import (
	"id-generator/models"
	"id-generator/repositories"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//----------
// Handlers
//----------
func GetUsers(c echo.Context) error {
	fmt.Println("GETUsers controller")

	// query all data
	fmt.Println("== query all data ==")

	users, err := repositories.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	fmt.Println("GETUser controller")

	uid := c.Param("uid")
	// query all data
	fmt.Println("== query all data ==")

	user, err := repositories.GetUser(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	
	fmt.Println("createUser controller")
	fmt.Printf("%#v", user)

	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	uid := c.Param("uid")

	user := models.User{}
	c.Bind(&user)
	updatedUser, err := repositories.UpdateUser(user, uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c echo.Context) error {
	uid := c.Param("uid")
	fmt.Println(uid)

	user := models.User{}
	c.Bind(&user)
	deletedUser, err := repositories.DeleteUser(user, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, deletedUser)
}
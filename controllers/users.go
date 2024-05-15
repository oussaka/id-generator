package controllers

import (
	"id-generator/models"
	"id-generator/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//----------
// Handlers
//----------
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
package controllers

import (
	"github.com/labstack/echo/v4"
	"id-generator/models"
	"net/http"
)

// Ping godoc
func Ping(c echo.Context) error {
	response := &models.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "pong",
	}

	return c.JSON(http.StatusOK, response)
}
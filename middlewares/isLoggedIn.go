package middlewares

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		log.Printf("dans middleware isLogged")

		_, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return errors.New("JWT token missing or invalid")
		}

		return next(c)
	}
}

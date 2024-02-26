package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClams struct {
	UserId float64 `json:"userId"`
	Role   string  `json:"role"`
}

func GetClamsJwt(c echo.Context) JwtClams {
	userJwt := c.Get("user")
	data := userJwt.(jwt.MapClaims)

	return JwtClams{
		UserId: data["userId"].(float64),
		Role:   data["role"].(string),
	}
}

package handler

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Respone struct {
	Status  int         `json:"status"`
	Massage string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func jwtConfig() echojwt.Config {
	tokensecret := viper.GetString("jwt_secret")
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey: []byte(tokensecret),
	}
}

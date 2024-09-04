package handler

import (
	"order-service-gb1/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kodinggo/user-service-gb1/pb/auth"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Respone struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
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

type cartsHandler struct {
	authClient    auth.JWTValidatorClient
	cartsServices model.ICartsServices
}

func NewcartsHandler() *cartsHandler {
	return new(cartsHandler)
}

func (h *cartsHandler) RegisterCartsServices(carts model.ICartsServices) {
	h.cartsServices = carts
}

func (h *cartsHandler) RegisterAuthClient(auth auth.JWTValidatorClient) {
	h.authClient = auth
}

func (h *cartsHandler) Routes(route *echo.Echo, auth echo.MiddlewareFunc) {
	v1 := route.Group("/api/v1")

	// private routes goes here
	routes := v1.Group("")
	routes.Use(auth)
	carts := routes.Group("/carts")

	carts.POST("", h.AddToCarts)
	carts.GET("", h.FindAllCarts)

}

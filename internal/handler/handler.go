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

type handler struct {
	authClient    auth.JWTValidatorClient
	cartsServices model.ICartsServices
	orderService  model.OrderService
}

func NewHandler() *handler {
	return new(handler)
}

func (h *handler) RegisterCartsServices(carts model.ICartsServices) {
	h.cartsServices = carts
}

func (h *handler) RegisterAuthClient(auth auth.JWTValidatorClient) {
	h.authClient = auth
}

func (h *handler) RegisterOrderService(order model.OrderService) {
	h.orderService = order
}

func (h *handler) Routes(route *echo.Echo, auth echo.MiddlewareFunc) {
	v1 := route.Group("/api/v1")

	// private routes goes here
	routes := v1.Group("")
	routes.Use(auth)

	carts := routes.Group("/carts")
	carts.POST("", h.AddToCarts)
	carts.GET("", h.FindAllCarts)

	orders := routes.Group("/orders")
	orders.POST("", h.Create)
}

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

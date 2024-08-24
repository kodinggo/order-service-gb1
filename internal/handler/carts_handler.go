package handler

import (
	"net/http"
	"order-service-gb1/internal/model"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CartsHandler struct {
	CartsServices model.ICartsServices
}

type NewStoryHandler struct {
	CartsServices model.ICartsServices
}

func NewCartsRepository(e *echo.Group, repo model.ICartsServices) {
	handler := &CartsHandler{
		CartsServices: repo,
	}
	carts := e.Group("/carts")
	carts.Use(echojwt.WithConfig(jwtConfig()))

	carts.POST("/addtocarts", handler.AddToCarts)
}

func (s *CartsHandler) AddToCarts(c echo.Context) error {
	var input model.CartsInput

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Respone{
			Status:  http.StatusBadRequest,
			Massage: err.Error(),
			Data:    nil,
		})
	}
	responeService, errServices := s.CartsServices.AddTocarts(input)
	if errServices != nil {
		return c.JSON(http.StatusBadRequest, Respone{
			Status:  http.StatusBadRequest,
			Massage: errServices.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, Respone{
		Status:  http.StatusOK,
		Massage: "success",
		Data:    responeService,
	})
}

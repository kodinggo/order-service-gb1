package handler

import (
	"net/http"
	"order-service-gb1/internal/model"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CartsHandler struct {
	CartsServices model.ICartsServices
}

type NewCartsHandler struct {
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
	log := logrus.WithContext(c.Request().Context())

	err := c.Bind(&input)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, Respone{
			Status:  http.StatusBadRequest,
			Massage: "invalid request body",
			Data:    nil,
		})
	}

	responeService, errServices := s.CartsServices.AddTocarts(c.Request().Context(), input)
	if errServices != nil {
		log.Error(errServices)

		return c.JSON(http.StatusInternalServerError, Respone{
			Status:  http.StatusInternalServerError,
			Massage: "internal server error",
			Data:    nil,
		})

	}
	return c.JSON(http.StatusOK, Respone{
		Status:  http.StatusOK,
		Massage: "success",
		Data:    responeService,
	})
}

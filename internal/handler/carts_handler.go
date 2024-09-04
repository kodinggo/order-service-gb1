package handler

import (
	"net/http"
	"order-service-gb1/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *cartsHandler) AddToCarts(c echo.Context) error {
	carts := model.CartsInput{}
	log := logrus.WithContext(c.Request().Context())

	err := c.Bind(&carts)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, Respone{
			Status:  http.StatusBadRequest,
			Message: "invalid request body binding",
		})
	}

	if err := carts.Validator(); err != nil {
		return c.JSON(http.StatusBadRequest, Respone{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	responeService, err := s.cartsServices.AddTocarts(c.Request().Context(), carts)
	if err != nil {
		log.Error(err)

		return c.JSON(http.StatusInternalServerError, Respone{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		})

	}
	return c.JSON(http.StatusOK, Respone{
		Status:  http.StatusOK,
		Message: "success",
		Data:    responeService,
	})
}

func (s *cartsHandler) FindAllCarts(c echo.Context) error {
	log := logrus.WithContext(c.Request().Context())

	carts, err := s.cartsServices.FindAllCarts(c.Request().Context())
	if err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, Respone{
		Status:  http.StatusOK,
		Message: "success",
		Data:    carts,
	})
}

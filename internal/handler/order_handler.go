package handler

import (
	"net/http"
	"order-service-gb1/internal/model"
	"order-service-gb1/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *handler) Create(c echo.Context) error {
	order := model.Order{}

	err := c.Bind(&order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: "invalid request body binding",
		})
	}

	logger := logrus.WithFields(logrus.Fields{
		"order": utils.Dump(order),
	})

	err = s.orderService.Create(c.Request().Context(), order)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "success",
	})
}

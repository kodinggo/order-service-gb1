package handler

import (
	"net/http"
	"order-service-gb1/internal/model"
	"order-service-gb1/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type orderHandler struct {
	orderService model.OrderService
}

func NewOrderHandler(e *echo.Group, service model.OrderService) {
	handler := &orderHandler{
		orderService: service,
	}

	order := e.Group("/orders")

	order.POST("", handler.Create)
}

func (s *orderHandler) Create(c echo.Context) error {
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

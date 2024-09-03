package handler

import (
	"net/http"
	"order-service-gb1/internal/model"

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
	//carts.Use(echojwt.WithConfig(jwtConfig()))

	carts.POST("/addtocarts", handler.AddToCarts)
	carts.GET("/findall", handler.FindAllCarts)
}

func (s *CartsHandler) AddToCarts(c echo.Context) error {
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

	responeService, err := s.CartsServices.AddTocarts(c.Request().Context(), carts)
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

func (s *CartsHandler) FindAllCarts(c echo.Context) error {
	log := logrus.WithContext(c.Request().Context())

	carts, err := s.CartsServices.FindAllCarts(c.Request().Context())
	if err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, Respone{
		Status:  http.StatusOK,
		Message: "success",
		Data:    carts,
	})
}

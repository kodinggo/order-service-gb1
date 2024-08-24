package services

import (
	"order-service-gb1/internal/model"

	"github.com/sirupsen/logrus"
)

type CartsServices struct {
	cartsRepo model.ICartsRepository
}

func NewCartsRepository(repo model.ICartsRepository) model.ICartsServices {
	return &CartsServices{cartsRepo: repo}
}

func (s *CartsServices) AddTocarts(input model.CartsInput) (model.CartsRespone, error) {
	log := logrus.WithFields(logrus.Fields{
		"carts": input,
	})
	carts := model.Carts{
		UserID:    input.UserID,
		ProductID: input.ProductID,
	}

	respone, err := s.cartsRepo.AddToCarts(carts)
	if err != nil {
		log.Error(err)
	}
	return respone, nil

}

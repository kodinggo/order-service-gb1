package services

import (
	"context"
	"order-service-gb1/internal/model"

	"github.com/sirupsen/logrus"
)

type CartsServices struct {
	cartsRepo model.ICartsRepository
}

func NewCartsRepository(repo model.ICartsRepository) model.ICartsServices {
	return &CartsServices{cartsRepo: repo}
}

func (s *CartsServices) AddTocarts(ctx context.Context, input model.CartsInput) (model.CartsRespone, error) {
	log := logrus.WithFields(logrus.Fields{
		"carts": input,
	})
	carts := model.Carts{
		UserID:    input.UserID,
		ProductID: input.ProductID,
	}

	respone, err := s.cartsRepo.AddToCarts(ctx, carts)
	if err != nil {
		log.Error(err)
		return model.CartsRespone{}, err
	}
	return respone, nil

}

func (s *CartsServices) FindAllCarts(ctx context.Context) ([]model.Carts, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})
	var carts []model.Carts
	carts, err := s.cartsRepo.FindAllCarts(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return carts, nil
}

package repository

import (
	"order-service-gb1/internal/model"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CartsRepository struct {
	db *gorm.DB
}

func NewCartsRepository(db *gorm.DB) *CartsRepository {
	return &CartsRepository{
		db: db,
	}
}

func (r *CartsRepository) AddToCarts(input model.Carts) (model.CartsRespone, error) {
	log := logrus.WithFields(logrus.Fields{
		"carts": input,
	})
	err := r.db.Select("user_id", "product_id").Create(&input).Error
	if err != nil {
		log.Error(err)
		return model.CartsRespone{}, err
	}

	return model.CartsRespone{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		CreatedAt: time.Now(),
	}, nil
}

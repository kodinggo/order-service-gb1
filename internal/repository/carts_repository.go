package repository

import (
	"order-service-gb1/internal/model"
	"time"

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

	r.db.Select("user_id", "product_id").Create(&input)

	return model.CartsRespone{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		CreatedAt: time.Now(),
	}, nil
}

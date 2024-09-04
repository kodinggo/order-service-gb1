package repository

import (
	"context"
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

func (r *CartsRepository) AddToCarts(ctx context.Context, input model.Carts) (model.CartsRespone, error) {
	log := logrus.WithFields(logrus.Fields{
		"carts": input,
	})
	err := r.db.WithContext(ctx).Select("user_id", "product_id").Create(&input).Error
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

func (r *CartsRepository) FindAllCarts(ctx context.Context) ([]model.Carts, error) {
	log := logrus.WithFields(logrus.Fields{
		"carts": ctx,
	})
	var carts []model.Carts
	err := r.db.WithContext(ctx).Find(&carts).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return carts, nil
}

func (r *CartsRepository) FindByUserID(ctx context.Context, userID int) ([]model.Carts, error) {
	log := logrus.WithFields(logrus.Fields{
		"user_id": userID,
	})

	var carts []model.Carts

	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&carts).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return carts, nil
}

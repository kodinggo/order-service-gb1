package repository

import (
	"context"
	"order-service-gb1/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) model.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(ctx context.Context, order model.Order) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":   ctx,
		"order": order,
	})

	// using sql transaction
	// to ensure all data is saved
	tx := r.db.Begin()

	err := tx.Create(&order).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	err = tx.Create(&model.OrderLog{OrderID: order.ID, Status: order.Status}).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	// commit the transaction if all data is saved
	tx.Commit()

	return nil
}

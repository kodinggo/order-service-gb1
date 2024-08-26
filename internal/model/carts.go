package model

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ICartsRepository interface {
	AddToCarts(ctx context.Context, input Carts) (CartsRespone, error)
}

type ICartsServices interface {
	AddTocarts(ctx context.Context, input CartsInput) (CartsRespone, error)
}

type Carts struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	ProductID int            `json:"product_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deteted_at"`
}

type CartsInput struct {
	UserID    int `json:"user_id" validate:"required"`
	ProductID int `json:"product_id" validate:"required"`
}

type CartsRespone struct {
	UserID    int       `json:"user_id" validate:"required"`
	ProductID int       `json:"product_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (c CartsInput) Validator() error {
	if c.UserID == 0 || c.ProductID == 0 {
		return errors.New("invalid request body ,user_id or product_id cannot be empty")
	}

	return nil
}

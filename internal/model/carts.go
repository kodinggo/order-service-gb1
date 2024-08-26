package model

import (
	"time"

	"gorm.io/gorm"
)

type ICartsRepository interface {
	AddToCarts(input Carts) (CartsRespone, error)
}

type ICartsServices interface {
	AddTocarts(input CartsInput) (CartsRespone, error)
}

type Carts struct {
	ID        *int           `json:"id" gorm:"primaryKey"`
	UserID    *int           `json:"user_id",omitempty"`
	ProductID *int           `json:"product_id,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index, omitempty"`
}

type CartsInput struct {
	UserID    *int `json:"user_id"" validate:"required"`
	ProductID *int `json:"product_id" validate:"required"`
}

type CartsRespone struct {
	UserID    *int      `json:"user_id"" validate:"required"`
	ProductID *int      `json:"product_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

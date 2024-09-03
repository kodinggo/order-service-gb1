package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID              int            `json:"id"`
	UserID          int            `json:"user_id"`
	InvoiceNo       string         `json:"invoice_no"`
	GrandTotal      float64        `json:"grand_total"`
	Status          string         `json:"status"`
	ShippingAddress string         `json:"shipping_address"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty"`
	OrderItems      []OrderItem    `json:"order_items" gorm:"foreignKey:OrderID"`
}

func (o *Order) SetGrandTotal() {
	var total float64
	for _, item := range o.OrderItems {
		total += item.SubTotal
	}

	o.GrandTotal = total
}

type OrderItem struct {
	ID           int            `json:"id"`
	OrderID      int            `json:"order_id"`
	ProductID    int            `json:"product_id"`
	ProductName  string         `json:"product_name"`
	ProductPrice float64        `json:"product_price"`
	Qty          int            `json:"qty"`
	SubTotal     float64        `json:"sub_total"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type OrderLog struct {
	ID        int            `json:"id"`
	OrderID   int            `json:"order_id"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type OrderRepository interface {
	Create(ctx context.Context, order Order) error
}

type OrderService interface {
	Create(ctx context.Context, order Order) error
}

package model

import (
	"context"

	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ProductID uint `gorm:"type:int;not null"`
	Quantity  uint `gorm:"type:int;not null"`
	OrderID   uint `gorm:"index;type:int;not null"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

type OrderItemQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewOrderItemQuery(ctx context.Context, db *gorm.DB) *OrderItemQuery {
	return &OrderItemQuery{
		ctx: ctx,
		db:  db,
	}
}

package model

import (
	"context"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint32      `gorm:"index;type:int;not null"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

func (o *Order) TableName() string {
	return "orders"
}

type OrdersQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewOrdersQuery(ctx context.Context, db *gorm.DB) *OrdersQuery {
	return &OrdersQuery{ctx: ctx, db: db}
}

func (oq *OrdersQuery) GetAll(user_id uint32) ([]*Order, error) {
	var orders []*Order
	err := oq.db.WithContext(oq.ctx).Model(&Order{}).Preload("OrderItems").Where("user_id = ?", user_id).Find(&orders).Error
	return orders, err
}
func (oq *OrdersQuery) CreateWithItems(userID uint32, items []OrderItem) (*Order, error) {
	order := &Order{UserID: userID}

	err := oq.db.WithContext(oq.ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		if len(items) > 0 {
			for i := range items {
				items[i].OrderID = uint(order.ID)
			}
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return order, nil
}

package model

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartItem struct {
	gorm.Model
	UserID    uint64 `gorm:"not null;uniqueIndex:uid_pid" json:"user_id"`
	ProductID uint64 `gorm:"not null;uniqueIndex:uid_pid" json:"product_id"`
	Quantity  uint32 `gorm:"not null;default:1" json:"quantity"`
}

func (CartItem) TableName() string { return "cart_items" }

type CartItemQuery struct {
	DB  *gorm.DB
	ctx context.Context
}

func NewCartItemQuery(ctx context.Context, DB *gorm.DB) *CartItemQuery {
	return &CartItemQuery{DB: DB, ctx: ctx}
}

func (q *CartItemQuery) SetCartQ(userID uint64, productID uint64, quantity uint32) error {
	return q.DB.WithContext(q.ctx).
		Model(&CartItem{}).
		Where("user_id=? AND product_id=?", userID, productID).
		Update("quantity", quantity).Error
}

func (q *CartItemQuery) AddOrIncr(userID, productID uint64, addQty uint32) error {
	item := CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  addQty,
	}
	return q.DB.WithContext(q.ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "user_id"},
				{Name: "product_id"},
			},
			DoUpdates: clause.Assignments(map[string]any{
				"quantity": gorm.Expr("quantity + VALUES(quantity)"),
			}),
		}).
		Create(&item).Error
}

func (q *CartItemQuery) AddCartItem(cartItem CartItem) error {
	return q.AddOrIncr(cartItem.UserID, cartItem.ProductID, cartItem.Quantity)
}

func (q *CartItemQuery) DeleteCartItem(userID uint64, productID uint64) error {
	return q.DB.WithContext(q.ctx).
		Where("user_id=? AND product_id=?", userID, productID).
		Delete(&CartItem{}).Error
}

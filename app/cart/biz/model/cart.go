package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint64     `gorm:"not null;uniqueIndex;" json:"user_id"`
	Items  []CartItem `gorm:"foreignKey:UserID;references:UserID"`
}

func (Cart) TableName() string { return "carts" }

type CartQuery struct {
	DB  *gorm.DB
	ctx context.Context
}

func NewCartQuery(ctx context.Context, DB *gorm.DB) *CartQuery {
	return &CartQuery{DB: DB, ctx: ctx}
}

func (q *CartQuery) GetByUser(userID uint64) (Cart, error) {
	var c Cart
	err := q.DB.WithContext(q.ctx).Preload("Items").
		Where("user_id = ?", userID).
		First(&c).Error
	return c, err
}

func (q *CartQuery) GetOrCreateByUser(userID uint64) (Cart, error) {
	var c Cart
	err := q.DB.WithContext(q.ctx).
		Where("user_id = ?", userID).
		First(&c).Error
	if err == nil {
		return c, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c = Cart{UserID: userID}
		if e := q.DB.WithContext(q.ctx).Create(&c).Error; e != nil {
			return Cart{}, e
		}
		return c, nil
	}
	return Cart{}, err
}

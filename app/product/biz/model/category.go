package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	Description string    `gorm:"type:text;not null"`
	Products    []Product `gorm:"many2many:product_category;"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	db  *gorm.DB
	ctx context.Context
}

func (q *CategoryQuery) GetProductByCategory(name string) ([]Category, error) {
	var Categories []Category
	err := q.db.WithContext(q.ctx).Model(&Category{}).Where("name = ?", name).Preload("Products").Find(&Categories).Error
	return Categories, err
}

func NewCategoryQuery(ctx context.Context, DB *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		db:  DB,
		ctx: ctx,
	}
}

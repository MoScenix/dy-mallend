package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Price       float32    `gorm:"type:float;not null" json:"price"`
	Description string     `gorm:"type:text;not null" json:"description"`
	Picture     string     `gorm:"type:varchar(255);not null" json:"picture"`
	UserID      int        `gorm:"type:int;not null;index" json:"user_id"`
	Categories  []Category `gorm:"many2many:product_category;" json:"categories"`
}

func (Product) TableName() string {
	return "products"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (q *ProductQuery) GetById(id int) (Product, error) {
	var product Product
	err := q.db.WithContext(q.ctx).Model(&Product{}).First(&product, id).Error
	return product, err
}
func (q *ProductQuery) SearchProducts(search string) ([]Product, error) {
	var products []Product
	err := q.db.WithContext(q.ctx).Model(&Product{}).Where("name LIKE ? or description like ?", "%"+search+"%", "%"+search+"%").Find(&products).Error
	return products, err
}
func NewProductQuery(ctx context.Context, DB *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  DB,
	}
}
func (q *ProductQuery) GetByUser(id int) ([]Product, error) {
	var Products []Product
	err := q.db.WithContext(q.ctx).Model(&Product{}).Where("user_id = ?", id).Find(&Products).Error
	return Products, err
}
func (q *ProductQuery) AddProduct(product Product) error {
	err := q.db.WithContext(q.ctx).Model(&Product{}).Create(&product).Error
	return err
}
func (q *ProductQuery) DeleteProduct(id int) error {
	err := q.db.WithContext(q.ctx).Model(&Product{}).Delete(&Product{}, id).Error
	return err
}

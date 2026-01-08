package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	Price       float32    `gorm:"type:float;not null" json:"price"`
	Description string     `gorm:"type:text;not null" json:"description"`
	Picture     string     `gorm:"type:varchar(255);not null" json:"picture"`
	UserID      int        `gorm:"type:int;not null;index" json:"user_id"`
	Categories  []Category `gorm:"many2many:product_category;" json:"categories"`
}

func (Product) TableName() string { return "products" }

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, DB *gorm.DB) *ProductQuery {
	return &ProductQuery{ctx: ctx, db: DB}
}

func (q *ProductQuery) GetById(id int) (Product, error) {
	var product Product
	err := q.db.WithContext(q.ctx).Model(&Product{}).First(&product, id).Error
	return product, err
}

func (q *ProductQuery) SearchProducts(search string) ([]Product, error) {
	var products []Product
	err := q.db.WithContext(q.ctx).
		Model(&Product{}).
		Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%").
		Find(&products).Error
	return products, err
}

func (q *ProductQuery) GetByUser(id int) ([]Product, error) {
	var products []Product
	err := q.db.WithContext(q.ctx).
		Model(&Product{}).
		Where("user_id = ?", id).
		Find(&products).Error
	return products, err
}

func (q *ProductQuery) AddProduct(product Product) error {
	return q.db.WithContext(q.ctx).Model(&Product{}).Create(&product).Error
}

func (q *ProductQuery) DeleteProduct(id int) error {
	return q.db.WithContext(q.ctx).Model(&Product{}).Delete(&Product{}, id).Error
}

func (q *ProductQuery) GetProductsById(userID int) ([]Product, error) {
	var products []Product
	err := q.db.WithContext(q.ctx).
		Model(&Product{}).
		Where("user_id = ?", userID).
		Find(&products).Error
	return products, err
}

type ProductProQuery struct {
	q      *ProductQuery
	rdb    *redis.Client
	prefix string
}

func NewProductProQuery(ctx context.Context, db *gorm.DB, rdb *redis.Client) *ProductProQuery {
	return &ProductProQuery{
		q:      NewProductQuery(ctx, db),
		rdb:    rdb,
		prefix: "dy:",
	}
}

func (p *ProductProQuery) keyProduct(id int) string {
	return fmt.Sprintf("%sproduct:%d", p.prefix, id)
}

func (p *ProductProQuery) keyUserProducts(userID int) string {
	return fmt.Sprintf("%suser_products:%d", p.prefix, userID)
}
func (p *ProductProQuery) GetById(id int) (Product, error) {
	if p.rdb != nil {
		if val, err := p.rdb.Get(p.q.ctx, p.keyProduct(id)).Result(); err == nil && val != "" {
			var prod Product
			if json.Unmarshal([]byte(val), &prod) == nil {
				return prod, nil
			}
		}
	}
	prod, err := p.q.GetById(id)
	if err != nil {
		return Product{}, err
	}
	if p.rdb != nil {
		if b, e := json.Marshal(prod); e == nil {
			_ = p.rdb.Set(p.q.ctx, p.keyProduct(id), b, time.Hour).Err()
		}
	}
	return prod, nil
}
func (p *ProductProQuery) GetByUser(userID int) ([]Product, error) {
	if p.rdb != nil {
		if val, err := p.rdb.Get(p.q.ctx, p.keyUserProducts(userID)).Result(); err == nil && val != "" {
			var list []Product
			if json.Unmarshal([]byte(val), &list) == nil {
				return list, nil
			}
		}
	}

	list, err := p.q.GetByUser(userID)
	if err != nil {
		return nil, err
	}

	if p.rdb != nil {
		if b, e := json.Marshal(list); e == nil {
			_ = p.rdb.Set(p.q.ctx, p.keyUserProducts(userID), b, time.Hour).Err()
		}
	}
	return list, nil
}
func (p *ProductProQuery) GetProductsById(userID int) ([]Product, error) {
	return p.GetByUser(userID)
}
func (p *ProductProQuery) SearchProducts(search string) ([]Product, error) {
	return p.q.SearchProducts(search)
}
func (p *ProductProQuery) AddProduct(product Product) error {
	if err := p.q.AddProduct(product); err != nil {
		return err
	}
	if p.rdb != nil {
		_ = p.rdb.Del(p.q.ctx, p.keyUserProducts(product.UserID)).Err()
	}
	return nil
}
func (p *ProductProQuery) DeleteProduct(id int) error {
	var uid int
	if prod, err := p.GetById(id); err == nil {
		uid = prod.UserID
	}

	if err := p.q.DeleteProduct(id); err != nil {
		return err
	}

	if p.rdb != nil {
		_ = p.rdb.Del(p.q.ctx, p.keyProduct(id)).Err()
		if uid != 0 {
			_ = p.rdb.Del(p.q.ctx, p.keyUserProducts(uid)).Err()
		}
	}
	return nil
}
func (q *ProductQuery) GetByIDsWithDeleted(ids []uint32) ([]Product, error) {
	var products []Product
	if len(ids) == 0 {
		return products, nil
	}

	err := q.db.WithContext(q.ctx).
		Unscoped().
		Model(&Product{}).
		Where("id IN ?", ids).
		Find(&products).Error
	return products, err
}

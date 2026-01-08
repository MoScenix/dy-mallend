package service

import (
	"context"
	"fmt"

	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/model"
	cart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
	"gorm.io/gorm"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	if req == nil || req.UserId == 0 || req.ProductId == 0 || req.Quantity == 0 {
		return &cart.AddItemResp{Success: false}, fmt.Errorf("invalid params")
	}
	if mysql.DB == nil {
		return &cart.AddItemResp{Success: false}, fmt.Errorf("db not initialized")
	}
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		cq := model.NewCartQuery(s.ctx, tx)
		if _, e := cq.GetOrCreateByUser(req.UserId); e != nil {
			return e
		}
		ciq := model.NewCartItemQuery(s.ctx, tx)
		if e := ciq.AddOrIncr(req.UserId, req.ProductId, req.Quantity); e != nil {
			return e
		}
		return nil
	})

	if err != nil {
		return &cart.AddItemResp{Success: false}, err
	}
	return &cart.AddItemResp{Success: true}, nil
}

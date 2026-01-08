package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/model"
	cart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
)

type SetItemService struct {
	ctx context.Context
} // NewSetItemService new SetItemService
func NewSetItemService(ctx context.Context) *SetItemService {
	return &SetItemService{ctx: ctx}
}

// Run create note info
func (s *SetItemService) Run(req *cart.SetItemReq) (resp *cart.SetItemResp, err error) {
	// Finish your business logic.
	cartItemQuery := model.NewCartItemQuery(s.ctx, mysql.DB)

	if req.Quantity > 0 {
		err = cartItemQuery.SetCartQ(req.UserId, req.ProductId, req.Quantity)
	} else {
		err = cartItemQuery.DeleteCartItem(req.UserId, req.ProductId)
	}
	success := false
	if err == nil {
		success = true
	}

	return &cart.SetItemResp{
		Success: success,
	}, err
}

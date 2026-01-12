package service

import (
	"context"
	orders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
)

type AddOrderService struct {
	ctx context.Context
} // NewAddOrderService new AddOrderService
func NewAddOrderService(ctx context.Context) *AddOrderService {
	return &AddOrderService{ctx: ctx}
}

// Run create note info
func (s *AddOrderService) Run(req *orders.AddOrderReq) (resp *orders.AddOrderResp, err error) {
	// Finish your business logic.

	return
}

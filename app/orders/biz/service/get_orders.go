package service

import (
	"context"
	orders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
)

type GetOrdersService struct {
	ctx context.Context
} // NewGetOrdersService new GetOrdersService
func NewGetOrdersService(ctx context.Context) *GetOrdersService {
	return &GetOrdersService{ctx: ctx}
}

// Run create note info
func (s *GetOrdersService) Run(req *orders.GetOrdersReq) (resp *orders.GetOrdersResp, err error) {
	// Finish your business logic.

	return
}

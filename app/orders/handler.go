package main

import (
	"context"
	orders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/service"
)

// OrdersServiceImpl implements the last service interface defined in the IDL.
type OrdersServiceImpl struct{}

// GetOrders implements the OrdersServiceImpl interface.
func (s *OrdersServiceImpl) GetOrders(ctx context.Context, req *orders.GetOrdersReq) (resp *orders.GetOrdersResp, err error) {
	resp, err = service.NewGetOrdersService(ctx).Run(req)

	return resp, err
}

// AddOrder implements the OrdersServiceImpl interface.
func (s *OrdersServiceImpl) AddOrder(ctx context.Context, req *orders.AddOrderReq) (resp *orders.AddOrderResp, err error) {
	resp, err = service.NewAddOrderService(ctx).Run(req)

	return resp, err
}

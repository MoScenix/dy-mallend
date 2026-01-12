package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/model"
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
	ordersQuery := model.NewOrdersQuery(s.ctx, mysql.DB)
	OrderItems := []model.OrderItem{}
	for _, item := range req.Items {
		OrderItems = append(OrderItems, model.OrderItem{
			ProductID: uint(item.ProductId),
			Quantity:  uint(item.Quantity),
		})
	}
	_, err = ordersQuery.CreateWithItems(req.UserId, OrderItems)
	return &orders.AddOrderResp{}, err
}

package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/model"
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
	ordersQuery := model.NewOrdersQuery(s.ctx, mysql.DB)
	Orders, err := ordersQuery.GetAll(req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &orders.GetOrdersResp{
		Orders: []*orders.Order{},
	}
	for _, order := range Orders {
		Order := &orders.Order{
			Items: []*orders.OrderItem{},
		}
		for _, item := range order.OrderItems {
			Order.Items = append(Order.Items, &orders.OrderItem{
				ProductId: uint32(item.ProductID),
				Quantity:  uint32(item.Quantity),
			})
		}
		resp.Orders = append(resp.Orders, Order)
	}
	return resp, nil
}

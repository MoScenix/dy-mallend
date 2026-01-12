package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	common "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/common"
	orders "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/orders"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	cart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
	rpcorders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddOrderService(Context context.Context, RequestContext *app.RequestContext) *AddOrderService {
	return &AddOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *AddOrderService) Run(req *orders.Order) (resp *common.Empty, err error) {
	add := &rpcorders.AddOrderReq{
		UserId: uint32(h.Context.Value(utils.UserIdKey).(float64)),
	}
	for _, v := range req.Items {
		add.Items = append(add.Items, &rpcorders.OrderItem{
			ProductId: uint32(v.ProductId),
			Quantity:  uint32(v.Quantity),
		})
	}
	for _, v := range req.Items {
		_, err = rpc.CartClient.SetItem(h.Context, &cart.SetItemReq{
			UserId:    uint64(h.Context.Value(utils.UserIdKey).(float64)),
			ProductId: uint64(v.ProductId),
			Quantity:  uint32(0),
		})
	}
	_, err = rpc.OrdersClient.AddOrder(h.Context, add)
	return &common.Empty{}, err
}

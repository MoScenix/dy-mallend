package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	common "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/common"
	orders "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/orders"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	rpcorders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
	rpcproduct "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetOrdersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetOrdersService(Context context.Context, RequestContext *app.RequestContext) *GetOrdersService {
	return &GetOrdersService{RequestContext: RequestContext, Context: Context}
}

func (h *GetOrdersService) Run(req *common.Empty) (resp map[string]any, err error) {

	res, err := rpc.OrdersClient.GetOrders(h.Context, &rpcorders.GetOrdersReq{
		UserId: uint32(h.Context.Value(utils.UserIdKey).(float64)),
	})
	if err != nil {
		return
	}
	var r []*orders.Order
	for _, v := range res.Orders {
		var q []uint32
		var sum float64

		m := make(map[int]int)
		for _, v := range v.Items {
			q = append(q, uint32(v.ProductId))
			m[int(v.ProductId)] = int(v.Quantity)
		}
		res2, err := rpc.ProductClient.GetProductsById(h.Context, &rpcproduct.GetProductsByIdReq{
			Ids: q,
		})
		if err != nil {
			return nil, err
		}

		var Items []*orders.OrderItem
		for _, r := range res2.Products {
			Items = append(Items, &orders.OrderItem{
				ProductId: uint32(r.Pid),
				Quantity:  uint32(m[int(r.Pid)]),
				Picture:   r.Picture,
				Name:      r.Name,
				Price:     r.Price,
			})
			sum += float64(r.Price) * float64(m[int(r.Pid)])
		}
		r = append(r, &orders.Order{
			Items: Items,
			Sum:   float32(sum),
		})
	}
	return map[string]any{
		"orders": r,
	}, nil
}

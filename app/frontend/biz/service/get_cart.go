package service

import (
	"context"
	"fmt"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	cart "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/cart"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	rpccart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.GetCartReq) (resp map[string]any, err error) {
	row, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint64(h.Context.Value(utils.UserIdKey).(float64))})
	if err != nil {
		fmt.Println(err)
		return
	}
	var p []uint32
	m := make(map[int]int)
	for _, v := range row.Cart.Items {
		p = append(p, uint32(v.ProductId))
		m[int(v.ProductId)] = int(v.Quantity)
	}
	rpcp, err := rpc.ProductClient.GetProductsById(h.Context, &product.GetProductsByIdReq{
		Ids: p,
	})
	if err != nil {
		return nil, err
	}
	Items := []*cart.CartItem{}
	for _, v := range rpcp.Products {
		Items = append(Items, &cart.CartItem{
			ProductId:   uint64(v.Pid),
			Quantity:    uint32(m[int(v.Pid)]),
			Picture:     v.Picture,
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			Isdeleted:   v.Isdeleted,
		})
	}
	return map[string]any{
		"cart": map[string]any{
			"items": Items,
		},
	}, nil
}

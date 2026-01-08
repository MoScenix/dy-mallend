package service

import (
	"context"
	"os"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	order "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/order"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderProdectDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderProdectDeleteService(Context context.Context, RequestContext *app.RequestContext) *OrderProdectDeleteService {
	return &OrderProdectDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderProdectDeleteService) Run(req *order.DeleteReq) (resp map[string]any, err error) {
	res, err := rpc.ProductClient.DeleteProduct(h.Context, &product.DeleteProductReq{
		UserId: uint32(h.Context.Value(utils.UserIdKey).(float64)),
		Id:     req.Id,
	})
	if err != nil {
		return map[string]any{
			"success": false,
		}, err
	}
	err = os.Remove("." + res.Picture)
	return map[string]any{
		"success": res.Success,
	}, err
}

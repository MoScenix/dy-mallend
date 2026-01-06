package service

import (
	"context"

	product "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/product"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	rpcproduct "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	res, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"item": res.Product,
	}, nil
}

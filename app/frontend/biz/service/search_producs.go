package service

import (
	"context"
	"fmt"

	product "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/product"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	rpcproduct "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProducsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProducsService(Context context.Context, RequestContext *app.RequestContext) *SearchProducsService {
	return &SearchProducsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProducsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	res, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return map[string]any{
		"products": res.Results,
	}, err
}

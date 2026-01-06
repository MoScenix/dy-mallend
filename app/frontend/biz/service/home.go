package service

import (
	"context"
	"fmt"

	common "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/common"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	rpcproduct "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	res, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: "",
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return map[string]any{
		"items": res.Results,
	}, err
}

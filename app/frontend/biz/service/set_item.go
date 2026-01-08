package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	cart "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/cart"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	rpccart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type SetItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSetItemService(Context context.Context, RequestContext *app.RequestContext) *SetItemService {
	return &SetItemService{RequestContext: RequestContext, Context: Context}
}

func (h *SetItemService) Run(req *cart.SetItemReq) (resp *cart.BoolResp, err error) {
	rpcresp, err := rpc.CartClient.SetItem(h.Context, &rpccart.SetItemReq{
		UserId:    uint64(h.Context.Value(utils.UserIdKey).(float64)),
		ProductId: uint64(req.ProductId),
		Quantity:  uint32(req.Quantity),
	})
	if err != nil {
		return nil, err
	}
	return &cart.BoolResp{Success: rpcresp.Success}, nil
}

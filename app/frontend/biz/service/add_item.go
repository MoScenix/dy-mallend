package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	cart "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/cart"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	rpccart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddItemService(Context context.Context, RequestContext *app.RequestContext) *AddItemService {
	return &AddItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddItemService) Run(req *cart.AddItemReq) (resp *cart.BoolResp, err error) {

	rpcReq := &rpccart.AddItemReq{
		UserId:    uint64(h.Context.Value(utils.UserIdKey).(float64)),
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}
	rpcResp, err := rpc.CartClient.AddItem(h.Context, rpcReq)
	if err != nil {
		return nil, err
	}
	return &cart.BoolResp{Success: rpcResp.Success}, nil
}

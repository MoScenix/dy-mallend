package service

import (
	"context"

	cart "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type RemoveItemsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRemoveItemsService(Context context.Context, RequestContext *app.RequestContext) *RemoveItemsService {
	return &RemoveItemsService{RequestContext: RequestContext, Context: Context}
}

func (h *RemoveItemsService) Run(req *cart.RemoveItemsReq) (resp *cart.BoolResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}

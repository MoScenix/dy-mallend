package orders

import (
	"context"
	orders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetOrders(ctx context.Context, req *orders.GetOrdersReq, callOptions ...callopt.Option) (resp *orders.GetOrdersResp, err error) {
	resp, err = defaultClient.GetOrders(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrders call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddOrder(ctx context.Context, req *orders.AddOrderReq, callOptions ...callopt.Option) (resp *orders.AddOrderResp, err error) {
	resp, err = defaultClient.AddOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

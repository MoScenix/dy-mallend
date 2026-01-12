package orders

import (
	"context"
	orders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"

	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders/ordersservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() ordersservice.Client
	Service() string
	GetOrders(ctx context.Context, Req *orders.GetOrdersReq, callOptions ...callopt.Option) (r *orders.GetOrdersResp, err error)
	AddOrder(ctx context.Context, Req *orders.AddOrderReq, callOptions ...callopt.Option) (r *orders.AddOrderResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := ordersservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient ordersservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() ordersservice.Client {
	return c.kitexClient
}

func (c *clientImpl) GetOrders(ctx context.Context, Req *orders.GetOrdersReq, callOptions ...callopt.Option) (r *orders.GetOrdersResp, err error) {
	return c.kitexClient.GetOrders(ctx, Req, callOptions...)
}

func (c *clientImpl) AddOrder(ctx context.Context, Req *orders.AddOrderReq, callOptions ...callopt.Option) (r *orders.AddOrderResp, err error) {
	return c.kitexClient.AddOrder(ctx, Req, callOptions...)
}

package service

import (
	"context"
	"testing"
	orders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
)

func TestGetOrders_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetOrdersService(ctx)
	// init req and assert value

	req := &orders.GetOrdersReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

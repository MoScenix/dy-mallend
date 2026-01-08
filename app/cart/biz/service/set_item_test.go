package service

import (
	"context"
	"testing"
	cart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
)

func TestSetItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSetItemService(ctx)
	// init req and assert value

	req := &cart.SetItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

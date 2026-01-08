package service

import (
	"context"
	"testing"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
)

func TestGetProductsById_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductsByIdService(ctx)
	// init req and assert value

	req := &product.GetProductsByIdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

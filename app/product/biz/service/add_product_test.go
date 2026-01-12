package service

import (
	"context"
	"testing"

	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
)

func TestAddProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddProductService(ctx)
	// init req and assert value

	req := &product.AddProductReq{
		Name:        "test",
		Price:       1.0,
		Picture:     "test",
		Description: "test",
		UserId:      1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}

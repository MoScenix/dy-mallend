package service

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/joho/godotenv"
)

func TestGetProduct_Run(t *testing.T) {
	_, thisFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(thisFile)
	target := filepath.Clean(filepath.Join(baseDir, "../../"))

	_ = os.Chdir(target)
	godotenv.Load()
	dal.Init()

	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value
	req := &product.GetProductReq{
		Id: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

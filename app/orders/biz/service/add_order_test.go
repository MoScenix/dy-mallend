package service

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/MoScenix/douyin-mall-backend/app/orders/biz/dal"
	orders "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/orders"
	"github.com/joho/godotenv"
)

func TestAddOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddOrderService(ctx)
	// init req and assert value
	_, thisFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(thisFile)
	target := filepath.Clean(filepath.Join(baseDir, "../../"))

	_ = os.Chdir(target)
	godotenv.Load()
	dal.Init()
	req := &orders.AddOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

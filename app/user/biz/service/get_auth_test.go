package service

import (
	"context"
	"testing"

	user "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
)

func TestGetAuth_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetAuthService(ctx)
	// init req and assert value

	req := &user.GetAuthReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

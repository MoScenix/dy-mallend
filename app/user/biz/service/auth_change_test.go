package service

import (
	"context"
	"testing"
	user "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
)

func TestAuthChange_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAuthChangeService(ctx)
	// init req and assert value

	req := &user.AuthChangeReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

package service

import (
	"context"
	"testing"

	user "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestRegisterService_Run(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req     *user.RegisterReq
		want    *user.RegisterResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewRegisterService(context.Background())
			got, gotErr := s.Run(tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Run() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Run() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

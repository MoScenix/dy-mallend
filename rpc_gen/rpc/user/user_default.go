package user

import (
	"context"
	user "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetAuth(ctx context.Context, req *user.GetAuthReq, callOptions ...callopt.Option) (resp *user.GetAuthResp, err error) {
	resp, err = defaultClient.GetAuth(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetAuth call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AuthChange(ctx context.Context, req *user.AuthChangeReq, callOptions ...callopt.Option) (resp *user.AuthChangeResp, err error) {
	resp, err = defaultClient.AuthChange(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AuthChange call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

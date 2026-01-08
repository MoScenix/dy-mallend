package main

import (
	"context"
	cart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/service"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// SetItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) SetItem(ctx context.Context, req *cart.SetItemReq) (resp *cart.SetItemResp, err error) {
	resp, err = service.NewSetItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

package order

import (
	"context"
	"fmt"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/service"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	common "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/common"
	order "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// OrderAuthChange .
// @router /order [POST]
func OrderAuthChange(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderAuthReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewOrderAuthChangeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// OrderAuth .
// @router /order [GET]
func OrderAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderAuthService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "center", utils.WarpResponse(ctx, c, resp))
}

// OrderProductPut .
// @router /order/product [POST]
func OrderProductPut(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderProductPutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderProductPutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// OrderProdectDelete .
// @router /order/product/delete [GET]
func OrderProdectDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.DeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewOrderProdectDeleteService(ctx, c).Run(&req)
	fmt.Println(err)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Println(resp)
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

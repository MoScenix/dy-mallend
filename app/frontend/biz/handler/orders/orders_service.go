package orders

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/service"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	common "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/common"
	orders "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/orders"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetOrders .
// @router /orders [GET]
func GetOrders(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewGetOrdersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "orders", utils.WarpResponse(ctx, c, resp))
}

// AddOrder .
// @router /orders [POST]
func AddOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req orders.Order
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewAddOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

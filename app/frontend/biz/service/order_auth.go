package service

import (
	"context"
	"errors"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	common "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/common"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderAuthService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderAuthService(Context context.Context, RequestContext *app.RequestContext) *OrderAuthService {
	return &OrderAuthService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderAuthService) Run(req *common.Empty) (resp map[string]any, err error) {

	userIDAny := h.Context.Value(utils.UserIdKey)

	resp = make(map[string]any)
	if userIDAny == nil {
		return nil, errors.New("user_id is nil")
	}
	uid32 := int32(userIDAny.(float64))
	userResp, err := rpc.UserClient.GetAuth(h.Context, &user.GetAuthReq{
		UserId: uid32,
	})
	if err != nil {
		return nil, err
	}
	productsResp, err := rpc.ProductClient.GetProDucts(h.Context, &product.GetProDuctsReq{
		UserId: uint32(uid32),
	})
	if err != nil {
		return nil, err
	}
	// userIdStr := strconv.Itoa(int(uid32))
	resp["user_name"] = userResp.UserName
	resp["description"] = userResp.Description

	items := make([]map[string]any, 0, len(productsResp.Products))
	for _, p := range productsResp.Products {
		items = append(items, map[string]any{
			"Id":          p.Id,
			"ProductName": p.Name,
			"Picture":     p.Picture,
			"Cost":        p.Price,
		})
	}
	resp["orders"] = []any{
		map[string]any{
			"Items": items,
		},
	}

	return resp, nil
}

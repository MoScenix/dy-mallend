package service

import (
	"context"
	"io"
	"os"
	"strconv"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	order "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/order"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type OrderAuthChangeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderAuthChangeService(Context context.Context, RequestContext *app.RequestContext) *OrderAuthChangeService {
	return &OrderAuthChangeService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderAuthChangeService) Run(req *order.OrderAuthReq) (resp map[string]any, err error) {
	user_id := uint32(h.Context.Value(utils.UserIdKey).(float64))
	avatar, err := h.RequestContext.FormFile("avatar")
	if err == nil && avatar != nil {
		avatar_file, err := avatar.Open()
		defer avatar_file.Close()

		avatarlocal, err := os.Create("./static/image/avatar/" + strconv.Itoa(int(user_id)) + ".jpg")
		if err != nil {
			hlog.Error("create avatar file error")
			return map[string]any{
				"success": false,
			}, err
		} else {
			_, err := io.Copy(avatarlocal, avatar_file)
			if err != nil {
				return map[string]any{
					"success": false,
				}, err
			}
		}
	}
	var change user.AuthChangeReq
	change.UserId = user_id
	change.UserName = req.UserName
	change.Description = req.Description
	res, err := rpc.UserClient.AuthChange(h.Context, &change)
	if !res.Success && err != nil {
		return map[string]any{
			"success": false,
		}, err
	}
	return map[string]any{
		"success": true,
	}, nil
}

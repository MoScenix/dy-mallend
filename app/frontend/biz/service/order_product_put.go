package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/frontend/biz/utils"
	order "github.com/MoScenix/douyin-mall-backend/app/frontend/hertz_gen/frontend/order"
	"github.com/MoScenix/douyin-mall-backend/app/frontend/infra/rpc"
	"github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
)

type OrderProductPutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderProductPutService(Context context.Context, RequestContext *app.RequestContext) *OrderProductPutService {
	return &OrderProductPutService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderProductPutService) Run(req *order.OrderProductPutReq) (resp map[string]any, err error) {
	picture, err := h.RequestContext.FormFile("picture")

	if err != nil || picture == nil || req.Name == "" || req.Price <= 0 {
		return map[string]any{"success": false}, kerrors.NewBizStatusError(2004003, "product name and price is required")
	}
	picturename := "./static/image/product/" + uuid.New().String() + ".jpg"
	err = h.RequestContext.SaveUploadedFile(picture, picturename)
	if err != nil {
		return map[string]any{"success": false}, kerrors.NewBizStatusError(2004004, "picture upload failed")
	}
	_, err = rpc.ProductClient.AddProduct(h.Context, &product.AddProductReq{
		UserId:      uint32(h.Context.Value(utils.UserIdKey).(float64)),
		Name:        req.Name,
		Price:       req.Price,
		Picture:     picturename[1:],
		Description: req.Description,
	})
	if err != nil {
		return map[string]any{"success": false}, err
	}
	return map[string]any{"success": true}, nil
}

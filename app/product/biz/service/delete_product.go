package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/redis"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/model"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// Finish your business logic.
	var ProductQuery = model.NewProductProQuery(s.ctx, mysql.DB, redis.RedisClient)
	DeleteProudct, err := ProductQuery.GetById(int(req.Id))
	if err != nil {
		return &product.DeleteProductResp{
			Success: false,
		}, kerrors.NewBizStatusError(2004001, "product id is required")
	}
	if req.UserId != uint32(DeleteProudct.UserID) {
		return &product.DeleteProductResp{
			Success: false,
		}, kerrors.NewBizStatusError(2004001, "you are not the owner of this product")
	}
	err = ProductQuery.DeleteProduct(int(req.Id))
	return &product.DeleteProductResp{
		Success: true,
		Picture: DeleteProudct.Picture,
	}, err
}

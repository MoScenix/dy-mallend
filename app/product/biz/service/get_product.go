package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/redis"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/model"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id <= 0 {
		return nil, kerrors.NewBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductProQuery(s.ctx, mysql.DB, redis.RedisClient)
	row, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(row.ID),
			Name:        row.Name,
			Price:       float32(row.Price),
			Picture:     row.Picture,
			Description: row.Description,
		},
	}, err
}

package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/model"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
)

type GetProductsByIdService struct {
	ctx context.Context
} // NewGetProductsByIdService new GetProductsByIdService
func NewGetProductsByIdService(ctx context.Context) *GetProductsByIdService {
	return &GetProductsByIdService{ctx: ctx}
}

// Run create note info
func (s *GetProductsByIdService) Run(req *product.GetProductsByIdReq) (resp *product.GetProductsByIdResp, err error) {
	// Finish your business logic.
	row, err := model.NewProductQuery(s.ctx, mysql.DB).GetByIDsWithDeleted(req.Ids)
	res := product.GetProductsByIdResp{
		Products: []*product.GetProductByIdRespIteam{},
	}
	for _, item := range row {
		var isDeleted bool
		if item.DeletedAt.Valid {
			isDeleted = true
		}
		res.Products = append(res.Products, &product.GetProductByIdRespIteam{
			Name:        item.Name,
			Description: item.Description,
			Price:       float32(item.Price),
			Picture:     item.Picture,
			Isdeleted:   isDeleted,
			Pid:         uint32(item.ID),
		})
	}
	return &res, err
}

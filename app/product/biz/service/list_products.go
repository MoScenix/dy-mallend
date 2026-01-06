package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/model"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categioryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	row, err := categioryQuery.GetProductByCategory(req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{
		Products: []*product.Product{},
	}
	for _, p := range row {
		for _, p2 := range p.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p2.ID),
				Name:        p2.Name,
				Description: p2.Description,
				Price:       float32(p2.Price),
				Picture:     p2.Picture,
			})
		}
	}
	return resp, err
}

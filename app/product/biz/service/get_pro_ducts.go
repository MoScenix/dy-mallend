package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/model"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
)

type GetProDuctsService struct {
	ctx context.Context
} // NewGetProDuctsService new GetProDuctsService
func NewGetProDuctsService(ctx context.Context) *GetProDuctsService {
	return &GetProDuctsService{ctx: ctx}
}

// Run create note info
func (s *GetProDuctsService) Run(req *product.GetProDuctsReq) (resp *product.GetProDuctsResp, err error) {
	// Finish your business logic.

	resp = &product.GetProDuctsResp{
		Products: []*product.Product{},
	}
	var Products []model.Product

	Products, err = model.NewProductQuery(s.ctx, mysql.DB).GetProductsById(int(req.UserId))
	for _, p := range Products {
		resp.Products = append(resp.Products, &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       float32(p.Price),
		})
	}
	return resp, err
}

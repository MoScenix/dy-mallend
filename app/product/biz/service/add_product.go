package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/product/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/product/biz/model"
	product "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	// Finish your business logic.
	var NewProduct model.Product
	NewProduct = model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		UserID:      int(req.UserId),
	}
	err = model.NewProductQuery(s.ctx, mysql.DB).AddProduct(NewProduct)
	return &product.AddProductResp{
		Id: uint32(NewProduct.ID),
	}, err
}

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/cart/biz/model"
	cart "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/cart"
	"gorm.io/gorm"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	fmt.Println(111)
	row, err := model.NewCartQuery(s.ctx, mysql.DB).GetByUser(req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		} else {
			return nil, err
		}
	}
	res := &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  make([]*cart.CartItem, 0),
		},
	}

	for _, item := range row.Items {
		res.Cart.Items = append(res.Cart.Items, &cart.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return res, nil
}

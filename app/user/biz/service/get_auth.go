package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/user/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/user/biz/model"
	user "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
)

type GetAuthService struct {
	ctx context.Context
} // NewGetAuthService new GetAuthService
func NewGetAuthService(ctx context.Context) *GetAuthService {
	return &GetAuthService{ctx: ctx}
}

// Run create note info
func (s *GetAuthService) Run(req *user.GetAuthReq) (resp *user.GetAuthResp, err error) {
	// Finish your business logic.
	User, err := model.GetByID(mysql.DB, int(req.UserId))
	resp = &user.GetAuthResp{
		UserName:    User.UserName,
		Description: User.Description,
	}
	return resp, err
}

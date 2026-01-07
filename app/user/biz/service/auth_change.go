package service

import (
	"context"

	"github.com/MoScenix/douyin-mall-backend/app/user/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/user/biz/model"
	user "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
)

type AuthChangeService struct {
	ctx context.Context
} // NewAuthChangeService new AuthChangeService
func NewAuthChangeService(ctx context.Context) *AuthChangeService {
	return &AuthChangeService{ctx: ctx}
}

// Run create note info
func (s *AuthChangeService) Run(req *user.AuthChangeReq) (resp *user.AuthChangeResp, err error) {
	// Finish your business logic.
	var u model.User
	u.ID = uint(req.UserId)
	u.Description = req.Description
	u.UserName = req.UserName
	err = model.UpdateUser(mysql.DB, u)
	if err != nil {
		return &user.AuthChangeResp{
			Success: true,
		}, err
	} else {
		return &user.AuthChangeResp{
			Success: false,
		}, nil
	}
}

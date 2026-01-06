package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/MoScenix/douyin-mall-backend/app/user/biz/dal/mysql"
	"github.com/MoScenix/douyin-mall-backend/app/user/biz/model"
	user "github.com/MoScenix/douyin-mall-backend/rpc_gen/kitex_gen/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.ConfirmPassword != req.Password {
		return nil, errors.New("password not equal confirm password")
	}
	if req.Email == "" || req.Password == "" {
		fmt.Println(req.Email)
		fmt.Println(req.Password)
		return nil, errors.New("email or password is empty")
	}
	PasswordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		UserName:       uuid.New().String(),
		PasswordHashed: string(PasswordHashed),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{
		UserId:  int32(newUser.ID),
		UerName: newUser.UserName,
	}, nil
}

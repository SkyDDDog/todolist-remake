package service

import (
	"context"

	"todolist-remake/cmd/user/dal/db"
	"todolist-remake/cmd/user/model"
	"todolist-remake/idl/pb/user"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (us *UserService) Login(req *user.LoginRequest) (resp *model.User, err error) {
	err = db.CheckUserPassword(us.ctx, req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	return db.GetUserByUsername(us.ctx, req.Username)
}

func (us *UserService) Register(req *user.RegisterRequest) error {
	return db.CreateUser(us.ctx, req)
}

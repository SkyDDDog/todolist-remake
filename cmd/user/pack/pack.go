package pack

import (
	"errors"
	"todolist-remake/cmd/user/model"
	"todolist-remake/idl/pb/user"
	"todolist-remake/pkg/errno"
)

func BuildBaseResp(err error) *user.Base {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func BuildUser(source *model.User) *user.User {
	return &user.User{
		Id:       source.ID,
		Username: source.Username,
	}
}

func baseResp(err errno.ErrNo) *user.Base {
	return &user.Base{
		Code:    err.ErrorCode,
		Message: err.ErrorMsg,
	}
}

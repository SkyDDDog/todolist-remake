package handler

import (
	"github.com/gin-gonic/gin"
	"todolist-remake/cmd/api/rpc"
	"todolist-remake/cmd/api/types"
	"todolist-remake/config"
	"todolist-remake/idl/pb/user"
	"todolist-remake/pkg/errno"
	"todolist-remake/pkg/utils"
)

func UserRegister(c *gin.Context) {
	var req types.UserRegisterRequest

	err := c.Bind(&req)

	if err != nil {
		BuildFailResponse(c, errno.ParamError)
		return
	}

	err = rpc.UserRegister(c, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		BuildFailResponse(c, err)
		return
	}

	BuildSuccessResponse(c, nil)
}

func UserLogin(c *gin.Context) {
	var req types.UserLoginRequest

	err := c.Bind(&req)

	if err != nil {
		BuildFailResponse(c, errno.ParamError)
		return
	}

	user, err := rpc.UserLogin(c, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		BuildFailResponse(c, err)
		return
	}

	token, err := utils.GenerateToken(user.Id, config.Server.Secret)

	if err != nil {
		BuildFailResponse(c, errno.AuthorizationFailError.WithMessage(err.Error()))
		return
	}

	BuildSuccessResponse(c, &types.UserLoginResponse{
		User: types.User{
			ID:       user.Id,
			Username: user.Username,
		},
		Token: token,
	})
}

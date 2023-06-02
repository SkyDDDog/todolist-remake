package middleware

import (
	"github.com/gin-gonic/gin"
	"todolist-remake/cmd/api/handler"
	"todolist-remake/config"
	"todolist-remake/consts"
	"todolist-remake/pkg/errno"
	"todolist-remake/pkg/utils"
)

func JWT(c *gin.Context) {
	var claims *utils.Claims
	var err error

	token := c.GetHeader(consts.AuthHeader)
	if token == "" {
		handler.BuildFailResponse(c, errno.AuthorizationFailError)
		c.Abort()
		return
	}

	claims, err = utils.ParseToken(token, config.Server.Secret)
	if err != nil {
		handler.BuildFailResponse(c, err)
		c.Abort()
		return
	}

	token, err = utils.GenerateToken(claims.UserID, config.Server.Secret)

	if err != nil {
		handler.BuildFailResponse(c, errno.NewErrNo(errno.AuthorizationFailedErrCode, err.Error()))
		c.Abort()
		return
	}
	c.Header(consts.AuthHeader, token)
	c.Set(consts.KeyUserId, claims.UserID)

	c.Next()
}

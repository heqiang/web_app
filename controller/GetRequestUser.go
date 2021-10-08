package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxtUserIDKEY = "user_id"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxtUserIDKEY)
	if !ok {
		err = ErrorUserNotLogin

		return
	}
	userId, ok := uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return userId, nil
}

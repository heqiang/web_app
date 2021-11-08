package middlerware

import (
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/utils"
	jwt2 "web_app/utils/jwt"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("token")
		if authHeader == "" {
			utils.ResponseError(c, utils.CodeNeedAuth)
			c.Abort()
			return
		}
		mc, err := jwt2.ParseToken(authHeader)
		if err != nil {
			utils.ResponseError(c, utils.CodeInvaildAuth)
			c.Abort()
			return
		}
		// 将当前请求的UserId信息保存到请求的上下文c上
		c.Set(controller.CtxtUserIDKEY, mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}
}

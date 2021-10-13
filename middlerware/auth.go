package middlerware

import (
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/pkg/jwt"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("token")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedAuth)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			controller.ResponseError(c, controller.CodeInvaildAuth)
			c.Abort()
			return
		}
		// 将当前请求的UserId信息保存到请求的上下文c上
		c.Set(controller.CtxtUserIDKEY, mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}
}

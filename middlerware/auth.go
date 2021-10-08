package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/pkg/jwt"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		//authHeader := c.Request.Header.Get("Authorization")
		authHeader := c.Request.Header.Get("token")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedAuth)
			c.Abort()
			return
		}
		// 按空格分割
		//parts := strings.SplitN(authHeader, " ", 2)
		//if !(len(parts) == 2 && parts[0] == "Bearer") {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 2004,
		//		"msg":  "请求头中auth格式有误",
		//	})
		//	c.Abort()
		//	return
		//}
		//// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(authHeader)
		fmt.Println(mc)
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
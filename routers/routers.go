package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller/community"
	"web_app/controller/posts"
	"web_app/controller/user"
	"web_app/logger"
	"web_app/settings"
)

func Setup(conf *settings.AppConfig) *gin.Engine {
	gin.SetMode(conf.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//v1 := r.Group("/api/v1",middlerware.JWTAuthMiddleware())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {})
		// 登录注册
		v1.POST("/register", user.RegisterHandle)
		v1.POST("/Login", user.LoginHadle)
		// 社区列表
		v1.GET("/communityDetail/:id", community.CommunityDetailHandle)
		// 获取社区列表
		v1.GET("/communityList", community.CommunityListHandle)
		v1.POST("/post", posts.PostCommunityHandle)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

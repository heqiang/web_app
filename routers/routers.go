package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller/community"
	"web_app/controller/posts"
	"web_app/controller/user"
	"web_app/logger"
	"web_app/middlerware"
	"web_app/settings"
)

func Setup(conf *settings.AppConfig) *gin.Engine {
	gin.SetMode(conf.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	v1.POST("/register", user.RegisterHandle)
	v1.POST("/Login", user.LoginHadle)
	v1.Use(middlerware.JWTAuthMiddleware())
	{
		// 社区列表
		v1.GET("/communityDetail/:id", community.CommunityDetailHandle)
		// 获取社区列表
		v1.GET("/communityList", community.CommunityListHandle)
		v1.POST("/post", posts.PostCommunityHandle)
		v1.GET("/post/:postId", posts.GetPostDeatilHandle)
		v1.POST("/GetPostList", posts.GetPostListHandle)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "网页消失了",
		})
	})
	return r
}

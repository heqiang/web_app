package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/controller"

	"web_app/dao/mysqlc/model"
	"web_app/logic"
	"web_app/pkg/snowflake"
)

// PostCommunityHandle 创建帖子
func PostCommunityHandle(c *gin.Context) {
	//获取帖子的参数然后进行校验
	var post model.Post
	if err := c.ShouldBind(&post); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			controller.ResponseError(c, controller.CodeInvaildParam)
			return
		}
		zap.L().Error("参数有误", zap.Error(errs))
		controller.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), controller.CodeInvaildParam)
	}
	id, _ := controller.GetCurrentUser(c)
	post.Post_id = snowflake.GetSnowId()
	post.Author_id = id
	err := logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		controller.ResponseSuccess(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, controller.CodeSuccess)

	//返回响应

}

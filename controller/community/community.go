package community

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/controller"
	"web_app/dao/mysqlc/model"
	"web_app/logic"
)

// --------社区相关handle

func CommunityHandle(c *gin.Context) {
	//查询到所有的社区(community_id,community_name)
	var comm model.Community
	if err := c.ShouldBind(comm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			controller.ResponseError(c, controller.CodeInvaildParam)
			return
		}
		zap.L().Error("社区获取 参数异常", zap.Error(err))
		controller.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), controller.CodeInvaildParam)
		return
	}
	// 业务处理
	//data, err := logic.GetCommunityList(comm)
	//if err != nil {
	//	zap.L().Error("logic.GetCommunityList Error", zap.Error(err))
	//	//不轻易将服务端报错是暴露到外面
	//	controller.ResponseError(c, controller.CodeServerBusy)
	//	return
	//}
	//controller.ResponseSuccess(c, data)
}

func CommunityListHandle(c *gin.Context) {
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Warn("logic.GetCommunityList() failed ", zap.Error(err))
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, communityList)
}

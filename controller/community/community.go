package community

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/controller"
	"web_app/logic"
)

// CommunityDetailHandle --------社区详情
func CommunityDetailHandle(c *gin.Context) {
	idStr := c.Param("id")
	communityId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		controller.ResponseError(c, controller.CodeInvaildParam)
		return
	}
	communityDtail, err := logic.GetCommunityDetail(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		controller.ResponseErrorWithMsg(c, err.Error(), controller.CodeInvaildParam)
		return
	}
	controller.ResponseSuccess(c, communityDtail)
}

// CommunityListHandle 社区列表
func CommunityListHandle(c *gin.Context) {
	userId, _ := c.Get("user_id")
	fmt.Println("用户id：", userId)
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Warn("logic.GetCommunityList() failed ", zap.Error(err))
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, communityList)
}

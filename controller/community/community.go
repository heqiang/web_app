package community

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/service/implement"
	"web_app/utils"
)

// CommunityDetailHandle --------社区详情
func CommunityDetailHandle(c *gin.Context) {
	idStr := c.Param("id")
	communityId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ResponseError(c, utils.CodeInvaildParam)
		return
	}
	var comm implement.Community
	communityDtail, err := comm.GetCommunityDetail(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		utils.ResponseErrorWithMsg(c, err.Error(), utils.CodeInvaildParam)
		return
	}
	utils.ResponseSuccess(c, communityDtail)
}

// CommunityListHandle 社区列表
func CommunityListHandle(c *gin.Context) {
	userId, _ := c.Get("user_id")
	fmt.Println("用户id：", userId)
	var comm implement.Community
	communityList, err := comm.GetCommunityList()
	if err != nil {
		zap.L().Warn("logic.GetCommunityList() failed ", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, communityList)
}

package posts

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web_app/controller"
	"web_app/controller/posts/postmodel"
	"web_app/dao/mysqlc"

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
		return
	}
	id, _ := controller.GetCurrentUser(c)
	post.Post_id = snowflake.GetSnowId()
	post.AuthorId = id
	err := logic.CreatePost(&post)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		controller.ResponseSuccess(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, controller.CodeSuccess)

}

// GetPostDeatilHadle 获取帖子的详情
func GetPostDeatilHandle(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("postId"), 10, 64)
	if err != nil {
		zap.L().Error("get postId failed ", zap.Error(err))
		controller.ResponseError(c, controller.CodeInvaildParam)
		return
	}
	postDetail, err := logic.GetPostDetail(postId)
	if err != nil {
		zap.L().Error("logic.GetPostDetail failed ", zap.Error(err))
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	user := mysqlc.QueryByUserId(postDetail.AuthorId)
	community := mysqlc.QueryByCommId(postDetail.CommunityId)

	postApiDetail := &postmodel.ApiPostDetail{
		AuthorName: user.UserName,
	}
	postApiDetail.Community = community
	postApiDetail.Post = postDetail
	controller.ResponseSuccess(c, postApiDetail)
}

func GetPostListHandle(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	size, _ := strconv.Atoi(c.Query("size"))
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	offset := (page - 1) * size
	allPosts, total, err := logic.GetPostList(offset, size)
	if err != nil {
		zap.L().Error(" logic.GetPostList failed ", zap.Error(err))
		return
	}
	paginationQ := &postmodel.PaginationQ{
		Ok:    true,
		Size:  uint(size),
		Page:  uint(page),
		Total: total,
		Data:  allPosts,
	}
	controller.ResponseSuccess(c, paginationQ)
}
func PostVotedHandle(c *gin.Context) {
	voted := new(postmodel.VoteData)
	if err := c.ShouldBind(voted); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("请求参数错误", zap.Error(err))
			controller.ResponseError(c, controller.CodeInvaildParam)
			return
		}
		zap.L().Error("请求参数验证错误", zap.Error(err))
		controller.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), controller.CodeInvaildParam)
		return
	}
	logic.PostVote(voted)
	controller.ResponseSuccess(c, nil)
}

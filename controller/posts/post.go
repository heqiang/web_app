package posts

import (
	"fmt"
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

const (
	orderTime  = "time"
	orderScore = "score"
)

// PostCommunityHandle
// @Tags 帖子相关接口
// @Security ApiKeyAuth
// @Description 用户发帖
// @Summary 用户发帖
// @title 用户发帖
// @Security
// @Param data body model.Post true "请示参数data"
// @Success 200 object controller.ResponseData "请求成功"
// @Router /api/v1/post  [post]
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

// GetPostDeatilHandle
// @Tags 帖子相关接口
// @Description 帖子详情
// @Summary 帖子接口
// @title 用户发帖
// @Security
// @Security ApiKeyAuth
// @Param postId path int true "postId"
// @Success 200 object controller.ResponseData "请求成功"
// @Router /api/v1/post/{postId}  [get]
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

// GetPostListHandle
// @Tags 帖子相关接口
// @Description 获取所有的帖子
// @Summary 获取所有的帖子
// @title 获取所有的帖子
// @Security ApiKeyAuth
// @Param page path string false "页数"
// @Param size path string false "size"
// @Success 200 object controller.ResponseData "请求成功"
// @Router /api/v1/GetPostList  [get]
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

// 帖子的列表按照时间还是分数
// GetPostListHandler2 /api/v1/post2?page=1&size=10&oreder=time
// GetPostListHandler2 /api/v1/post2?page=1&size=10&oreder=score
func GetPostListHandler2(c *gin.Context) {
	paramlist := postmodel.ParamPostList{
		Page:  1,
		Size:  10,
		Order: orderTime,
	}
	err := c.ShouldBindQuery(&paramlist)
	if err != nil {
		zap.L().Error("GetPostListHandle2 with invaild param", zap.Error(err))
		controller.ResponseError(c, controller.CodeInvaildParam)
		return
	}
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
	category := c.Query("order")
	fmt.Println(category)
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

// PostVotedHandle
// @Tags 帖子相关接口
// @Description 帖子投票
// @Summary 帖子投票
// @title 帖子投票
// @Security ApiKeyAuth
// @Param voteData body  postmodel.VoteData true "投票参数data"
// @Success 200 object controller.ResponseData "请求成功"
// @Router /api/v1/vote  [post]
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
	userId, err := controller.GetCurrentUser(c)
	if err != nil {
		controller.ResponseError(c, controller.CodeNeedAuth)
		return
	}
	err = logic.PostVote(voted, userId)
	if err != nil {
		zap.L().Error("votedForpost", zap.Error(err))
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, nil)
}

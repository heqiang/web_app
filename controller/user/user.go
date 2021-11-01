package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/controller"
	"web_app/dao/mysqlc/model"
	"web_app/logic"
)

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Accept application/json
// @Produce application/json
// @termsOfService http://swagger.io/terms/
// @Param object query model.User false "查询参数"
// @Router /posts2 [get]
// @host 127.0.0.1
// @Router /register [post]

func RegisterHandle(c *gin.Context) {
	var p model.User
	// 1 参数校验
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非vaildator类型的错误 直接返回
			controller.ResponseError(c, controller.CodeInvaildParam)
			return
		}
		//请求失败
		zap.L().Error("用户注册 invaild param", zap.Error(err))
		controller.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), controller.CodeInvaildParam)
		return
	}
	// 2业务处理
	err := logic.UserRegister(&p)
	if err != nil {
		controller.ResponseError(c, controller.CodeUserExist)
		return
	}
	controller.ResponseSuccess(c, nil)
}

// LoginHadle 登录
func LoginHadle(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			controller.ResponseError(c, controller.CodeInvaildParam)
			return
		}
		zap.L().Error("登录失败", zap.Error(err))
		controller.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), controller.CodeInvaildParam)
		return
	}
	token, err := logic.UserLogin(&u)
	if err != nil {
		controller.ResponseError(c, controller.CodeInvaildPasswordorUserName)
		return
	}
	//获取用户信息
	controller.ResponseSuccess(c, token)
}

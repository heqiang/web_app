package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/controller"
	"web_app/dao/mysqlc/model"
	"web_app/service/implement"
	"web_app/utils"
)

// RegisterHandle 注册
// @Tags 用户相关接口
// @Summary 用户注册接口
// @title 用户注册
// @Param data body swagtype.UserRegiter true "请示参数data"
// @Success 200 object controller.ResponseData "请求成功"
// @Router /api/v1/register  [post]
func RegisterHandle(c *gin.Context) {
	var p model.User
	// 1 参数校验
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非vaildator类型的错误 直接返回
			utils.ResponseError(c, utils.CodeInvaildParam)
			return
		}
		//请求失败
		zap.L().Error("用户注册 invaild param", zap.Error(err))
		utils.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), utils.CodeInvaildParam)
		return
	}
	// 2业务处理
	var user implement.User
	err := user.Register(&p)
	if err != nil {
		utils.ResponseError(c, utils.CodeUserExist)
		return
	}
	utils.ResponseSuccess(c, nil)
}

// LoginHadle 登录
// @Tags 用户相关接口
// @Summary 用户登录接口
// @title 用户登录
// @Param logindata body swagtype.UserLogin true "请示参数data"
// @Success 200 object controller.ResponseData "请求成功"
// @Host 127.0.0.1
// @Router /api/v1/Login  [post]
func LoginHadle(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			utils.ResponseError(c, utils.CodeInvaildParam)
			return
		}
		zap.L().Error("登录失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), utils.CodeInvaildParam)
		return
	}
	var user implement.User
	token, err := user.Login(&u)
	if err != nil {
		utils.ResponseError(c, utils.CodeInvaildPasswordorUserName)
		return
	}
	//获取用户信息
	utils.ResponseSuccess(c, token)
}

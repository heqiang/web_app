package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strings"
	"web_app/dao/mysqlc/model"
	"web_app/logic"
)

// RegisterHandle用户注册
func RegisterHandle(c *gin.Context) {
	var p model.User
	// 1 参数校验
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非vaildator类型的错误 直接返回
			ResponseError(c, CodeInvaildParam)
			return
		}
		//请求失败
		zap.L().Error("用户注册 invaild param", zap.Error(err))
		ResponseErrorWithMsg(c, removeTopStruct(errs.Translate(trans)), CodeInvaildParam)
		return
	}
	// 2业务处理
	err := logic.UserRegister(&p)
	if err != nil {
		ResponseError(c, CodeUserExist)
		return
	}
	ResponseSuccess(c, nil)
}
func LoginHadle(c *gin.Context) {

	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseError(c, CodeInvaildParam)
			return
		}
		zap.L().Error("登录失败", zap.Error(err))
		ResponseErrorWithMsg(c, removeTopStruct(errs.Translate(trans)), CodeInvaildParam)
		return
	}
	token, err := logic.UserLogin(&u)
	if err != nil {
		ResponseError(c, CodeInvaildPasswordorUserName)
		return
	}
	ResponseSuccess(c, token)
}

// removeTopStruct 去除提示中的结构体名称
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

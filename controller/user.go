package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		//请求失败
		zap.L().Error("用户注册 invaild param", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数有误",
			"err": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	fmt.Println(p)
	// 2业务处理
	err := logic.UserRegister(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

// removeTopStruct 去除提示中的结构体名称
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

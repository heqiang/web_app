package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseData 数据接口相应数据
type ResponseData struct {
	Code ResCode     `json:"code"`           //业务相应状态码
	Msg  interface{} `json:"msg"`            //提示信息
	Data interface{} `json:"data,omitempty"` // 数据
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseErrorWithMsg(c *gin.Context, msg interface{}, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

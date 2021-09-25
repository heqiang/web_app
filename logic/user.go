package logic

import (
	"errors"
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
	"web_app/pkg/snowflake"
)

//用户注册业务逻辑
func UserRegister(userinfo *model.User) (err error) {
	//判断用户是否存在
	err = mysqlc.QueryUserByName(userinfo.UserName)
	if err != nil {
		return errors.New("用户已存在")
	}
	// 生成Uid
	UUID := snowflake.GetSnowId()
	// 构造一个user实例
	u := model.User{
		UserId:   UUID,
		UserName: userinfo.UserName,
		Password: userinfo.Password,
	}
	// mysql入库
	return mysqlc.InsertUser(&u)
}

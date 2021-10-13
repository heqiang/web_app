package logic

import (
	"errors"
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

// UserRegister 用户注册业务逻辑
func UserRegister(userinfo *model.User) (err error) {
	err = mysqlc.QueryUserByName(userinfo.UserName)
	if err != nil {
		return errors.New("用户已存在")
	}
	// 生成Uid
	UUID := snowflake.GetSnowId()
	u := model.User{
		UserId:   UUID,
		UserName: userinfo.UserName,
		Password: userinfo.Password,
	}
	// mysql入库
	return mysqlc.InsertUser(&u)
}

func UserLogin(userinfo *model.User) (token string, err error) {

	err = mysqlc.QueryByUser(userinfo.UserName, userinfo.Password)
	if err != nil {
		return "", err
	}
	token, err1 := jwt.GenToken(userinfo.UserName, userinfo.UserId)
	if err1 != nil {
		return "", nil
	}
	return token, nil

}

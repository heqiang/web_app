package implement

import (
	"errors"
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
	snowflake2 "web_app/utils/snowflake"
)

type User struct {
}

func (user *User) Login(user1 *model.User) (token string, err error) {
	return mysqlc.QueryByUser(user1.UserName, user1.Password)

}
func (user *User) Register(user1 *model.User) (err error) {
	err = mysqlc.QueryUserByName(user1.UserName)
	if err != nil {
		return errors.New("用户已存在")
	}
	// 生成Uid
	UUID := snowflake2.GetSnowId()
	u := model.User{
		UserId:   UUID,
		UserName: user1.UserName,
		Password: user1.Password,
	}
	// mysql入库
	return mysqlc.InsertUser(&u)
}

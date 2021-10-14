package mysqlc

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"web_app/dao/mysqlc/model"
	"web_app/pkg/jwt"
)

const serct = "1422127065@qq.com"

func QueryUserByName(username string) (err error) {
	var user model.User
	db.Where("username=?", username).Take(&user)
	fmt.Println(user.UserName)
	if user.UserName != "" {
		return errors.New("已存在该用户")
	}
	return nil
}

// QueryByUser 用户信息查询
func QueryByUser(username, password string) (token string, err error) {
	var user model.User
	queryRes := db.Where("username=? and password=?", username, encryptPassword(password)).Take(&user)
	if queryRes.RowsAffected == 0 {
		err = errors.New("用户不存在")
		return
	}
	fmt.Println(user)
	return jwt.GenToken(user.UserName, user.UserId)
}

func InsertUser(user *model.User) (err error) {
	user.Password = encryptPassword(user.Password)
	res := db.Create(user)
	if res.RowsAffected >= 1 {
		return nil
	}
	return errors.New("新增失败")
}

// QueryByUserId 获取user信息
func QueryByUserId(id int64) (user *model.User) {
	user = new(model.User)
	db.Where("user_id=?", id).Take(&user)
	return
}

// 用户密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(serct))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

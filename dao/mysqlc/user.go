package mysqlc

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"web_app/dao/mysqlc/model"
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
func QueryByUser(userinfo *model.User) (err error) {
	var user model.User
	db.Where("username=?", userinfo.UserName).Take(&user)
	if user.UserName != "" {
		password := encryptPassword(userinfo.Password)
		if user.Password != password {
			return errors.New("密码错误")
		}
		return nil
	}
	return errors.New("用户不存在")
}

func InsertUser(user *model.User) (err error) {
	user.Password = encryptPassword(user.Password)
	res := db.Create(user)
	if res.RowsAffected >= 1 {
		return nil
	}
	return errors.New("新增失败")
}

// 用户密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(serct))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

package mysqlc

import (
	"errors"
	"web_app/dao/mysqlc/model"
)

func InsertPost(postmode *model.Post) (err error) {
	insertRes := db.Create(&postmode)
	if insertRes.RowsAffected == 0 {
		err = errors.New("帖子创建失败 系统错误")
		return
	}
	return
}

package logic

import (
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
)

func CreatePost(postmode *model.Post) (err error) {
	return mysqlc.InsertPost(postmode)
}

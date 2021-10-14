package logic

import (
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
)

func CreatePost(postmode *model.Post) (err error) {
	return mysqlc.InsertPost(postmode)
}

func GetPostDetail(postId int64) (postDetail *model.Post, err error) {
	return mysqlc.QueryPostDetail(postId)
}

func GetPostList() (postList []*model.Post) {
	return mysqlc.QueryAllPosts()
}

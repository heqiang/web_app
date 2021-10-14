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

// QueryPostDetail 获取某个帖子的详情
func QueryPostDetail(id int64) (postDetail *model.Post, err error) {
	postDetail = new(model.Post)
	queryRes := db.Where("post_id=?", id).Find(&postDetail)
	if queryRes.RowsAffected == 0 {
		err = errors.New("不存在该帖子")
		return
	}
	return

}

//GetPostListHandle 获取所有的贴子
func QueryAllPosts() (postList []*model.Post) {
	var postlist []*model.Post
	db.Find(&postlist)
	return postlist
}

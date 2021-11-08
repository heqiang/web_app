package service

import (
	"web_app/controller/posts/postmodel"
	"web_app/dao/mysqlc/model"
)

type PostInterface interface {
	CreatePost(postmode *model.Post) (err error)
	GetPostDetail(postId int64) (postDetail *model.Post, err error)
	GetPostList(page, size int) (postList []postmodel.ApiPostDetail, total int64, err error)
	PostVote(voted *postmodel.VoteData, userid int64) error
}

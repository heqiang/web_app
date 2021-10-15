package logic

import (
	"web_app/controller/posts/postmodel"
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
)

func CreatePost(postmode *model.Post) (err error) {
	return mysqlc.InsertPost(postmode)
}

func GetPostDetail(postId int64) (postDetail *model.Post, err error) {
	return mysqlc.QueryPostDetail(postId)
}

func GetPostList(page, size int) (postList []postmodel.ApiPostDetail, total int64, err error) {
	var posts []*model.Post
	postList = []postmodel.ApiPostDetail{}
	posts, total, err = mysqlc.QueryAllPosts(page, size)
	for _, post := range posts {
		postDetail, _ := GetPostDetail(post.Post_id)
		user := mysqlc.QueryByUserId(postDetail.AuthorId)
		community := mysqlc.QueryByCommId(postDetail.CommunityId)
		postApiDetail := postmodel.ApiPostDetail{
			AuthorName: user.UserName,
		}
		postApiDetail.Community = community
		postApiDetail.Post = postDetail
		postList = append(postList, postApiDetail)
	}
	return
}

func PostVote(voted *postmodel.VoteData) {

}

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

/*
 帖子的热度 如果在一段时间内没有变化 那么旧帖子就不应该参与排名了
 投一票就加432分  86400/200->200张帖子就可以让帖子续一天
 投票的几种情况
	direction = 1 时,有两种情况：
			1.之前没投过票，现在投赞成票
			2.之前投反对票 现在改投赞成票
*/

// PostVote 帖子投票
func PostVote(voted *postmodel.VoteData, userid int64) {

}

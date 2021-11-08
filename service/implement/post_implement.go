package implement

import (
	"strconv"
	"web_app/controller/posts/postmodel"
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
	"web_app/dao/redis"
)

type Post struct{}

func (p *Post) CreatePost(postmode *model.Post) (err error) {
	err = mysqlc.InsertPost(postmode)
	if err != nil {
		return err
	}
	err = redis.CreatePostTime(postmode.Post_id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) GetPostDetail(postId int64) (postDetail *model.Post, err error) {
	return mysqlc.QueryPostDetail(postId)
}

func (p *Post) GetPostList(page, size int) (postList []postmodel.ApiPostDetail, total int64, err error) {
	var posts []*model.Post
	postList = []postmodel.ApiPostDetail{}
	posts, total, err = mysqlc.QueryAllPosts(page, size)
	for _, post := range posts {
		postDetail, _ := p.GetPostDetail(post.Post_id)
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
			1.之前没有投过票，现在投赞成票   差值的绝对值为 1  1-0  +432
			2.之前投反对票 现在改投赞成票    差值的绝对值为2  1-(-1)  +432*2
	direction = 0 时,有两种情况：
			1.之前投过赞成票，现在要取消   差值的绝对值为1   |0-1| -432
			2.之前投反对票 现在要取消     差值的绝对值为1	 |0-(1)| +432
	direction = -1 时,有两种情况：
			1.之前没投过票，现在投反对票   差值的绝对值1	|-1-0| -432
			2.之前投过赞成票 现在改投反对票  差值的绝对值2  |-1-1|  -432*2
投票的限制：
	帖子发表之日一个星期呢允许用户投票，超过一个星期就不再允许投票
		1、到期之后将redis中保存的赞成票数及反对票数储存到mysql表中
		2 到期之后删除那个 KeyPostVoted
*/

// PostVote 帖子投票
func (p *Post) PostVote(voted *postmodel.VoteData, userid int64) error {
	return redis.VoteForPost(strconv.Itoa(int(userid)), strconv.Itoa(int(voted.PostId)), float64(voted.Direection))

}

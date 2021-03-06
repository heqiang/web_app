package model

import "gorm.io/gorm"

// User 用户注册的参数
type User struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id" `
	UserName string `gorm:"column:username" binding:"required"` //用户姓名
	Password string `gorm:"column:password" binding:"required"` //用户密码
	Email    string `gorm:"column:email"`
}

type Community struct {
	gorm.Model
	CommunityId   int64  `gorm:"column:communityid" binding:"required" `
	CommunityName string `gorm:"column:communityname" binding:"required" `
	Introducion   string `gorm:"column:introducion" `
}

type Post struct {
	gorm.Model
	Post_id     int64  `gorm:"column:post_id" json:"post_id"`
	AuthorId    int64  `gorm:"column:author_id"`
	Status      int    `gorm:"column:status"`
	Title       string `binding:"required" json:"title"`
	Content     string `binding:"required" json:"content"`
	CommunityId int64  `binding:"required" json:"communityid"`
}

// VotedParam  投票
type Voted struct {
	gorm.Model
	// UserId  从token中获取
	PostId int64 `json:"post_id,string"`
	// 赞同 1或者反对 -1
	Direction int `json:"direection,string"`
}

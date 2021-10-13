package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id" `
	UserName string `gorm:"column:username" binding:"required"`
	Password string `gorm:"column:password" binding:"required"`
	Email    string `gorm:"column:email"`
}

type Community struct {
	gorm.Model
	CommunityId   int64  `gorm:"column:communityid" binding:"required" `
	CommunityName string `gorm:"column:communityname" binding:"required" `
	Introducion   string `gorm:"column:introducion" `
}

type Communitydetail struct {
	gorm.Model
	Name         string `gorm:"column:name"`
	Introduction string `gorm:"column:introduction"`
}

type Post struct {
	gorm.Model
	Post_id      int64  `gorm:"column:post_id" `
	Author_id    int64  `gorm:"column:author_id"  binding:"required"`
	Status       int    `gorm:"column:status"`
	Title        string `gorm:"column:title"  binding:"required"`
	Content      string `gorm:"column:content"  binding:"required"`
	Community_id int64  `gorm:"column:communityid"  binding:"required"`
}

package postmodel

import (
	"web_app/dao/mysqlc/model"
)

type Post struct {
	Title   string `gorm:"column:title" binding:"required"`
	Content string `gorm:"column:content" binding:"required"`
}

func (Post) TableName() string {
	return "post"
}

type PostList struct {
}
type ApiPostDetail struct {
	AuthorName       string
	*model.Community `json:"community_detail"`
	*model.Post      `json:"post"`
}

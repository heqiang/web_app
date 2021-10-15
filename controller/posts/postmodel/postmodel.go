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
	*model.Community `json:"community"`
	*model.Post      `json:"post"`
}

type PaginationQ struct {
	Ok    bool        `json:"ok"`
	Page  uint        `form:"page" json:"page"`
	Size  uint        `form:"size" json:"size"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list

}

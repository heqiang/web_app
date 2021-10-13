package postmodel

type Post struct {
	Post_id      int64  `gorm:"column:post_id" `
	Author_id    int64  `gorm:"column:author_id"  binding:"required" `
	Community_id int64  `gorm:"column:communityid" binding:"required" `
	Status       int    `gorm:"column:status" `
	Title        string `gorm:"column:title" binding:"required"`
	Content      string `gorm:"column:content" binding:"required"`
}

func (Post) TableName() string {
	return "post"
}

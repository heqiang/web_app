package communitymodel

type Community struct {
	CommunityId   int64  `gorm:"column:communityid" binding:"required" `
	CommunityName string `gorm:"column:communityname" binding:"required" `
}

func (Community) TableName() string { //实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为users（结构体+s）
	return "community"
}

type CommunityDetail struct {
	CommunityId   int64  `gorm:"column:communityid"`
	CommunityName string `gorm:"column:communityname"`
	Introduction  string `gorm:"column:introduction"`
}

func (CommunityDetail) TableName() string { //实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为users（结构体+s）
	return "communitydetail"
}

package swagtype

type UserLogin struct {
	Username string `gorm:"column:username" binding:"required" json:"username"` //用户姓名
	Password string `gorm:"column:password" binding:"required" json:"password"` //用户密码
}

type UserRegiter struct {
	Username string `gorm:"column:username" binding:"required" json:"username"` //用户姓名
	Password string `gorm:"column:password" binding:"required" json:"password"` //用户密码
	Email    string `gorm:"column:email"`                                       //邮箱
}

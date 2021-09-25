package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id" `
	UserName string `gorm:"column:username" binding:"required"`
	Password string `gorm:"column:password" binding:"required"`
	Email    string `gorm:"column:email"`
}

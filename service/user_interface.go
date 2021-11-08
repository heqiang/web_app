package service

import "web_app/dao/mysqlc/model"

type UserInterface interface {
	Login(user *model.User) (token string, err error)
	Register(user *model.User) (err error)
}

package mysqlc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"web_app/dao/mysqlc/model"
	"web_app/settings"
)

var db *gorm.DB

func InitMySQL(conf *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("mysql  conn err:", err)
		return err
	}
	var user model.User
	var community model.Community
	var post model.Post
	err1 := db.AutoMigrate(user, community, post)
	if err1 != nil {
		return err1
	}
	return nil
}

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
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.DbName)
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
	var communitydetail model.Communitydetail
	//var post model.Post
	err1 := db.AutoMigrate(user, community, communitydetail)
	if err1 != nil {
		return err1
	}
	return nil
}

func Db() *gorm.DB {
	return db
}

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"web_app/dao/mysqlc/model"
)

func main() {
	dsn := fmt.Sprintf("%s:%d@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4&parseTime=True&loc=Local",
		"root", 142212)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("mysql  conn err:", err)
		return
	}
	var communitys []model.Community
	for x := 3; x < 8; x++ {
		community := model.Community{
			CommunityId:   int64(x),
			CommunityName: fmt.Sprintf("测试name%s", strconv.Itoa(x)),
			Introducion:   fmt.Sprintf("测试描述%s", strconv.Itoa(x)),
		}
		communitys = append(communitys, community)
	}
	db.Create(&communitys)
}

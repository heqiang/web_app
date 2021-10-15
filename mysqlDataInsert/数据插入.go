package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	"web_app/dao/mysqlc/model"
)

var db *gorm.DB

func main() {
	var err error
	dsn := fmt.Sprintf("%s:%d@tcp(127.0.0.1:3306)/sqk_demo?charset=utf8mb4&parseTime=True&loc=Local",
		"root", 142212)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
	//community()
	//communityDetail()
	//QueryByName("admin")
	//InsertPost()
	QueryPost()

}
func community() {
	var communitys []model.Community
	for x := 0; x < 8; x++ {
		community := model.Community{
			CommunityId:   int64(x),
			CommunityName: fmt.Sprintf("测试name%s", strconv.Itoa(x)),
			Introducion:   fmt.Sprintf("测试描述%s", strconv.Itoa(x)),
		}
		communitys = append(communitys, community)
	}
	db.Create(&communitys)

}
func communityDetail() {
	var communityDetails []model.Community
	for x := 0; x < 8; x++ {
		communityDetail := model.Community{
			CommunityName: fmt.Sprintf("Name%s号", strconv.Itoa(x)),
			Introducion: fmt.Sprintf("介绍描述%s"+
				"", strconv.Itoa(x)),
		}
		communityDetails = append(communityDetails, communityDetail)
	}
	db.Create(&communityDetails)
}

func QueryByName(name string) {
	var user model.User
	db.Where("username=?", name).Take(&user)
	fmt.Println(user)
}

func InsertPost() {
	var posts []model.Post
	for x := 0; x < 20; x++ {
		post := model.Post{
			Title:   fmt.Sprintf("title%s", strconv.Itoa(x)),
			Content: fmt.Sprintf("content%s", strconv.Itoa(x)),
		}
		posts = append(posts, post)
	}
	db.Create(&posts)

}

func QueryPost() {
	posts := []*model.Post{}
	db.Offset(1).Limit(10).Find(&posts)
	//a := &posts
	for _, x := range posts {
		fmt.Println(*x)
	}
}

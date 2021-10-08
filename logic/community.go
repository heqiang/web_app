package logic

import (
	"web_app/dao/mysqlc"
	"web_app/dao/mysqlc/model"
)

func GetCommunityList(comm model.Community) (data []model.Community, err error) {
	data = mysqlc.QueryAllCommunitys(comm)
	return data, nil
}

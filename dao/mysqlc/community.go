package mysqlc

import "web_app/dao/mysqlc/model"

func QueryAllCommunitys(comm model.Community) (comms []model.Community) {
	db.Where(&model.Community{CommunityId: comm.CommunityId, CommunityName: comm.CommunityName}).Find(&comms)
	return
}

package service

import (
	"web_app/controller/community/communitymodel"
	"web_app/dao/mysqlc"
)

type Community struct{}

func (comm *Community) GetCommunityDetail(id int64) (commdetail communitymodel.CommunityDetail, err error) {
	commdetail, err = mysqlc.QueryCommunityDetail(id)
	if err != nil {
		return
	}
	return
}

func (comm *Community) GetCommunityList() (data []communitymodel.Community, err error) {
	return mysqlc.QueryAllCommunitys()
}

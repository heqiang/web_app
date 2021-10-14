package logic

import (
	"web_app/controller/community/communitymodel"
	"web_app/dao/mysqlc"
)

func GetCommunityList() (data []communitymodel.Community, err error) {
	return mysqlc.QueryAllCommunitys()
}

func GetCommunityDetail(id int64) (commdetail communitymodel.CommunityDetail, err error) {
	commdetail, err = mysqlc.QueryCommunityDetail(id)
	if err != nil {
		return
	}
	return

}

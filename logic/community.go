package logic

import (
	"web_app/dao/mysqlc"
	"web_app/logic/modeltype"
)

func GetCommunityList() (data []modeltype.Community, err error) {
	return mysqlc.QueryAllCommunitys()
}

func GetCommunityDetail(id int64) (commdetail modeltype.CommunityDetail, err error) {
	commdetail, err = mysqlc.QueryCommunityDetail(id)
	if err != nil {
		return
	}
	return

}

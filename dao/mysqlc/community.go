package mysqlc

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"web_app/controller/community/communitymodel"
	"web_app/dao/mysqlc/model"
)

func QueryAllCommunitys() (comms []communitymodel.Community, err error) {

	res := db.Select("CommunityId", "CommunityName").Find(&comms)
	if res.RowsAffected == 0 {
		zap.L().Warn("no community in db")
		err = errors.New("communityList is nill")
		return
	}
	return
}

func QueryCommunityDetail(id int64) (commdetail communitymodel.CommunityDetail, err error) {
	queryRes := db.Where("id=?", id).Find(&commdetail)
	fmt.Println(queryRes.RowsAffected)
	if queryRes.RowsAffected == 0 {
		zap.L().Warn("no such community")
		err = errors.New("no such community")
		return
	}
	return

}

func QueryByCommId(id int64) (community *model.Community) {
	community = new(model.Community)
	db.Where("communityid=?", id).Take(&community)
	return
}

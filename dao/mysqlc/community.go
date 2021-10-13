package mysqlc

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"web_app/logic/modeltype"
)

func QueryAllCommunitys() (comms []modeltype.Community, err error) {

	res := db.Select("CommunityId", "CommunityName").Find(&comms)
	if res.RowsAffected == 0 {
		zap.L().Warn("no community in db")
		err = errors.New("communityList is nill")
		return
	}
	return
}

func QueryCommunityDetail(id int64) (commdetail modeltype.CommunityDetail, err error) {
	queryRes := db.Where("id=?", id).Find(&commdetail)
	fmt.Println(queryRes.RowsAffected)
	if queryRes.RowsAffected == 0 {
		zap.L().Warn("no such community")
		err = errors.New("no such community")
		return
	}
	return

}

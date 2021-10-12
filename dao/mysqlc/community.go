package mysqlc

import (
	"errors"
	"go.uber.org/zap"
	"web_app/logic"
)

func QueryAllCommunitys() (comms []logic.Community, err error) {

	res := db.Select("CommunityId", "CommunityName").Find(&comms)
	if res.RowsAffected == 0 {
		zap.L().Warn("no community in db")
		err = errors.New("communityList is nill")
	}
	return
}

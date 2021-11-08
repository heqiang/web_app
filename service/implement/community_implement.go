package implement

import "web_app/controller/community/communitymodel"

type Community interface {
	GetCommunityList() (data []communitymodel.Community, err error)
	GetCommunityDetail(id int64) (commdetail communitymodel.CommunityDetail, err error)
}

package v20140526

import "aliyuncs/core"

type DeleteHaVipRequest struct {
	core.Request
}

func NewDeleteHaVipRequest() *DeleteHaVipRequest {
	r := new(DeleteHaVipRequest)
	r.Init("Ecs", "2014-05-26", "DeleteHaVip")
	return r
}

func (this *DeleteHaVipRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DeleteHaVipRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DeleteHaVipRequest) SetHaVipId(v string) {
	this.Set("HaVipId", v)
}

func (this *DeleteHaVipRequest) GetHaVipId() string {
	return this.Get("HaVipId")
}

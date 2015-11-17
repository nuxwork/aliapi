package model

import "aliyuncs/ecs"

type DeleteHaVipRequest struct {
	ecs.Request
}

func NewDeleteHaVipRequest() *DeleteHaVipRequest{
	r := new(DeleteHaVipRequest)
	r.Set("Action", "DeleteHaVip")
	return r
}

func (this *DeleteHaVipRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DeleteHaVipRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DeleteHaVipRequest) SetHaVipId(v string){
	this.Set("HaVipId", v)
}

func (this *DeleteHaVipRequest) GetHaVipId() string {
	return this.Get("HaVipId")
}

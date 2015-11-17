package model

import "aliyuncs/ecs"

type AssociateHaVipRequest struct {
	ecs.Request
}

func NewAssociateHaVipRequest() *AssociateHaVipRequest{
	r := new(AssociateHaVipRequest)
	r.Set("Action", "AssociateHaVip")
	return r
}

func (this *AssociateHaVipRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *AssociateHaVipRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *AssociateHaVipRequest) SetHaVipId(v string){
	this.Set("HaVipId", v)
}

func (this *AssociateHaVipRequest) GetHaVipId() string {
	return this.Get("HaVipId")
}

func (this *AssociateHaVipRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *AssociateHaVipRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

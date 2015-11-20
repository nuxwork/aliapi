package v20140526

import "aliyuncs/core"

type AssociateHaVipRequest struct {
	core.Request
}

func NewAssociateHaVipRequest() *AssociateHaVipRequest {
	r := new(AssociateHaVipRequest)
	r.Init("Ecs", "2014-05-26", "AssociateHaVip")
	return r
}

func (this *AssociateHaVipRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *AssociateHaVipRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *AssociateHaVipRequest) SetHaVipId(v string) {
	this.Set("HaVipId", v)
}

func (this *AssociateHaVipRequest) GetHaVipId() string {
	return this.Get("HaVipId")
}

func (this *AssociateHaVipRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *AssociateHaVipRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

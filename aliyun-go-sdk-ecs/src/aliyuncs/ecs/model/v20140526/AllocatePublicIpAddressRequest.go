package v20140526

import "aliyuncs/core"

type AllocatePublicIpAddressRequest struct {
	core.Request
}

func NewAllocatePublicIpAddressRequest() *AllocatePublicIpAddressRequest {
	r := new(AllocatePublicIpAddressRequest)
	r.Init("Ecs", "2014-05-26", "AllocatePublicIpAddress")
	return r
}

func (this *AllocatePublicIpAddressRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *AllocatePublicIpAddressRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

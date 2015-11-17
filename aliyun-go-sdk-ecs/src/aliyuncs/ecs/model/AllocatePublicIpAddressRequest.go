package model

import "aliyuncs/ecs"

type AllocatePublicIpAddressRequest struct {
	ecs.Request
}

func NewAllocatePublicIpAddressRequest() *AllocatePublicIpAddressRequest{
	r := new(AllocatePublicIpAddressRequest)
	r.Set("Action", "AllocatePublicIpAddress")
	return r
}

func (this *AllocatePublicIpAddressRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *AllocatePublicIpAddressRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

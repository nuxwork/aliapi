package model

import "aliyuncs/ecs"

type ReleaseEipAddressRequest struct {
	ecs.Request
}

func NewReleaseEipAddressRequest() *ReleaseEipAddressRequest{
	r := new(ReleaseEipAddressRequest)
	r.Set("Action", "ReleaseEipAddress")
	return r
}

func (this *ReleaseEipAddressRequest) SetAllocationId(v string){
	this.Set("AllocationId", v)
}

func (this *ReleaseEipAddressRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

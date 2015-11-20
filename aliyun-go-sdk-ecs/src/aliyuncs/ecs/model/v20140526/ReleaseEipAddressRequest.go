package v20140526

import "aliyuncs/core"

type ReleaseEipAddressRequest struct {
	core.Request
}

func NewReleaseEipAddressRequest() *ReleaseEipAddressRequest {
	r := new(ReleaseEipAddressRequest)
	r.Init("Ecs", "2014-05-26", "ReleaseEipAddress")
	return r
}

func (this *ReleaseEipAddressRequest) SetAllocationId(v string) {
	this.Set("AllocationId", v)
}

func (this *ReleaseEipAddressRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

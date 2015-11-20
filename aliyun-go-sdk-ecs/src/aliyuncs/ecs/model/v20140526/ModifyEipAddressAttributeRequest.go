package v20140526

import "aliyuncs/core"

type ModifyEipAddressAttributeRequest struct {
	core.Request
}

func NewModifyEipAddressAttributeRequest() *ModifyEipAddressAttributeRequest {
	r := new(ModifyEipAddressAttributeRequest)
	r.Init("Ecs", "2014-05-26", "ModifyEipAddressAttribute")
	return r
}

func (this *ModifyEipAddressAttributeRequest) SetAllocationId(v string) {
	this.Set("AllocationId", v)
}

func (this *ModifyEipAddressAttributeRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

func (this *ModifyEipAddressAttributeRequest) SetBandwidth(v string) {
	this.Set("Bandwidth", v)
}

func (this *ModifyEipAddressAttributeRequest) GetBandwidth() string {
	return this.Get("Bandwidth")
}

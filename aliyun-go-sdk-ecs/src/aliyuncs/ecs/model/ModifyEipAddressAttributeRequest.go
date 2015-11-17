package model

import "aliyuncs/ecs"

type ModifyEipAddressAttributeRequest struct {
	ecs.Request
}

func NewModifyEipAddressAttributeRequest() *ModifyEipAddressAttributeRequest{
	r := new(ModifyEipAddressAttributeRequest)
	r.Set("Action", "ModifyEipAddressAttribute")
	return r
}

func (this *ModifyEipAddressAttributeRequest) SetAllocationId(v string){
	this.Set("AllocationId", v)
}

func (this *ModifyEipAddressAttributeRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

func (this *ModifyEipAddressAttributeRequest) SetBandwidth(v string){
	this.Set("Bandwidth", v)
}

func (this *ModifyEipAddressAttributeRequest) GetBandwidth() string {
	return this.Get("Bandwidth")
}

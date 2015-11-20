package v20140526

import "aliyuncs/core"

type AssociateEipAddressRequest struct {
	core.Request
}

func NewAssociateEipAddressRequest() *AssociateEipAddressRequest {
	r := new(AssociateEipAddressRequest)
	r.Init("Ecs", "2014-05-26", "AssociateEipAddress")
	return r
}

func (this *AssociateEipAddressRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *AssociateEipAddressRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *AssociateEipAddressRequest) SetAllocationId(v string) {
	this.Set("AllocationId", v)
}

func (this *AssociateEipAddressRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

func (this *AssociateEipAddressRequest) SetInstanceType(v string) {
	this.Set("InstanceType", v)
}

func (this *AssociateEipAddressRequest) GetInstanceType() string {
	return this.Get("InstanceType")
}

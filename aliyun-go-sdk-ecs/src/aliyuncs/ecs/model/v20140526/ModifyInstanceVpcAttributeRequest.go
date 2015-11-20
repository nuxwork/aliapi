package v20140526

import "aliyuncs/core"

type ModifyInstanceVpcAttributeRequest struct {
	core.Request
}

func NewModifyInstanceVpcAttributeRequest() *ModifyInstanceVpcAttributeRequest {
	r := new(ModifyInstanceVpcAttributeRequest)
	r.Init("Ecs", "2014-05-26", "ModifyInstanceVpcAttribute")
	return r
}

func (this *ModifyInstanceVpcAttributeRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *ModifyInstanceVpcAttributeRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *ModifyInstanceVpcAttributeRequest) SetVSwitchId(v string) {
	this.Set("VSwitchId", v)
}

func (this *ModifyInstanceVpcAttributeRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

func (this *ModifyInstanceVpcAttributeRequest) SetPrivateIpAddress(v string) {
	this.Set("PrivateIpAddress", v)
}

func (this *ModifyInstanceVpcAttributeRequest) GetPrivateIpAddress() string {
	return this.Get("PrivateIpAddress")
}

package model

import "aliyuncs/ecs"

type ModifyInstanceVpcAttributeRequest struct {
	ecs.Request
}

func NewModifyInstanceVpcAttributeRequest() *ModifyInstanceVpcAttributeRequest{
	r := new(ModifyInstanceVpcAttributeRequest)
	r.Set("Action", "ModifyInstanceVpcAttribute")
	return r
}

func (this *ModifyInstanceVpcAttributeRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *ModifyInstanceVpcAttributeRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *ModifyInstanceVpcAttributeRequest) SetVSwitchId(v string){
	this.Set("VSwitchId", v)
}

func (this *ModifyInstanceVpcAttributeRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

func (this *ModifyInstanceVpcAttributeRequest) SetPrivateIpAddress(v string){
	this.Set("PrivateIpAddress", v)
}

func (this *ModifyInstanceVpcAttributeRequest) GetPrivateIpAddress() string {
	return this.Get("PrivateIpAddress")
}

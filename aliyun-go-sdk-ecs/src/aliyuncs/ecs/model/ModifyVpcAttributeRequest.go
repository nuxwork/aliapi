package model

import "aliyuncs/ecs"

type ModifyVpcAttributeRequest struct {
	ecs.Request
}

func NewModifyVpcAttributeRequest() *ModifyVpcAttributeRequest{
	r := new(ModifyVpcAttributeRequest)
	r.Set("Action", "ModifyVpcAttribute")
	return r
}

func (this *ModifyVpcAttributeRequest) SetVpcId(v string){
	this.Set("VpcId", v)
}

func (this *ModifyVpcAttributeRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *ModifyVpcAttributeRequest) SetVpcName(v string){
	this.Set("VpcName", v)
}

func (this *ModifyVpcAttributeRequest) GetVpcName() string {
	return this.Get("VpcName")
}

func (this *ModifyVpcAttributeRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *ModifyVpcAttributeRequest) GetDescription() string {
	return this.Get("Description")
}

package v20140526

import "aliyuncs/core"

type ModifyVpcAttributeRequest struct {
	core.Request
}

func NewModifyVpcAttributeRequest() *ModifyVpcAttributeRequest {
	r := new(ModifyVpcAttributeRequest)
	r.Init("Ecs", "2014-05-26", "ModifyVpcAttribute")
	return r
}

func (this *ModifyVpcAttributeRequest) SetVpcId(v string) {
	this.Set("VpcId", v)
}

func (this *ModifyVpcAttributeRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *ModifyVpcAttributeRequest) SetVpcName(v string) {
	this.Set("VpcName", v)
}

func (this *ModifyVpcAttributeRequest) GetVpcName() string {
	return this.Get("VpcName")
}

func (this *ModifyVpcAttributeRequest) SetDescription(v string) {
	this.Set("Description", v)
}

func (this *ModifyVpcAttributeRequest) GetDescription() string {
	return this.Get("Description")
}

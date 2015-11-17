package model

import "aliyuncs/ecs"

type ModifyVSwitchAttributeRequest struct {
	ecs.Request
}

func NewModifyVSwitchAttributeRequest() *ModifyVSwitchAttributeRequest{
	r := new(ModifyVSwitchAttributeRequest)
	r.Set("Action", "ModifyVSwitchAttribute")
	return r
}

func (this *ModifyVSwitchAttributeRequest) SetVSwitchId(v string){
	this.Set("VSwitchId", v)
}

func (this *ModifyVSwitchAttributeRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

func (this *ModifyVSwitchAttributeRequest) SetVSwitchName(v string){
	this.Set("VSwitchName", v)
}

func (this *ModifyVSwitchAttributeRequest) GetVSwitchName() string {
	return this.Get("VSwitchName")
}

func (this *ModifyVSwitchAttributeRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *ModifyVSwitchAttributeRequest) GetDescription() string {
	return this.Get("Description")
}

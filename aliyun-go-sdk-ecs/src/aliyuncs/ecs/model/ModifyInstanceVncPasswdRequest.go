package model

import "aliyuncs/ecs"

type ModifyInstanceVncPasswdRequest struct {
	ecs.Request
}

func NewModifyInstanceVncPasswdRequest() *ModifyInstanceVncPasswdRequest{
	r := new(ModifyInstanceVncPasswdRequest)
	r.Set("Action", "ModifyInstanceVncPasswd")
	return r
}

func (this *ModifyInstanceVncPasswdRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *ModifyInstanceVncPasswdRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *ModifyInstanceVncPasswdRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *ModifyInstanceVncPasswdRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *ModifyInstanceVncPasswdRequest) SetVncPassword(v string){
	this.Set("VncPassword", v)
}

func (this *ModifyInstanceVncPasswdRequest) GetVncPassword() string {
	return this.Get("VncPassword")
}

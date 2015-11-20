package v20140526

import "aliyuncs/core"

type ModifyInstanceAttributeRequest struct {
	core.Request
}

func NewModifyInstanceAttributeRequest() *ModifyInstanceAttributeRequest {
	r := new(ModifyInstanceAttributeRequest)
	r.Init("Ecs", "2014-05-26", "ModifyInstanceAttribute")
	return r
}

func (this *ModifyInstanceAttributeRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *ModifyInstanceAttributeRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *ModifyInstanceAttributeRequest) SetInstanceName(v string) {
	this.Set("InstanceName", v)
}

func (this *ModifyInstanceAttributeRequest) GetInstanceName() string {
	return this.Get("InstanceName")
}

func (this *ModifyInstanceAttributeRequest) SetDescription(v string) {
	this.Set("Description", v)
}

func (this *ModifyInstanceAttributeRequest) GetDescription() string {
	return this.Get("Description")
}

func (this *ModifyInstanceAttributeRequest) SetPassword(v string) {
	this.Set("Password", v)
}

func (this *ModifyInstanceAttributeRequest) GetPassword() string {
	return this.Get("Password")
}

func (this *ModifyInstanceAttributeRequest) SetHostName(v string) {
	this.Set("HostName", v)
}

func (this *ModifyInstanceAttributeRequest) GetHostName() string {
	return this.Get("HostName")
}

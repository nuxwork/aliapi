package v20140526

import "aliyuncs/core"

type ModifyVRouterAttributeRequest struct {
	core.Request
}

func NewModifyVRouterAttributeRequest() *ModifyVRouterAttributeRequest {
	r := new(ModifyVRouterAttributeRequest)
	r.Init("Ecs", "2014-05-26", "ModifyVRouterAttribute")
	return r
}

func (this *ModifyVRouterAttributeRequest) SetVRouterId(v string) {
	this.Set("VRouterId", v)
}

func (this *ModifyVRouterAttributeRequest) GetVRouterId() string {
	return this.Get("VRouterId")
}

func (this *ModifyVRouterAttributeRequest) SetVRouterName(v string) {
	this.Set("VRouterName", v)
}

func (this *ModifyVRouterAttributeRequest) GetVRouterName() string {
	return this.Get("VRouterName")
}

func (this *ModifyVRouterAttributeRequest) SetDescription(v string) {
	this.Set("Description", v)
}

func (this *ModifyVRouterAttributeRequest) GetDescription() string {
	return this.Get("Description")
}

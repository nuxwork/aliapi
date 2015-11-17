package model

import "aliyuncs/ecs"

type ModifyImageAttributeRequest struct {
	ecs.Request
}

func NewModifyImageAttributeRequest() *ModifyImageAttributeRequest{
	r := new(ModifyImageAttributeRequest)
	r.Set("Action", "ModifyImageAttribute")
	return r
}

func (this *ModifyImageAttributeRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *ModifyImageAttributeRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *ModifyImageAttributeRequest) SetImageId(v string){
	this.Set("ImageId", v)
}

func (this *ModifyImageAttributeRequest) GetImageId() string {
	return this.Get("ImageId")
}

func (this *ModifyImageAttributeRequest) SetImageName(v string){
	this.Set("ImageName", v)
}

func (this *ModifyImageAttributeRequest) GetImageName() string {
	return this.Get("ImageName")
}

func (this *ModifyImageAttributeRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *ModifyImageAttributeRequest) GetDescription() string {
	return this.Get("Description")
}

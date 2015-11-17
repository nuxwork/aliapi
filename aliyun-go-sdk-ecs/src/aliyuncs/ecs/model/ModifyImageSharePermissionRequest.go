package model

import "aliyuncs/ecs"

type ModifyImageSharePermissionRequest struct {
	ecs.Request
}

func NewModifyImageSharePermissionRequest() *ModifyImageSharePermissionRequest{
	r := new(ModifyImageSharePermissionRequest)
	r.Set("Action", "ModifyImageSharePermission")
	return r
}

func (this *ModifyImageSharePermissionRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *ModifyImageSharePermissionRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *ModifyImageSharePermissionRequest) SetImageId(v string){
	this.Set("ImageId", v)
}

func (this *ModifyImageSharePermissionRequest) GetImageId() string {
	return this.Get("ImageId")
}

func (this *ModifyImageSharePermissionRequest) SetAddAccountn(v string){
	this.Set("AddAccount.n", v)
}

func (this *ModifyImageSharePermissionRequest) GetAddAccountn() string {
	return this.Get("AddAccount.n")
}

func (this *ModifyImageSharePermissionRequest) SetRemoveAccountn(v string){
	this.Set("RemoveAccount.n", v)
}

func (this *ModifyImageSharePermissionRequest) GetRemoveAccountn() string {
	return this.Get("RemoveAccount.n")
}

package v20140526

import "strconv"
import "aliyuncs/core"

type ModifyImageSharePermissionRequest struct {
	core.Request
}

func NewModifyImageSharePermissionRequest() *ModifyImageSharePermissionRequest {
	r := new(ModifyImageSharePermissionRequest)
	r.Init("Ecs", "2014-05-26", "ModifyImageSharePermission")
	return r
}

func (this *ModifyImageSharePermissionRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *ModifyImageSharePermissionRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *ModifyImageSharePermissionRequest) SetImageId(v string) {
	this.Set("ImageId", v)
}

func (this *ModifyImageSharePermissionRequest) GetImageId() string {
	return this.Get("ImageId")
}

func (this *ModifyImageSharePermissionRequest) SetAddAccountn(v string, n int) {
	k := "AddAccount." + strconv.Itoa(n)
	this.Set(k, v)
}

func (this *ModifyImageSharePermissionRequest) GetAddAccountn(n int) string {
	k := "AddAccount." + strconv.Itoa(n)
	return this.Get(k)
}

func (this *ModifyImageSharePermissionRequest) SetRemoveAccountn(v string, n int) {
	k := "RemoveAccount." + strconv.Itoa(n)
	this.Set(k, v)
}

func (this *ModifyImageSharePermissionRequest) GetRemoveAccountn(n int) string {
	k := "RemoveAccount." + strconv.Itoa(n)
	return this.Get(k)
}

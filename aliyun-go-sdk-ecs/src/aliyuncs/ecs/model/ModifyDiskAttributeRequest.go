package model

import "aliyuncs/ecs"

type ModifyDiskAttributeRequest struct {
	ecs.Request
}

func NewModifyDiskAttributeRequest() *ModifyDiskAttributeRequest{
	r := new(ModifyDiskAttributeRequest)
	r.Set("Action", "ModifyDiskAttribute")
	return r
}

func (this *ModifyDiskAttributeRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *ModifyDiskAttributeRequest) GetDiskId() string {
	return this.Get("DiskId")
}

func (this *ModifyDiskAttributeRequest) SetDiskName(v string){
	this.Set("DiskName", v)
}

func (this *ModifyDiskAttributeRequest) GetDiskName() string {
	return this.Get("DiskName")
}

func (this *ModifyDiskAttributeRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *ModifyDiskAttributeRequest) GetDescription() string {
	return this.Get("Description")
}

func (this *ModifyDiskAttributeRequest) SetDeleteWithInstance(v bool){
	this.SetBool("DeleteWithInstance", v)
}

func (this *ModifyDiskAttributeRequest) GetDeleteWithInstance() bool {
	return this.GetBool("DeleteWithInstance")
}

func (this *ModifyDiskAttributeRequest) SetDeleteAutoSnapshot(v bool){
	this.SetBool("DeleteAutoSnapshot", v)
}

func (this *ModifyDiskAttributeRequest) GetDeleteAutoSnapshot() bool {
	return this.GetBool("DeleteAutoSnapshot")
}

func (this *ModifyDiskAttributeRequest) SetEnableAutoSnapshot(v bool){
	this.SetBool("EnableAutoSnapshot", v)
}

func (this *ModifyDiskAttributeRequest) GetEnableAutoSnapshot() bool {
	return this.GetBool("EnableAutoSnapshot")
}

package model

import "aliyuncs/ecs"

type ReplaceSystemDiskRequest struct {
	ecs.Request
}

func NewReplaceSystemDiskRequest() *ReplaceSystemDiskRequest{
	r := new(ReplaceSystemDiskRequest)
	r.Set("Action", "ReplaceSystemDisk")
	return r
}

func (this *ReplaceSystemDiskRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *ReplaceSystemDiskRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *ReplaceSystemDiskRequest) SetImageId(v string){
	this.Set("ImageId", v)
}

func (this *ReplaceSystemDiskRequest) GetImageId() string {
	return this.Get("ImageId")
}

func (this *ReplaceSystemDiskRequest) SetClientToken(v string){
	this.Set("ClientToken", v)
}

func (this *ReplaceSystemDiskRequest) GetClientToken() string {
	return this.Get("ClientToken")
}

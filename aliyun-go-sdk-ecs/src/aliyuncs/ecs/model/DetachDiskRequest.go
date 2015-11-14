package model

import "aliyuncs/ecs"

type DetachDiskRequest struct {
	ecs.Request
}

func NewDetachDiskRequest() *DetachDiskRequest{
	r := new(DetachDiskRequest)
	r.Set("Action", "DetachDisk")
	return r
}

func (this *DetachDiskRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *DetachDiskRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *DetachDiskRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *DetachDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

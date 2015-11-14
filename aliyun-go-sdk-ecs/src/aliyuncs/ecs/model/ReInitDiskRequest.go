package model

import "aliyuncs/ecs"

type ReInitDiskRequest struct {
	ecs.Request
}

func NewReInitDiskRequest() *ReInitDiskRequest{
	r := new(ReInitDiskRequest)
	r.Set("Action", "ReInitDisk")
	return r
}

func (this *ReInitDiskRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *ReInitDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

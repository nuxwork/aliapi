package model

import "aliyuncs/ecs"

type DeleteDiskRequest struct {
	ecs.Request
}

func NewDeleteDiskRequest() *DeleteDiskRequest{
	r := new(DeleteDiskRequest)
	r.Set("Action", "DeleteDisk")
	return r
}

func (this *DeleteDiskRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *DeleteDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

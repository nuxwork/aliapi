package model

import "aliyuncs/ecs"

type ResetDiskRequest struct {
	ecs.Request
}

func NewResetDiskRequest() *ResetDiskRequest{
	r := new(ResetDiskRequest)
	r.Set("Action", "ResetDisk")
	return r
}

func (this *ResetDiskRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *ResetDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

func (this *ResetDiskRequest) SetSnapshotId(v string){
	this.Set("SnapshotId", v)
}

func (this *ResetDiskRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}

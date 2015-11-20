package v20140526

import "aliyuncs/core"

type ResetDiskRequest struct {
	core.Request
}

func NewResetDiskRequest() *ResetDiskRequest {
	r := new(ResetDiskRequest)
	r.Init("Ecs", "2014-05-26", "ResetDisk")
	return r
}

func (this *ResetDiskRequest) SetDiskId(v string) {
	this.Set("DiskId", v)
}

func (this *ResetDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

func (this *ResetDiskRequest) SetSnapshotId(v string) {
	this.Set("SnapshotId", v)
}

func (this *ResetDiskRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}

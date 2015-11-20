package v20140526

import "aliyuncs/core"

type DetachDiskRequest struct {
	core.Request
}

func NewDetachDiskRequest() *DetachDiskRequest {
	r := new(DetachDiskRequest)
	r.Init("Ecs", "2014-05-26", "DetachDisk")
	return r
}

func (this *DetachDiskRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *DetachDiskRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *DetachDiskRequest) SetDiskId(v string) {
	this.Set("DiskId", v)
}

func (this *DetachDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

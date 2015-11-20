package v20140526

import "aliyuncs/core"

type ReInitDiskRequest struct {
	core.Request
}

func NewReInitDiskRequest() *ReInitDiskRequest {
	r := new(ReInitDiskRequest)
	r.Init("Ecs", "2014-05-26", "ReInitDisk")
	return r
}

func (this *ReInitDiskRequest) SetDiskId(v string) {
	this.Set("DiskId", v)
}

func (this *ReInitDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

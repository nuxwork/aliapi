package v20140526

import "aliyuncs/core"

type DeleteDiskRequest struct {
	core.Request
}

func NewDeleteDiskRequest() *DeleteDiskRequest {
	r := new(DeleteDiskRequest)
	r.Init("Ecs", "2014-05-26", "DeleteDisk")
	return r
}

func (this *DeleteDiskRequest) SetDiskId(v string) {
	this.Set("DiskId", v)
}

func (this *DeleteDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

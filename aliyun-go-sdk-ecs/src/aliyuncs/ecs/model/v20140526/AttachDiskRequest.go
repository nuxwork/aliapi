package v20140526

import "aliyuncs/core"

type AttachDiskRequest struct {
	core.Request
}

func NewAttachDiskRequest() *AttachDiskRequest {
	r := new(AttachDiskRequest)
	r.Init("Ecs", "2014-05-26", "AttachDisk")
	return r
}

func (this *AttachDiskRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *AttachDiskRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *AttachDiskRequest) SetDiskId(v string) {
	this.Set("DiskId", v)
}

/*
空表示由系统默认分配，/dev/xvdb 开始到 /dev/xvdz
默认值：空
*/
func (this *AttachDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

func (this *AttachDiskRequest) SetDevice(v string) {
	this.Set("Device", v)
}

func (this *AttachDiskRequest) GetDevice() string {
	return this.Get("Device")
}

func (this *AttachDiskRequest) SetDeleteWithInstance(v bool) {
	this.SetBool("DeleteWithInstance", v)
}

func (this *AttachDiskRequest) GetDeleteWithInstance() bool {
	return this.GetBool("DeleteWithInstance")
}

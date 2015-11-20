package v20140526

import "aliyuncs/core"

type CreateDiskRequest struct {
	core.Request
}

func NewCreateDiskRequest() *CreateDiskRequest {
	r := new(CreateDiskRequest)
	r.Init("Ecs", "2014-05-26", "CreateDisk")
	return r
}

func (this *CreateDiskRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *CreateDiskRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CreateDiskRequest) SetZoneId(v string) {
	this.Set("ZoneId", v)
}

func (this *CreateDiskRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *CreateDiskRequest) SetDiskName(v string) {
	this.Set("DiskName", v)
}

func (this *CreateDiskRequest) GetDiskName() string {
	return this.Get("DiskName")
}

func (this *CreateDiskRequest) SetDescription(v string) {
	this.Set("Description", v)
}

func (this *CreateDiskRequest) GetDescription() string {
	return this.Get("Description")
}

/*
数据盘的磁盘种类
可选值：
cloud – 普通云盘
cloud_efficiency – 高效云盘
cloud_ssd – SSD云盘
默认值：cloud
*/
func (this *CreateDiskRequest) SetDiskCategory(v string) {
	this.Set("DiskCategory", v)
}

func (this *CreateDiskRequest) GetDiskCategory() string {
	return this.Get("DiskCategory")
}

/*
容量大小，以GB为单位：
cloud：5 ~ 2000
cloud_efficiency：20 ~ 2048
cloud_ssd：20 ~ 1024
指定该参数后，Size大小必须 ≥ 指定快照SnapshotId的大小。
*/
func (this *CreateDiskRequest) SetSize(v int) {
	this.SetInt("Size", v)
}

func (this *CreateDiskRequest) GetSize() int {
	return this.GetInt("Size")
}

func (this *CreateDiskRequest) SetSnapshotId(v string) {
	this.Set("SnapshotId", v)
}

func (this *CreateDiskRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}

func (this *CreateDiskRequest) SetClientToken(v string) {
	this.Set("ClientToken", v)
}

func (this *CreateDiskRequest) GetClientToken() string {
	return this.Get("ClientToken")
}

package v20140526

import "aliyuncs/core"

type DescribeSnapshotsRequest struct {
	core.Request
}

func NewDescribeSnapshotsRequest() *DescribeSnapshotsRequest {
	r := new(DescribeSnapshotsRequest)
	r.Init("Ecs", "2014-05-26", "DescribeSnapshots")
	return r
}

func (this *DescribeSnapshotsRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeSnapshotsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeSnapshotsRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *DescribeSnapshotsRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *DescribeSnapshotsRequest) SetDiskId(v string) {
	this.Set("DiskId", v)
}

func (this *DescribeSnapshotsRequest) GetDiskId() string {
	return this.Get("DiskId")
}

func (this *DescribeSnapshotsRequest) SetSnapshotIds(v []string) {
	this.SetStringArray("SnapshotIds", v)
}

func (this *DescribeSnapshotsRequest) GetSnapshotIds() []string {
	return this.GetStringArray("SnapshotIds")
}

func (this *DescribeSnapshotsRequest) SetSnapshotName(v string) {
	this.Set("SnapshotName", v)
}

func (this *DescribeSnapshotsRequest) GetSnapshotName() string {
	return this.Get("SnapshotName")
}

/*
快照状态，progressing | accomplished | failed | all
*/
func (this *DescribeSnapshotsRequest) SetStatus(v string) {
	this.Set("Status", v)
}

func (this *DescribeSnapshotsRequest) GetStatus() string {
	return this.Get("Status")
}

/*
快照类型，auto | user | all
*/
func (this *DescribeSnapshotsRequest) SetSnapshotType(v string) {
	this.Set("SnapshotType", v)
}

func (this *DescribeSnapshotsRequest) GetSnapshotType() string {
	return this.Get("SnapshotType")
}

/*
快照源磁盘的磁盘类型，System | Data
*/
func (this *DescribeSnapshotsRequest) SetSourceDiskType(v string) {
	this.Set("SourceDiskType", v)
}

func (this *DescribeSnapshotsRequest) GetSourceDiskType() string {
	return this.Get("SourceDiskType")
}

/*
有引用关系的资源类型，image | disk | image_disk | none
*/
func (this *DescribeSnapshotsRequest) SetUsage(v string) {
	this.Set("Usage", v)
}

func (this *DescribeSnapshotsRequest) GetUsage() string {
	return this.Get("Usage")
}

/*
磁盘状态列表的页码，起始值为 1，默认值为 1
*/
func (this *DescribeSnapshotsRequest) SetPageNumber(v int) {
	this.SetInt("PageNumber", v)
}

func (this *DescribeSnapshotsRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

/*
分页查询时设置的每页行数，最大值 100 行，默认为 10
*/
func (this *DescribeSnapshotsRequest) SetPageSize(v int) {
	this.SetInt("PageSize", v)
}

func (this *DescribeSnapshotsRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

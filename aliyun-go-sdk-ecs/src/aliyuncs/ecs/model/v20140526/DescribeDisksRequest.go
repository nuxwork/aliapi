package v20140526

import "aliyuncs/core"

type DescribeDisksRequest struct {
	core.Request
}

func NewDescribeDisksRequest() *DescribeDisksRequest {
	r := new(DescribeDisksRequest)
	r.Init("Ecs", "2014-05-26", "DescribeDisks")
	return r
}

func (this *DescribeDisksRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeDisksRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeDisksRequest) SetZoneId(v string) {
	this.Set("ZoneId", v)
}

func (this *DescribeDisksRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *DescribeDisksRequest) SetDiskIds(v []string) {
	this.SetStringArray("DiskIds", v)
}

func (this *DescribeDisksRequest) GetDiskIds() []string {
	return this.GetStringArray("DiskIds")
}

func (this *DescribeDisksRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *DescribeDisksRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

/*
all | system | data，默认值为 all。
*/
func (this *DescribeDisksRequest) SetDiskType(v string) {
	this.Set("DiskType", v)
}

func (this *DescribeDisksRequest) GetDiskType() string {
	return this.Get("DiskType")
}

/*
磁盘种类
all | cloud | cloud_efficiency | cloud_ssd | ephemeral | ephemeral_ssd。默认值为all
*/
func (this *DescribeDisksRequest) SetCategory(v string) {
	this.Set("Category", v)
}

func (this *DescribeDisksRequest) GetCategory() string {
	return this.Get("Category")
}

/*
磁盘状态
In_use | Available | Attaching | Detaching | Creating | ReIniting | All，默认值为 All。
*/
func (this *DescribeDisksRequest) SetStatus(v string) {
	this.Set("Status", v)
}

func (this *DescribeDisksRequest) GetStatus() string {
	return this.Get("Status")
}

func (this *DescribeDisksRequest) SetSnapshotId(v string) {
	this.Set("SnapshotId", v)
}

func (this *DescribeDisksRequest) GetSnapshotId() string {
	return this.Get("SnapshotId")
}

func (this *DescribeDisksRequest) SetDiskName(v string) {
	this.Set("DiskName", v)
}

func (this *DescribeDisksRequest) GetDiskName() string {
	return this.Get("DiskName")
}

func (this *DescribeDisksRequest) SetPortable(v bool) {
	this.SetBool("Portable", v)
}

func (this *DescribeDisksRequest) GetPortable() bool {
	return this.GetBool("Portable")
}

func (this *DescribeDisksRequest) SetDeleteWithInstance(v bool) {
	this.SetBool("DeleteWithInstance", v)
}

func (this *DescribeDisksRequest) GetDeleteWithInstance() bool {
	return this.GetBool("DeleteWithInstance")
}

func (this *DescribeDisksRequest) SetDeleteAutoSnapshot(v bool) {
	this.SetBool("DeleteAutoSnapshot", v)
}

func (this *DescribeDisksRequest) GetDeleteAutoSnapshot() bool {
	return this.GetBool("DeleteAutoSnapshot")
}

func (this *DescribeDisksRequest) SetEnableAutoSnapshot(v bool) {
	this.SetBool("EnableAutoSnapshot", v)
}

func (this *DescribeDisksRequest) GetEnableAutoSnapshot() bool {
	return this.GetBool("EnableAutoSnapshot")
}

/*
磁盘的付费方式。
PrePaid：预付费，即包年包月
PostPaid：后付费，即按量付费
*/
func (this *DescribeDisksRequest) SetDiskChargeType(v string) {
	this.Set("DiskChargeType", v)
}

func (this *DescribeDisksRequest) GetDiskChargeType() string {
	return this.Get("DiskChargeType")
}

/*
磁盘状态列表的页码，起始值为 1，默认值为 1
*/
func (this *DescribeDisksRequest) SetPageNumber(v int) {
	this.SetInt("PageNumber", v)
}

func (this *DescribeDisksRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

/*
分页查询时设置的每页行数，最大值 100 行，默认为 10
*/
func (this *DescribeDisksRequest) SetPageSize(v int) {
	this.SetInt("PageSize", v)
}

func (this *DescribeDisksRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

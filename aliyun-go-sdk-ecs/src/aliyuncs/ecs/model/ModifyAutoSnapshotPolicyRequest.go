package model

import "aliyuncs/ecs"

type ModifyAutoSnapshotPolicyRequest struct {
	ecs.Request
}

func NewModifyAutoSnapshotPolicyRequest() *ModifyAutoSnapshotPolicyRequest{
	r := new(ModifyAutoSnapshotPolicyRequest)
	r.Set("Action", "ModifyAutoSnapshotPolicy")
	return r
}

func (this *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyEnabled(v bool){
	this.SetBool("SystemDiskPolicyEnabled", v)
}

func (this *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyEnabled() bool {
	return this.GetBool("SystemDiskPolicyEnabled")
}

func (this *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyTimePeriod(v int){
	this.SetInt("SystemDiskPolicyTimePeriod", v)
}

/*
系统盘自动快照策略的时间段
4 个时间段：
1：1:00 - 7:00
2：7:00 - 13:00
3：13:00 - 19:00
4：19:00 - 1:00
默认值：无，表示保留当前值
*/
func (this *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyTimePeriod() int {
	return this.GetInt("SystemDiskPolicyTimePeriod")
}

/*
系统盘自动快照策略的保留天数 
可选值：1，2，3 
默认值：无，表示保留当前值
*/
func (this *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyRetentionDays(v int){
	this.SetInt("SystemDiskPolicyRetentionDays", v)
}

func (this *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyRetentionDays() int {
	return this.GetInt("SystemDiskPolicyRetentionDays")
}

func (this *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyRetentionLastWeek(v bool){
	this.SetBool("SystemDiskPolicyRetentionLastWeek", v)
}

func (this *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyRetentionLastWeek() bool {
	return this.GetBool("SystemDiskPolicyRetentionLastWeek")
}

func (this *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyEnabled(v bool){
	this.SetBool("DataDiskPolicyEnabled", v)
}

func (this *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyEnabled() bool {
	return this.GetBool("DataDiskPolicyEnabled")
}

/*
数据盘自动快照策略的时间段 
4 个时间段：
1：1:00 - 7:00
2：7:00 - 13:00
3：13:00 - 19:00
4：19:00 - 1:00
默认值：无，表示保留当前值
*/
func (this *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyTimePeriod(v int){
	this.SetInt("DataDiskPolicyTimePeriod", v)
}

func (this *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyTimePeriod() int {
	return this.GetInt("DataDiskPolicyTimePeriod")
}

func (this *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyRetentionDays(v int){
	this.SetInt("DataDiskPolicyRetentionDays", v)
}

func (this *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyRetentionDays() int {
	return this.GetInt("DataDiskPolicyRetentionDays")
}

func (this *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyRetentionLastWeek(v bool){
	this.SetBool("DataDiskPolicyRetentionLastWeek", v)
}

func (this *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyRetentionLastWeek() bool {
	return this.GetBool("DataDiskPolicyRetentionLastWeek")
}

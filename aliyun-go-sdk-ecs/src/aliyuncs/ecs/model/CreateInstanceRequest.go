package model

import "aliyuncs/ecs"

type CreateInstanceRequest struct {
	ecs.Request
}

func NewCreateInstanceRequest() *CreateInstanceRequest{
	r := new(CreateInstanceRequest)
	r.Set("Action", "CreateInstance")
	return r
}

func (this *CreateInstanceRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *CreateInstanceRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CreateInstanceRequest) SetZoneId(v string){
	this.Set("ZoneId", v)
}

func (this *CreateInstanceRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *CreateInstanceRequest) SetImageId(v string){
	this.Set("ImageId", v)
}

func (this *CreateInstanceRequest) GetImageId() string {
	return this.Get("ImageId")
}

/*
 * "Tiny", "Standard", "High Memory", "High CPU"
 */
func (this *CreateInstanceRequest) SetInstanceType(v string){
	this.Set("InstanceType", v)
}

func (this *CreateInstanceRequest) GetInstanceType() string {
	return this.Get("InstanceType")
}

func (this *CreateInstanceRequest) SetSecurityGroupId(v string){
	this.Set("SecurityGroupId", v)
}

func (this *CreateInstanceRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *CreateInstanceRequest) SetInstanceName(v string){
	this.Set("InstanceName", v)
}

func (this *CreateInstanceRequest) GetInstanceName() string {
	return this.Get("InstanceName")
}

func (this *CreateInstanceRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *CreateInstanceRequest) GetDescription() string {
	return this.Get("Description")
}

/* 
 * "PayByBandwidth", "PayByTraffic"
 * default is "PayByBandwidth"
 */
func (this *CreateInstanceRequest) SetInternetChargeType(v string){
	this.Set("InternetChargeType", v)
}

func (this *CreateInstanceRequest) GetInternetChargeType() string {
	return this.Get("InternetChargeType")
}

/*
 * [1, 200]
 * default is 200
 */
func (this *CreateInstanceRequest) SetInternetMaxBandwidthIn(v int){
	this.SetInt("InternetMaxBandwidthIn", v)
}

func (this *CreateInstanceRequest) GetInternetMaxBandwidthIn() int {
	return this.GetInt("InternetMaxBandwidthIn")
}

/*
 * PayByBandwidth [0, 100], default is 0
 * PayByTraffic [1, 100], not null
 */
func (this *CreateInstanceRequest) SetInternetMaxBandwidthOut(v int){
	this.SetInt("InternetMaxBandwidthOut", v)
}

func (this *CreateInstanceRequest) GetInternetMaxBandwidthOut() int {
	return this.GetInt("InternetMaxBandwidthOut")
}

func (this *CreateInstanceRequest) SetHostName(v string){
	this.Set("HostName", v)
}

func (this *CreateInstanceRequest) GetHostName() string {
	return this.Get("HostName")
}

func (this *CreateInstanceRequest) SetPassword(v string){
	this.Set("Password", v)
}

func (this *CreateInstanceRequest) GetPassword() string {
	return this.Get("Password")
}

/*
 * none：非IO优化， default is none
 * optimized：IO优化
 */
func (this *CreateInstanceRequest) SetIoOptimized(v string){
	this.Set("IoOptimized", v)
}

func (this *CreateInstanceRequest) GetIoOptimized() string {
	return this.Get("IoOptimized")
}

/*
	cloud – 普通云盘
	cloud_efficiency – 高效云盘
	cloud_ssd – SSD云盘
	ephemeral_ssd - 本地 SSD 盘
	默认值：cloud
*/
func (this *CreateInstanceRequest) SetSystemDiskCategory(v string){
	this.Set("SystemDisk.Category", v)
}

func (this *CreateInstanceRequest) GetSystemDiskCategory() string {
	return this.Get("SystemDisk.Category")
}

func (this *CreateInstanceRequest) SetSystemDiskDiskName(v string){
	this.Set("SystemDisk.DiskName", v)
}

func (this *CreateInstanceRequest) GetSystemDiskDiskName() string {
	return this.Get("SystemDisk.DiskName")
}

func (this *CreateInstanceRequest) SetSystemDiskDescription(v string){
	this.Set("SystemDisk.Description", v)
}

func (this *CreateInstanceRequest) GetSystemDiskDescription() string {
	return this.Get("SystemDisk.Description")
}

/*
数据盘 n 的磁盘大小（n 从 1 开始编号）。 以 GB 为单位，取值范围为：
cloud：5 ~ 2000
cloud_efficiency：20 ~ 2048
cloud_ssd：20 ~ 1024
ephemeral_ssd：5 ~ 800
指定该参数后，Size大小必须 ≥ 指定快照SnapshotId的大小。
*/
func (this *CreateInstanceRequest) SetDataDisknSize(v int){
	this.SetInt("DataDisk.n.Size", v)
}

func (this *CreateInstanceRequest) GetDataDisknSize() int {
	return this.GetInt("DataDisk.n.Size")
}

/*
数据盘 n 的磁盘种类 
可选值：
cloud – 普通云盘
cloud_efficiency – 高效云盘
cloud_ssd – SSD云盘
ephemeral_ssd - 本地 SSD 盘
默认值：cloud
*/
func (this *CreateInstanceRequest) SetDataDisknCategory(v string){
	this.Set("DataDisk.n.Category", v)
}

func (this *CreateInstanceRequest) GetDataDisknCategory() string {
	return this.Get("DataDisk.n.Category")
}

func (this *CreateInstanceRequest) SetDataDisknSnapshotId(v string){
	this.Set("DataDisk.n.SnapshotId", v)
}

func (this *CreateInstanceRequest) GetDataDisknSnapshotId() string {
	return this.Get("DataDisk.n.SnapshotId")
}

func (this *CreateInstanceRequest) SetDataDisknDiskName(v string){
	this.Set("DataDisk.n.DiskName", v)
}

func (this *CreateInstanceRequest) GetDataDisknDiskName() string {
	return this.Get("DataDisk.n.DiskName")
}

func (this *CreateInstanceRequest) SetDataDisknDescription(v string){
	this.Set("DataDisk.n.Description", v)
}

func (this *CreateInstanceRequest) GetDataDisknDescription() string {
	return this.Get("DataDisk.n.Description")
}

/*
空表示由系统默认分配，/dev/xvdb 开始到 /dev/xvdz
默认值：空
*/
func (this *CreateInstanceRequest) SetDataDisknDevice(v string){
	this.Set("DataDisk.n.Device", v)
}

func (this *CreateInstanceRequest) GetDataDisknDevice() string {
	return this.Get("DataDisk.n.Device")
}

func (this *CreateInstanceRequest) SetDataDisknDeleteWithInstance(v bool){
	this.SetBool("DataDisk.n.DeleteWithInstance", v)
}

func (this *CreateInstanceRequest) GetDataDisknDeleteWithInstance() bool {
	return this.GetBool("DataDisk.n.DeleteWithInstance")
}

func (this *CreateInstanceRequest) SetVSwitchId(v string){
	this.Set("VSwitchId", v)
}

func (this *CreateInstanceRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

func (this *CreateInstanceRequest) SetPrivateIpAddress(v string){
	this.Set("PrivateIpAddress", v)
}

func (this *CreateInstanceRequest) GetPrivateIpAddress() string {
	return this.Get("PrivateIpAddress")
}

func (this *CreateInstanceRequest) SetClientToken(v string){
	this.Set("ClientToken", v)
}

func (this *CreateInstanceRequest) GetClientToken() string {
	return this.Get("ClientToken")
}

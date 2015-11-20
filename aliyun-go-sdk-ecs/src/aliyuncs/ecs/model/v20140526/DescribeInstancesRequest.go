package v20140526

import "aliyuncs/core"

type DescribeInstancesRequest struct {
	core.Request
}

func NewDescribeInstancesRequest() *DescribeInstancesRequest {
	r := new(DescribeInstancesRequest)
	r.Init("Ecs", "2014-05-26", "DescribeInstances")
	return r
}

func (this *DescribeInstancesRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeInstancesRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeInstancesRequest) SetVpcId(v string) {
	this.Set("VpcId", v)
}

func (this *DescribeInstancesRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *DescribeInstancesRequest) SetVSwitchId(v string) {
	this.Set("VSwitchId", v)
}

func (this *DescribeInstancesRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

func (this *DescribeInstancesRequest) SetZoneId(v string) {
	this.Set("ZoneId", v)
}

func (this *DescribeInstancesRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *DescribeInstancesRequest) SetInstanceIds(v []string) {
	this.SetStringArray("InstanceIds", v)
}

func (this *DescribeInstancesRequest) GetInstanceIds() []string {
	return this.GetStringArray("InstanceIds")
}

func (this *DescribeInstancesRequest) SetInstanceNetworkType(v string) {
	this.Set("InstanceNetworkType", v)
}

func (this *DescribeInstancesRequest) GetInstanceNetworkType() string {
	return this.Get("InstanceNetworkType")
}

func (this *DescribeInstancesRequest) SetPrivateIpAddresses(v string) {
	this.Set("PrivateIpAddresses", v)
}

func (this *DescribeInstancesRequest) GetPrivateIpAddresses() string {
	return this.Get("PrivateIpAddresses")
}

func (this *DescribeInstancesRequest) SetInnerIpAddresses(v string) {
	this.Set("InnerIpAddresses", v)
}

func (this *DescribeInstancesRequest) GetInnerIpAddresses() string {
	return this.Get("InnerIpAddresses")
}

func (this *DescribeInstancesRequest) SetPublicIpAddresses(v string) {
	this.Set("PublicIpAddresses", v)
}

func (this *DescribeInstancesRequest) GetPublicIpAddresses() string {
	return this.Get("PublicIpAddresses")
}

func (this *DescribeInstancesRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *DescribeInstancesRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

/*
实例的付费方式。
PrePaid：预付费，即包年包月
PostPaid：后付费，即按量付费
*/
func (this *DescribeInstancesRequest) SetInstanceChargeType(v string) {
	this.Set("InstanceChargeType", v)
}

func (this *DescribeInstancesRequest) GetInstanceChargeType() string {
	return this.Get("InstanceChargeType")
}

/*
网络计费类型，PayByBandwidth | PayByTraffic两个值中的一个。预付费实例显示PayByBandwidth（按带宽计费）。
PayByTraffic：按流量计费
PayByBandwidth：按带宽计费
*/
func (this *DescribeInstancesRequest) SetInternetChargeType(v string) {
	this.Set("InternetChargeType", v)
}

func (this *DescribeInstancesRequest) GetInternetChargeType() string {
	return this.Get("InternetChargeType")
}

func (this *DescribeInstancesRequest) SetInstanceName(v string) {
	this.Set("InstanceName", v)
}

func (this *DescribeInstancesRequest) GetInstanceName() string {
	return this.Get("InstanceName")
}

func (this *DescribeInstancesRequest) SetImageId(v string) {
	this.Set("ImageId", v)
}

func (this *DescribeInstancesRequest) GetImageId() string {
	return this.Get("ImageId")
}

/*
实例状态，可选值：
Running
Starting
Stopping
Stopped
*/
func (this *DescribeInstancesRequest) SetStatus(v string) {
	this.Set("Status", v)
}

func (this *DescribeInstancesRequest) GetStatus() string {
	return this.Get("Status")
}

func (this *DescribeInstancesRequest) SetDeviceAvailable(v bool) {
	this.SetBool("DeviceAvailable", v)
}

func (this *DescribeInstancesRequest) GetDeviceAvailable() bool {
	return this.GetBool("DeviceAvailable")
}

func (this *DescribeInstancesRequest) SetIoOptimized(v bool) {
	this.SetBool("IoOptimized", v)
}

func (this *DescribeInstancesRequest) GetIoOptimized() bool {
	return this.GetBool("IoOptimized")
}

func (this *DescribeInstancesRequest) SetPageNumber(v int) {
	this.SetInt("PageNumber", v)
}

func (this *DescribeInstancesRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeInstancesRequest) SetPageSize(v int) {
	this.SetInt("PageSize", v)
}

func (this *DescribeInstancesRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

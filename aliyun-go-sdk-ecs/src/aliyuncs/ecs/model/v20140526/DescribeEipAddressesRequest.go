package v20140526

import "aliyuncs/core"

type DescribeEipAddressesRequest struct {
	core.Request
}

func NewDescribeEipAddressesRequest() *DescribeEipAddressesRequest {
	r := new(DescribeEipAddressesRequest)
	r.Init("Ecs", "2014-05-26", "DescribeEipAddresses")
	return r
}

func (this *DescribeEipAddressesRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeEipAddressesRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeEipAddressesRequest) SetStatus(v string) {
	this.Set("Status", v)
}

func (this *DescribeEipAddressesRequest) GetStatus() string {
	return this.Get("Status")
}

func (this *DescribeEipAddressesRequest) SetEipAddress(v string) {
	this.Set("EipAddress", v)
}

func (this *DescribeEipAddressesRequest) GetEipAddress() string {
	return this.Get("EipAddress")
}

func (this *DescribeEipAddressesRequest) SetAllocationId(v string) {
	this.Set("AllocationId", v)
}

func (this *DescribeEipAddressesRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

func (this *DescribeEipAddressesRequest) SetAssociatedInstanceType(v string) {
	this.Set("AssociatedInstanceType", v)
}

func (this *DescribeEipAddressesRequest) GetAssociatedInstanceType() string {
	return this.Get("AssociatedInstanceType")
}

func (this *DescribeEipAddressesRequest) SetAssociatedInstanceId(v string) {
	this.Set("AssociatedInstanceId", v)
}

func (this *DescribeEipAddressesRequest) GetAssociatedInstanceId() string {
	return this.Get("AssociatedInstanceId")
}

func (this *DescribeEipAddressesRequest) SetPageNumber(v int) {
	this.SetInt("PageNumber", v)
}

func (this *DescribeEipAddressesRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeEipAddressesRequest) SetPageSize(v int) {
	this.SetInt("PageSize", v)
}

func (this *DescribeEipAddressesRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

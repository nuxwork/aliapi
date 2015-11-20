package v20140526

import "aliyuncs/core"

type DescribeInstanceStatusRequest struct {
	core.Request
}

func NewDescribeInstanceStatusRequest() *DescribeInstanceStatusRequest {
	r := new(DescribeInstanceStatusRequest)
	r.Init("Ecs", "2014-05-26", "DescribeInstanceStatus")
	return r
}

func (this *DescribeInstanceStatusRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeInstanceStatusRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeInstanceStatusRequest) SetZoneId(v string) {
	this.Set("ZoneId", v)
}

func (this *DescribeInstanceStatusRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *DescribeInstanceStatusRequest) SetPageNumber(v int) {
	this.SetInt("PageNumber", v)
}

func (this *DescribeInstanceStatusRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeInstanceStatusRequest) SetPageSize(v int) {
	this.SetInt("PageSize", v)
}

func (this *DescribeInstanceStatusRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

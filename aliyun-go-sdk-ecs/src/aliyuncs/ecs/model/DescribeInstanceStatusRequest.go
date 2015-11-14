package model

import "aliyuncs/ecs"

type DescribeInstanceStatusRequest struct {
	ecs.Request
}

func NewDescribeInstanceStatusRequest() *DescribeInstanceStatusRequest{
	r := new(DescribeInstanceStatusRequest)
	r.Set("Action", "DescribeInstanceStatus")
	return r
}

func (this *DescribeInstanceStatusRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DescribeInstanceStatusRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeInstanceStatusRequest) SetZoneId(v string){
	this.Set("ZoneId", v)
}

func (this *DescribeInstanceStatusRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *DescribeInstanceStatusRequest) SetPageNumber(v int){
	this.SetInt("PageNumber", v)
}

func (this *DescribeInstanceStatusRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeInstanceStatusRequest) SetPageSize(v int){
	this.SetInt("PageSize", v)
}

func (this *DescribeInstanceStatusRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

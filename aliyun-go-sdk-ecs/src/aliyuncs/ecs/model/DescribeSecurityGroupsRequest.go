package model

import "aliyuncs/ecs"

type DescribeSecurityGroupsRequest struct {
	ecs.Request
}

func NewDescribeSecurityGroupsRequest() *DescribeSecurityGroupsRequest{
	r := new(DescribeSecurityGroupsRequest)
	r.Set("Action", "DescribeSecurityGroups")
	return r
}

func (this *DescribeSecurityGroupsRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DescribeSecurityGroupsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeSecurityGroupsRequest) SetVpcId(v string){
	this.Set("VpcId", v)
}

func (this *DescribeSecurityGroupsRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *DescribeSecurityGroupsRequest) SetPageNumber(v int){
	this.SetInt("PageNumber", v)
}

func (this *DescribeSecurityGroupsRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeSecurityGroupsRequest) SetPageSize(v int){
	this.SetInt("PageSize", v)
}

func (this *DescribeSecurityGroupsRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

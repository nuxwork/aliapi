package model

import "aliyuncs/ecs"

type DescribeVRoutersRequest struct {
	ecs.Request
}

func NewDescribeVRoutersRequest() *DescribeVRoutersRequest{
	r := new(DescribeVRoutersRequest)
	r.Set("Action", "DescribeVRouters")
	return r
}

func (this *DescribeVRoutersRequest) SetVRouterId(v string){
	this.Set("VRouterId", v)
}

func (this *DescribeVRoutersRequest) GetVRouterId() string {
	return this.Get("VRouterId")
}

func (this *DescribeVRoutersRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DescribeVRoutersRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeVRoutersRequest) SetPageNumber(v int){
	this.SetInt("PageNumber", v)
}

func (this *DescribeVRoutersRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeVRoutersRequest) SetPageSize(v int){
	this.SetInt("PageSize", v)
}

func (this *DescribeVRoutersRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

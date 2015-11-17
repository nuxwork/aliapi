package model

import "aliyuncs/ecs"

type DescribeRouteTablesRequest struct {
	ecs.Request
}

func NewDescribeRouteTablesRequest() *DescribeRouteTablesRequest{
	r := new(DescribeRouteTablesRequest)
	r.Set("Action", "DescribeRouteTables")
	return r
}

func (this *DescribeRouteTablesRequest) SetVRouterId(v string){
	this.Set("VRouterId", v)
}

func (this *DescribeRouteTablesRequest) GetVRouterId() string {
	return this.Get("VRouterId")
}

func (this *DescribeRouteTablesRequest) SetRouteTableId(v string){
	this.Set("RouteTableId", v)
}

func (this *DescribeRouteTablesRequest) GetRouteTableId() string {
	return this.Get("RouteTableId")
}

func (this *DescribeRouteTablesRequest) SetPageNumber(v int){
	this.SetInt("PageNumber", v)
}

func (this *DescribeRouteTablesRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeRouteTablesRequest) SetPageSize(v int){
	this.SetInt("PageSize", v)
}

func (this *DescribeRouteTablesRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}

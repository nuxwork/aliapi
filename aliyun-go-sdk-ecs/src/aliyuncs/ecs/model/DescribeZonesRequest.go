package model

import "aliyuncs/ecs"

type DescribeZonesRequest struct {
	ecs.Request
}

func NewDescribeZonesRequest() *DescribeZonesRequest{
	r := new(DescribeZonesRequest)
	r.Set("Action", "DescribeZones")
	return r
}

func (this *DescribeZonesRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DescribeZonesRequest) GetRegionId() string {
	return this.Get("RegionId")
}

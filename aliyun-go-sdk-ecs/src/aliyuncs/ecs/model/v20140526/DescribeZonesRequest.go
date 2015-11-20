package v20140526

import "aliyuncs/core"

type DescribeZonesRequest struct {
	core.Request
}

func NewDescribeZonesRequest() *DescribeZonesRequest {
	r := new(DescribeZonesRequest)
	r.Init("Ecs", "2014-05-26", "DescribeZones")
	return r
}

func (this *DescribeZonesRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeZonesRequest) GetRegionId() string {
	return this.Get("RegionId")
}

package v20140526

import "aliyuncs/core"

type DescribeInstanceVncUrlRequest struct {
	core.Request
}

func NewDescribeInstanceVncUrlRequest() *DescribeInstanceVncUrlRequest {
	r := new(DescribeInstanceVncUrlRequest)
	r.Init("Ecs", "2014-05-26", "DescribeInstanceVncUrl")
	return r
}

func (this *DescribeInstanceVncUrlRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *DescribeInstanceVncUrlRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *DescribeInstanceVncUrlRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeInstanceVncUrlRequest) GetRegionId() string {
	return this.Get("RegionId")
}

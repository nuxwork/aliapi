package model

import "aliyuncs/ecs"

type DescribeInstanceVncUrlRequest struct {
	ecs.Request
}

func NewDescribeInstanceVncUrlRequest() *DescribeInstanceVncUrlRequest{
	r := new(DescribeInstanceVncUrlRequest)
	r.Set("Action", "DescribeInstanceVncUrl")
	return r
}

func (this *DescribeInstanceVncUrlRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *DescribeInstanceVncUrlRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *DescribeInstanceVncUrlRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DescribeInstanceVncUrlRequest) GetRegionId() string {
	return this.Get("RegionId")
}

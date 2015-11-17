package model

import "aliyuncs/ecs"

type DescribeEipMonitorDataRequest struct {
	ecs.Request
}

func NewDescribeEipMonitorDataRequest() *DescribeEipMonitorDataRequest{
	r := new(DescribeEipMonitorDataRequest)
	r.Set("Action", "DescribeEipMonitorData")
	return r
}

func (this *DescribeEipMonitorDataRequest) SetAllocationId(v string){
	this.Set("AllocationId", v)
}

func (this *DescribeEipMonitorDataRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

func (this *DescribeEipMonitorDataRequest) SetStartTime(v string){
	this.Set("StartTime", v)
}

func (this *DescribeEipMonitorDataRequest) GetStartTime() string {
	return this.Get("StartTime")
}

func (this *DescribeEipMonitorDataRequest) SetEndTime(v string){
	this.Set("EndTime", v)
}

func (this *DescribeEipMonitorDataRequest) GetEndTime() string {
	return this.Get("EndTime")
}

func (this *DescribeEipMonitorDataRequest) SetPeriod(v int){
	this.SetInt("Period", v)
}

func (this *DescribeEipMonitorDataRequest) GetPeriod() int {
	return this.GetInt("Period")
}

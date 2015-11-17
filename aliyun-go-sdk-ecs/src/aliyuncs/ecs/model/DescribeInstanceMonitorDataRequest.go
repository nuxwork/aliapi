package model

import "aliyuncs/ecs"

type DescribeInstanceMonitorDataRequest struct {
	ecs.Request
}

func NewDescribeInstanceMonitorDataRequest() *DescribeInstanceMonitorDataRequest{
	r := new(DescribeInstanceMonitorDataRequest)
	r.Set("Action", "DescribeInstanceMonitorData")
	return r
}

func (this *DescribeInstanceMonitorDataRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *DescribeInstanceMonitorDataRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *DescribeInstanceMonitorDataRequest) SetStartTime(v string){
	this.Set("StartTime", v)
}

func (this *DescribeInstanceMonitorDataRequest) GetStartTime() string {
	return this.Get("StartTime")
}

func (this *DescribeInstanceMonitorDataRequest) SetEndTime(v string){
	this.Set("EndTime", v)
}

func (this *DescribeInstanceMonitorDataRequest) GetEndTime() string {
	return this.Get("EndTime")
}

func (this *DescribeInstanceMonitorDataRequest) SetPeriod(v int){
	this.SetInt("Period", v)
}

func (this *DescribeInstanceMonitorDataRequest) GetPeriod() int {
	return this.GetInt("Period")
}

package model

import "aliyuncs/ecs"

type DescribeDiskMonitorDataRequest struct {
	ecs.Request
}

func NewDescribeDiskMonitorDataRequest() *DescribeDiskMonitorDataRequest{
	r := new(DescribeDiskMonitorDataRequest)
	r.Set("Action", "DescribeDiskMonitorData")
	return r
}

func (this *DescribeDiskMonitorDataRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *DescribeDiskMonitorDataRequest) GetDiskId() string {
	return this.Get("DiskId")
}

func (this *DescribeDiskMonitorDataRequest) SetStartTime(v string){
	this.Set("StartTime", v)
}

func (this *DescribeDiskMonitorDataRequest) GetStartTime() string {
	return this.Get("StartTime")
}

func (this *DescribeDiskMonitorDataRequest) SetEndTime(v string){
	this.Set("EndTime", v)
}

func (this *DescribeDiskMonitorDataRequest) GetEndTime() string {
	return this.Get("EndTime")
}

func (this *DescribeDiskMonitorDataRequest) SetPeriod(v int){
	this.SetInt("Period", v)
}

func (this *DescribeDiskMonitorDataRequest) GetPeriod() int {
	return this.GetInt("Period")
}

package model

import "aliyuncs/ecs"

type StopInstanceRequest struct {
	ecs.Request
}

func NewStopInstanceRequest() *StopInstanceRequest{
	r := new(StopInstanceRequest)
	r.Set("Action", "StopInstance")
	return r
}

func (this *StopInstanceRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *StopInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *StopInstanceRequest) SetForceStop(v bool){
	this.SetBool("ForceStop", v)
}

func (this *StopInstanceRequest) GetForceStop() bool {
	return this.GetBool("ForceStop")
}

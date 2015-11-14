package model

import "aliyuncs/ecs"

type RebootInstanceRequest struct {
	ecs.Request
}

func NewRebootInstanceRequest() *RebootInstanceRequest{
	r := new(RebootInstanceRequest)
	r.Set("Action", "RebootInstance")
	return r
}

func (this *RebootInstanceRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *RebootInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

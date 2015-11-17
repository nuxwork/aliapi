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

func (this *RebootInstanceRequest) SetForceStop(v bool){
	this.SetBool("ForceStop", v)
}

func (this *RebootInstanceRequest) GetForceStop() bool {
	return this.GetBool("ForceStop")
}

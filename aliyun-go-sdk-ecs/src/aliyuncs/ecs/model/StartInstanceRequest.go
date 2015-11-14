package model

import "aliyuncs/ecs"

type StartInstanceRequest struct {
	ecs.Request
}

func NewStartInstanceRequest() *StartInstanceRequest{
	r := new(StartInstanceRequest)
	r.Set("Action", "StartInstance")
	return r
}

func (this *StartInstanceRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *StartInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

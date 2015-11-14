package model

import "aliyuncs/ecs"

type DeleteInstanceRequest struct {
	ecs.Request
}

func NewDeleteInstanceRequest() *DeleteInstanceRequest{
	r := new(DeleteInstanceRequest)
	r.Set("Action", "DeleteInstance")
	return r
}

func (this *DeleteInstanceRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *DeleteInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

package v20140526

import "aliyuncs/core"

type DeleteInstanceRequest struct {
	core.Request
}

func NewDeleteInstanceRequest() *DeleteInstanceRequest {
	r := new(DeleteInstanceRequest)
	r.Init("Ecs", "2014-05-26", "DeleteInstance")
	return r
}

func (this *DeleteInstanceRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *DeleteInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

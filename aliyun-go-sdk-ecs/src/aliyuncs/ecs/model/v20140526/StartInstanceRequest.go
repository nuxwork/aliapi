package v20140526

import "aliyuncs/core"

type StartInstanceRequest struct {
	core.Request
}

func NewStartInstanceRequest() *StartInstanceRequest {
	r := new(StartInstanceRequest)
	r.Init("Ecs", "2014-05-26", "StartInstance")
	return r
}

func (this *StartInstanceRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *StartInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

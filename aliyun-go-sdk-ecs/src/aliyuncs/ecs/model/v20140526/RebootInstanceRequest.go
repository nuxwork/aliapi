package v20140526

import "aliyuncs/core"

type RebootInstanceRequest struct {
	core.Request
}

func NewRebootInstanceRequest() *RebootInstanceRequest {
	r := new(RebootInstanceRequest)
	r.Init("Ecs", "2014-05-26", "RebootInstance")
	return r
}

func (this *RebootInstanceRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *RebootInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *RebootInstanceRequest) SetForceStop(v bool) {
	this.SetBool("ForceStop", v)
}

func (this *RebootInstanceRequest) GetForceStop() bool {
	return this.GetBool("ForceStop")
}

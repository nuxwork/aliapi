package v20140526

import "aliyuncs/core"

type DeleteVSwitchRequest struct {
	core.Request
}

func NewDeleteVSwitchRequest() *DeleteVSwitchRequest {
	r := new(DeleteVSwitchRequest)
	r.Init("Ecs", "2014-05-26", "DeleteVSwitch")
	return r
}

func (this *DeleteVSwitchRequest) SetVSwitchId(v string) {
	this.Set("VSwitchId", v)
}

func (this *DeleteVSwitchRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

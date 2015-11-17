package model

import "aliyuncs/ecs"

type DeleteVSwitchRequest struct {
	ecs.Request
}

func NewDeleteVSwitchRequest() *DeleteVSwitchRequest{
	r := new(DeleteVSwitchRequest)
	r.Set("Action", "DeleteVSwitch")
	return r
}

func (this *DeleteVSwitchRequest) SetVSwitchId(v string){
	this.Set("VSwitchId", v)
}

func (this *DeleteVSwitchRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

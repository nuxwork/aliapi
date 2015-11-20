package v20140526

import "aliyuncs/core"

type DeleteVpcRequest struct {
	core.Request
}

func NewDeleteVpcRequest() *DeleteVpcRequest {
	r := new(DeleteVpcRequest)
	r.Init("Ecs", "2014-05-26", "DeleteVpc")
	return r
}

func (this *DeleteVpcRequest) SetVpcId(v string) {
	this.Set("VpcId", v)
}

func (this *DeleteVpcRequest) GetVpcId() string {
	return this.Get("VpcId")
}

package model

import "aliyuncs/ecs"

type DeleteVpcRequest struct {
	ecs.Request
}

func NewDeleteVpcRequest() *DeleteVpcRequest{
	r := new(DeleteVpcRequest)
	r.Set("Action", "DeleteVpc")
	return r
}

func (this *DeleteVpcRequest) SetVpcId(v string){
	this.Set("VpcId", v)
}

func (this *DeleteVpcRequest) GetVpcId() string {
	return this.Get("VpcId")
}

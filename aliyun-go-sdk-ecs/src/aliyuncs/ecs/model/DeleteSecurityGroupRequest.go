package model

import "aliyuncs/ecs"

type DeleteSecurityGroupRequest struct {
	ecs.Request
}

func NewDeleteSecurityGroupRequest() *DeleteSecurityGroupRequest{
	r := new(DeleteSecurityGroupRequest)
	r.Set("Action", "DeleteSecurityGroup")
	return r
}

func (this *DeleteSecurityGroupRequest) SetSecurityGroupId(v string){
	this.Set("SecurityGroupId", v)
}

func (this *DeleteSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *DeleteSecurityGroupRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *DeleteSecurityGroupRequest) GetRegionId() string {
	return this.Get("RegionId")
}

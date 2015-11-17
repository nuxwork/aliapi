package model

import "aliyuncs/ecs"

type CreateSecurityGroupRequest struct {
	ecs.Request
}

func NewCreateSecurityGroupRequest() *CreateSecurityGroupRequest{
	r := new(CreateSecurityGroupRequest)
	r.Set("Action", "CreateSecurityGroup")
	return r
}

func (this *CreateSecurityGroupRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *CreateSecurityGroupRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CreateSecurityGroupRequest) SetSecurityGroupName(v string){
	this.Set("SecurityGroupName", v)
}

func (this *CreateSecurityGroupRequest) GetSecurityGroupName() string {
	return this.Get("SecurityGroupName")
}

func (this *CreateSecurityGroupRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *CreateSecurityGroupRequest) GetDescription() string {
	return this.Get("Description")
}

func (this *CreateSecurityGroupRequest) SetVpcId(v string){
	this.Set("VpcId", v)
}

func (this *CreateSecurityGroupRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *CreateSecurityGroupRequest) SetClientToken(v string){
	this.Set("ClientToken", v)
}

func (this *CreateSecurityGroupRequest) GetClientToken() string {
	return this.Get("ClientToken")
}

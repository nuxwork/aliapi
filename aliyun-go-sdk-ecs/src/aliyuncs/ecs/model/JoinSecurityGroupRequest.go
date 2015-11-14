package model

import "aliyuncs/ecs"

type JoinSecurityGroupRequest struct {
	ecs.Request
}

func NewJoinSecurityGroupRequest() *JoinSecurityGroupRequest{
	r := new(JoinSecurityGroupRequest)
	r.Set("Action", "JoinSecurityGroup")
	return r
}

func (this *JoinSecurityGroupRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *JoinSecurityGroupRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *JoinSecurityGroupRequest) SetSecurityGroupId(v string){
	this.Set("SecurityGroupId", v)
}

func (this *JoinSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

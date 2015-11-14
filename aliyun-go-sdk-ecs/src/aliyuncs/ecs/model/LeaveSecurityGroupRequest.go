package model

import "aliyuncs/ecs"

type LeaveSecurityGroupRequest struct {
	ecs.Request
}

func NewLeaveSecurityGroupRequest() *LeaveSecurityGroupRequest{
	r := new(LeaveSecurityGroupRequest)
	r.Set("Action", "LeaveSecurityGroup")
	return r
}

func (this *LeaveSecurityGroupRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *LeaveSecurityGroupRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *LeaveSecurityGroupRequest) SetSecurityGroupId(v string){
	this.Set("SecurityGroupId", v)
}

func (this *LeaveSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

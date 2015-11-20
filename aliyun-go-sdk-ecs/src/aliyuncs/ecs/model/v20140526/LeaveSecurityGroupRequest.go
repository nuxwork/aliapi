package v20140526

import "aliyuncs/core"

type LeaveSecurityGroupRequest struct {
	core.Request
}

func NewLeaveSecurityGroupRequest() *LeaveSecurityGroupRequest {
	r := new(LeaveSecurityGroupRequest)
	r.Init("Ecs", "2014-05-26", "LeaveSecurityGroup")
	return r
}

func (this *LeaveSecurityGroupRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *LeaveSecurityGroupRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *LeaveSecurityGroupRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *LeaveSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

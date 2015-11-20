package v20140526

import "aliyuncs/core"

type DeleteSecurityGroupRequest struct {
	core.Request
}

func NewDeleteSecurityGroupRequest() *DeleteSecurityGroupRequest {
	r := new(DeleteSecurityGroupRequest)
	r.Init("Ecs", "2014-05-26", "DeleteSecurityGroup")
	return r
}

func (this *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *DeleteSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *DeleteSecurityGroupRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DeleteSecurityGroupRequest) GetRegionId() string {
	return this.Get("RegionId")
}

package v20140526

import "aliyuncs/core"

type RevokeSecurityGroupEgressRequest struct {
	core.Request
}

func NewRevokeSecurityGroupEgressRequest() *RevokeSecurityGroupEgressRequest {
	r := new(RevokeSecurityGroupEgressRequest)
	r.Init("Ecs", "2014-05-26", "RevokeSecurityGroupEgress")
	return r
}

func (this *RevokeSecurityGroupEgressRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *RevokeSecurityGroupEgressRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *RevokeSecurityGroupEgressRequest) SetIpProtocol(v string) {
	this.Set("IpProtocol", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetIpProtocol() string {
	return this.Get("IpProtocol")
}

func (this *RevokeSecurityGroupEgressRequest) SetPortRange(v string) {
	this.Set("PortRange", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetPortRange() string {
	return this.Get("PortRange")
}

func (this *RevokeSecurityGroupEgressRequest) SetDestGroupId(v string) {
	this.Set("DestGroupId", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetDestGroupId() string {
	return this.Get("DestGroupId")
}

func (this *RevokeSecurityGroupEgressRequest) SetDestGroupOwnerAccount(v string) {
	this.Set("DestGroupOwnerAccount", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetDestGroupOwnerAccount() string {
	return this.Get("DestGroupOwnerAccount")
}

func (this *RevokeSecurityGroupEgressRequest) SetDestCidrIp(v string) {
	this.Set("DestCidrIp", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetDestCidrIp() string {
	return this.Get("DestCidrIp")
}

func (this *RevokeSecurityGroupEgressRequest) SetPolicy(v string) {
	this.Set("Policy", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetPolicy() string {
	return this.Get("Policy")
}

func (this *RevokeSecurityGroupEgressRequest) SetPriority(v string) {
	this.Set("Priority", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetPriority() string {
	return this.Get("Priority")
}

func (this *RevokeSecurityGroupEgressRequest) SetNicType(v string) {
	this.Set("NicType", v)
}

func (this *RevokeSecurityGroupEgressRequest) GetNicType() string {
	return this.Get("NicType")
}

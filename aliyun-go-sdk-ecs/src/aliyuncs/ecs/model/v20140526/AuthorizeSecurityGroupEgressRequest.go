package v20140526

import "aliyuncs/core"

type AuthorizeSecurityGroupEgressRequest struct {
	core.Request
}

func NewAuthorizeSecurityGroupEgressRequest() *AuthorizeSecurityGroupEgressRequest {
	r := new(AuthorizeSecurityGroupEgressRequest)
	r.Init("Ecs", "2014-05-26", "AuthorizeSecurityGroupEgress")
	return r
}

func (this *AuthorizeSecurityGroupEgressRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetIpProtocol(v string) {
	this.Set("IpProtocol", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetIpProtocol() string {
	return this.Get("IpProtocol")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetPortRange(v string) {
	this.Set("PortRange", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetPortRange() string {
	return this.Get("PortRange")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetDestGroupId(v string) {
	this.Set("DestGroupId", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetDestGroupId() string {
	return this.Get("DestGroupId")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetDestGroupOwnerAccount(v string) {
	this.Set("DestGroupOwnerAccount", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetDestGroupOwnerAccount() string {
	return this.Get("DestGroupOwnerAccount")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetDestCidrIp(v string) {
	this.Set("DestCidrIp", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetDestCidrIp() string {
	return this.Get("DestCidrIp")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetPolicy(v string) {
	this.Set("Policy", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetPolicy() string {
	return this.Get("Policy")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetPriority(v string) {
	this.Set("Priority", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetPriority() string {
	return this.Get("Priority")
}

func (this *AuthorizeSecurityGroupEgressRequest) SetNicType(v string) {
	this.Set("NicType", v)
}

func (this *AuthorizeSecurityGroupEgressRequest) GetNicType() string {
	return this.Get("NicType")
}

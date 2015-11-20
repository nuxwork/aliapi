package v20140526

import "aliyuncs/core"

type AuthorizeSecurityGroupRequest struct {
	core.Request
}

func NewAuthorizeSecurityGroupRequest() *AuthorizeSecurityGroupRequest {
	r := new(AuthorizeSecurityGroupRequest)
	r.Init("Ecs", "2014-05-26", "AuthorizeSecurityGroup")
	return r
}

func (this *AuthorizeSecurityGroupRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *AuthorizeSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *AuthorizeSecurityGroupRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *AuthorizeSecurityGroupRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *AuthorizeSecurityGroupRequest) SetIpProtocol(v string) {
	this.Set("IpProtocol", v)
}

func (this *AuthorizeSecurityGroupRequest) GetIpProtocol() string {
	return this.Get("IpProtocol")
}

func (this *AuthorizeSecurityGroupRequest) SetPortRange(v string) {
	this.Set("PortRange", v)
}

func (this *AuthorizeSecurityGroupRequest) GetPortRange() string {
	return this.Get("PortRange")
}

func (this *AuthorizeSecurityGroupRequest) SetSourceGroupId(v string) {
	this.Set("SourceGroupId", v)
}

func (this *AuthorizeSecurityGroupRequest) GetSourceGroupId() string {
	return this.Get("SourceGroupId")
}

func (this *AuthorizeSecurityGroupRequest) SetSourceGroupOwnerAccount(v string) {
	this.Set("SourceGroupOwnerAccount", v)
}

func (this *AuthorizeSecurityGroupRequest) GetSourceGroupOwnerAccount() string {
	return this.Get("SourceGroupOwnerAccount")
}

func (this *AuthorizeSecurityGroupRequest) SetSourceCidrIp(v string) {
	this.Set("SourceCidrIp", v)
}

func (this *AuthorizeSecurityGroupRequest) GetSourceCidrIp() string {
	return this.Get("SourceCidrIp")
}

func (this *AuthorizeSecurityGroupRequest) SetPolicy(v string) {
	this.Set("Policy", v)
}

func (this *AuthorizeSecurityGroupRequest) GetPolicy() string {
	return this.Get("Policy")
}

func (this *AuthorizeSecurityGroupRequest) SetPriority(v string) {
	this.Set("Priority", v)
}

func (this *AuthorizeSecurityGroupRequest) GetPriority() string {
	return this.Get("Priority")
}

func (this *AuthorizeSecurityGroupRequest) SetNicType(v string) {
	this.Set("NicType", v)
}

func (this *AuthorizeSecurityGroupRequest) GetNicType() string {
	return this.Get("NicType")
}

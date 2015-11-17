package model

import "aliyuncs/ecs"

type RevokeSecurityGroupRequest struct {
	ecs.Request
}

func NewRevokeSecurityGroupRequest() *RevokeSecurityGroupRequest{
	r := new(RevokeSecurityGroupRequest)
	r.Set("Action", "RevokeSecurityGroup")
	return r
}

func (this *RevokeSecurityGroupRequest) SetSecurityGroupId(v string){
	this.Set("SecurityGroupId", v)
}

func (this *RevokeSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *RevokeSecurityGroupRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *RevokeSecurityGroupRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *RevokeSecurityGroupRequest) SetIpProtocol(v string){
	this.Set("IpProtocol", v)
}

func (this *RevokeSecurityGroupRequest) GetIpProtocol() string {
	return this.Get("IpProtocol")
}

func (this *RevokeSecurityGroupRequest) SetPortRange(v string){
	this.Set("PortRange", v)
}

func (this *RevokeSecurityGroupRequest) GetPortRange() string {
	return this.Get("PortRange")
}

func (this *RevokeSecurityGroupRequest) SetSourceGroupId(v string){
	this.Set("SourceGroupId", v)
}

func (this *RevokeSecurityGroupRequest) GetSourceGroupId() string {
	return this.Get("SourceGroupId")
}

func (this *RevokeSecurityGroupRequest) SetSourceGroupOwnerAccount(v string){
	this.Set("SourceGroupOwnerAccount", v)
}

func (this *RevokeSecurityGroupRequest) GetSourceGroupOwnerAccount() string {
	return this.Get("SourceGroupOwnerAccount")
}

func (this *RevokeSecurityGroupRequest) SetSourceCidrIp(v string){
	this.Set("SourceCidrIp", v)
}

func (this *RevokeSecurityGroupRequest) GetSourceCidrIp() string {
	return this.Get("SourceCidrIp")
}

func (this *RevokeSecurityGroupRequest) SetPolicy(v string){
	this.Set("Policy", v)
}

func (this *RevokeSecurityGroupRequest) GetPolicy() string {
	return this.Get("Policy")
}

func (this *RevokeSecurityGroupRequest) SetPriority(v string){
	this.Set("Priority", v)
}

func (this *RevokeSecurityGroupRequest) GetPriority() string {
	return this.Get("Priority")
}

func (this *RevokeSecurityGroupRequest) SetNicType(v string){
	this.Set("NicType", v)
}

func (this *RevokeSecurityGroupRequest) GetNicType() string {
	return this.Get("NicType")
}

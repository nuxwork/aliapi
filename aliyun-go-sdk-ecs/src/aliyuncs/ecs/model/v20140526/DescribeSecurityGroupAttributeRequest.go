package v20140526

import "aliyuncs/core"

type DescribeSecurityGroupAttributeRequest struct {
	core.Request
}

func NewDescribeSecurityGroupAttributeRequest() *DescribeSecurityGroupAttributeRequest {
	r := new(DescribeSecurityGroupAttributeRequest)
	r.Init("Ecs", "2014-05-26", "DescribeSecurityGroupAttribute")
	return r
}

func (this *DescribeSecurityGroupAttributeRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *DescribeSecurityGroupAttributeRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *DescribeSecurityGroupAttributeRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeSecurityGroupAttributeRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeSecurityGroupAttributeRequest) SetNicType(v string) {
	this.Set("NicType", v)
}

func (this *DescribeSecurityGroupAttributeRequest) GetNicType() string {
	return this.Get("NicType")
}

func (this *DescribeSecurityGroupAttributeRequest) SetDirection(v string) {
	this.Set("Direction", v)
}

func (this *DescribeSecurityGroupAttributeRequest) GetDirection() string {
	return this.Get("Direction")
}
